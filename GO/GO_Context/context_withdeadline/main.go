package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//ExampleWithDeadline()
	selecttry()
	wg.Wait()
}

var wg sync.WaitGroup

func ExampleWithDeadline() {
	d := time.Now().Add(11 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()
	fmt.Println(d)
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	// Output:
	// context deadline exceeded
}
func selecttry() {
	ctx, cancel := context.WithCancel(context.Background())
	intervalTimeout := time.After(time.Duration(5) * time.Second)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println(2)
		for ctx.Err() == nil {
			select {
			case <-ctx.Done():
			case <-intervalTimeout:
				fmt.Println("intervalTimeout")
				cancel()
				return
			case <-time.After(5 * time.Second):
				// if ntp.cancel {
				// 	ntp.wg.Done()
				// 	cancel()
				// 	return
				// }
				fmt.Println("5 seconds check")
			}
		}
	}()

}
