package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	timeoutError = errors.New("timeout error")
)

func doWork(ctx context.Context, duration time.Duration) error {
	timer := time.NewTimer(duration)
	defer timer.Stop()

	workEmulationTimer := time.NewTimer(500 * time.Millisecond)
	defer workEmulationTimer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return timeoutError
	case <-workEmulationTimer.C:
		return nil
	}
}

func main() {
	// close ctx error
	fmt.Println(1)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if err := doWork(ctx, time.Second); err != nil {
		fmt.Printf("%v\n", err)
	}
	// timeout error
	fmt.Println(2)
	if err := doWork(context.Background(), 0); err != nil {
		fmt.Printf("%v\n", err)
	}
	// no error
	fmt.Println(3)
	if err := doWork(context.Background(), time.Second); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println("Work successfuly done")
	}
}
