package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var overLimitError = errors.New("over limit")

func main() {
	ctx := context.Background()

	rl := NewRateLimiter2(5)

	for i := range 100 {
		go func() {
			if err := rl.Wait(ctx); err != nil {
				fmt.Printf("%d %v\n", i+1, err)
			} else {
				fmt.Printf("%d ok\n", i)
			}
		}()
		time.Sleep(150 * time.Millisecond)
	}
}
