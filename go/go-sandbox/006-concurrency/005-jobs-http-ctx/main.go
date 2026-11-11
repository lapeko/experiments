package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

const swapiUrl = "https://swapi.info/api/people/"

type person struct {
	Name string `json:"name"`
}

func getSwapiPerson(ctx context.Context, id int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%d", swapiUrl, id), nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var p person
	if err := json.NewDecoder(res.Body).Decode(&p); err != nil {
		return "", err
	}
	return p.Name, nil
}

func jobRunner(ctx context.Context, in <-chan int, out chan<- string) {
	for {
		select {
		case <-ctx.Done():
			{
				fmt.Println("Done")
				in = nil
				return
			}
		case id, ok := <-in:
			{
				if ok {
					p, err := getSwapiPerson(ctx, id)
					if err != nil {
						fmt.Printf("%v", err)
					} else {
						select {
						case out <- p:
						case <-ctx.Done():
							{
								return
							}
						}
					}
				} else {
					in = nil
					return
				}
			}
		}
	}
}

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	size := 5
	ids := make([]int, size)
	for i := range size {
		ids[i] = i + 1
	}

	in := make(chan int)
	out := make(chan string)
	var wg sync.WaitGroup

	go func() {
		defer close(in)
		for _, id := range ids {
			select {
			case in <- id:
			case <-ctx.Done():
				{
					return
				}
			}
		}
	}()

	for range 3 {
		wg.Go(func() {
			jobRunner(ctx, in, out)
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var results []string

	for res := range out {
		if res != "" {
			results = append(results, res)
		}
	}

	fmt.Println(strings.Join(results, "\n"))
}
