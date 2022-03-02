package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	Run(ctx, t)
	for err := range Run(ctx, t) {
		fmt.Println("Err:", err)
		// cancel()
	}
	// Run(func() error {
	// 	if time.Now().Second()%2 == 0 {
	// 		return errors.New("2 test error")
	// 	}
	// 	return nil
	// },
	// 	func() error {
	// 		if time.Now().Second()%3 == 0 {
	// 			return errors.New("3 test error")
	// 		}
	// 		return nil
	// 	},
	// 	func() error {
	// 		if time.Now().Second()%4 == 0 {
	// 			return errors.New("4 test error")
	// 		}
	// 		return nil
	// 	})
}

// func Run(tasks ...func() error) (err error) {
// 	g := new(errgroup.Group)
// 	for _, t := range tasks {
// 		g.Go(t)
// 	}
// 	// Wait for all go routines to complete.
// 	if err = g.Wait(); err != nil {
// 		log.Error(err)

// 		return
// 	}
// 	log.Info("No errors occurred")

// 	return
// }

func Run(ctx context.Context, tasks ...func(context.Context, *sync.WaitGroup, chan<- error)) (errCh chan error) {
	n := len(tasks)
	errCh = make(chan error, n)

	w := &sync.WaitGroup{}
	w.Add(n)

	for _, t := range tasks {
		go t(ctx, w, errCh)
	}

	w.Wait()
	close(errCh)

	return
}

func t(ctx context.Context, wg *sync.WaitGroup, errChan chan<- error) {
	if time.Now().Second()%2 == 0 {
		errChan <- errors.New("2 test error")
	}
}
