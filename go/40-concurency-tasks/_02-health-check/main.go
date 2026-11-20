package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"https://google.com",
	"https://github.com",
	"https://stackoverflow.com",
	"https://golang.org",
	"https://example.com",
	"https://httpbin.org/status/500",
}

type outChItem struct {
	url string
	err error
}

func healhCheck(ctx context.Context, urls []string) map[string]error {
	workers := 3
	result := make(map[string]error, len(urls))

	var wg sync.WaitGroup
	in, out := make(chan string), make(chan outChItem)

	go func() {
		defer close(in)

		for _, url := range urls {
			select {
			case <-ctx.Done():
				return
			default:
				in <- url
			}
		}
	}()

	for range workers {
		wg.Go(func() {
			for url := range in {
				select {
				case <-ctx.Done():
					return
				default:
					func() {
						ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
						defer cancel()

						if req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil); err != nil {
							out <- outChItem{url, err}
						} else {
							res, err := http.DefaultClient.Do(req)
							if err != nil {
								out <- outChItem{url, err}
							} else {
								defer res.Body.Close()

								if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices {
									out <- outChItem{url, nil}
								} else {
									out <- outChItem{url, fmt.Errorf("status code: %d", res.StatusCode)}
								}
							}
						}
					}()
				}
			}
		})
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	for r := range out {
		result[r.url] = r.err
	}

	return result
}

func main() {
	m := healhCheck(context.Background(), urls)
	for k, e := range m {
		fmt.Printf("%s: %v\n", k, e)
	}
}
