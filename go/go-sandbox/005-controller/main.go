package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
)

// Controller инкапсулирует informer, lister и workqueue
type Controller struct {
	clientset kubernetes.Interface
	queue     workqueue.TypedRateLimitingInterface[string]
	informer  cache.SharedIndexInformer
}

func NewController(clientset kubernetes.Interface, informer cache.SharedIndexInformer) *Controller {
	queue := workqueue.NewTypedRateLimitingQueue(
		workqueue.DefaultTypedControllerRateLimiter[string](),
	)

	c := &Controller{
		clientset: clientset,
		queue:     queue,
		informer:  informer,
	}

	// Подписываемся на события Informer и отправляем ключи в очередь
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				fmt.Printf("[EVENT] Added Pod: %s\n", key)
				c.queue.Add(key)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			newPod := newObj.(*corev1.Pod)
			oldPod := oldObj.(*corev1.Pod)

			// Фильтруем пустые апдейты (например, только ResourceVersion)
			if oldPod.ResourceVersion == newPod.ResourceVersion {
				return
			}

			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err == nil {
				fmt.Printf("[EVENT] Updated Pod: %s\n", key)
				c.queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				fmt.Printf("[EVENT] Deleted Pod: %s\n", key)
				c.queue.Add(key)
			}
		},
	})

	return c
}

// Run запускает воркера и ожидает синхронизации кэша
func (c *Controller) Run(ctx context.Context) {
	defer c.queue.ShutDown()

	fmt.Println("Ожидание синхронизации кэша Informer...")
	if !cache.WaitForCacheSync(ctx.Done(), c.informer.HasSynced) {
		fmt.Println("Ошибка: не удалось синхронизировать кэш")
		return
	}
	fmt.Println("Кэш успешно синхронизирован. Запуск воркера...")

	// Запускаем воркер в цикле
	go func() {
		for c.processNextItem(ctx) {
		}
	}()

	<-ctx.Done()
	fmt.Println("Завершение работы контроллера...")
}

// processNextItem достает элемент из очереди и обрабатывает
func (c *Controller) processNextItem(ctx context.Context) bool {
	key, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Done(key)

	err := c.reconcile(ctx, key)
	if err != nil {
		fmt.Printf("Ошибка при обработке %s: %v. Повтор...\n", key, err)
		c.queue.AddRateLimited(key)
		return true
	}

	c.queue.Forget(key)
	return true
}

// reconcile — основа контроллера (Reconciliation Loop)
func (c *Controller) reconcile(ctx context.Context, key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}

	// Читаем из локального кэша Informer
	item, exists, err := c.informer.GetIndexer().GetByKey(key)
	if err != nil {
		return err
	}

	if !exists {
		fmt.Printf("[RECONCILE] Pod %s/%s был удален, очищаем связанные ресурсы если нужно\n", namespace, name)
		return nil
	}

	pod := item.(*corev1.Pod)
	fmt.Printf("[RECONCILE] Pod %s/%s в фазе: %s (IP: %s)\n",
		pod.Namespace, pod.Name, pod.Status.Phase, pod.Status.PodIP)

	return nil
}

func getKubeConfig() (*rest.Config, error) {
	// Сначала пробуем in-cluster конфиг (если запущен внутри пода)
	config, err := rest.InClusterConfig()
	if err == nil {
		return config, nil
	}

	// Иначе берем ~/.kube/config
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		if home := homedir.HomeDir(); home != "" {
			kubeconfigPath = filepath.Join(home, ".kube", "config")
		}
	}

	return clientcmd.BuildConfigFromFlags("", kubeconfigPath)
}

func main() {
	config, err := getKubeConfig()
	if err != nil {
		panic(fmt.Sprintf("Не удалось загрузить kubeconfig: %v", err))
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("Не удалось создать clientset: %v", err))
	}

	// SharedInformerFactory с ресинком раз в 10 минут
	factory := informers.NewSharedInformerFactory(clientset, 10*time.Minute)
	podInformer := factory.Core().V1().Pods().Informer()

	controller := NewController(clientset, podInformer)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Запускаем информеры
	factory.Start(ctx.Done())

	// Запускаем контроллер
	controller.Run(ctx)
}
