package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var callEnd bool
var wg sync.WaitGroup
var updateTime int = 0

func main() {
	chann()
	time.Sleep(time.Duration(60) * time.Second)
}

func chann() {
	defer fmt.Println("chann major function end")
	callEnd = true
	wg.Wait()
	wg.Add(1)
	callEnd = false
	go func() {
		defer wg.Done()
		for !callEnd {
			updateTime++
			if updateTime == 5 {
				return
			}
			ctx, cancel := context.WithCancel(context.TODO())
			go noises(ctx, updateTime)
			intervalTimeout := time.After(time.Duration(5) * time.Second)
			fmt.Println(time.Now().Add(5))
			fmt.Println("for select start")
			for ctx.Err() == nil {
				select {
				case <-ctx.Done():
				case <-intervalTimeout:
					fmt.Println("interval Timeout")
					cancel()
				case <-time.After(5 * time.Second):
					if callEnd {
						cancel()
					}
				}
			}
			fmt.Println("for select end")
		}
	}()

}

func noises(ctx context.Context, times int) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("shut up")
			return
		default:
			fmt.Print(times)
			fmt.Print(" noise ")
			fmt.Println(time.Now())
			time.Sleep(time.Duration(1) * time.Second)
		}

	}
}
