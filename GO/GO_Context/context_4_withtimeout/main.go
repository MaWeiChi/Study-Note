package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go doForever(ctx)
	time.Sleep(8 * time.Second)
	cancel()
	time.Sleep(8 * time.Second)
}

func run(ctx context.Context) {
	for i := 0; i < 20; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

}

func doForever(ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("End")
			return
		default:
			i++
			if i == 1 {
				go run(ctx)
			}
			if ctx.Done() {

			}
			fmt.Println("do")
			time.Sleep(1 * time.Second)
		}
	}
}
