package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

var logger *log.Logger
var key string = "name"

func main() {
	// // base: func handle
	// // ---------------------------------------------------------------------
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cancel()
	// go handle(ctx, 2000*time.Millisecond)
	// select {
	// case <-ctx.Done():
	// 	fmt.Println("main", ctx.Err())
	// }
	// // ---------------------------------------------------------------------

	// // watch - context.WithValue()
	// // ---------------------------------------------------------------------
	// logger = log.New(os.Stdout, "", log.Ltime)

	// ctx, cancel := context.WithCancel(context.Background())

	// valueCtx1 := context.WithValue(ctx, key, 1)
	// valueCtx2 := context.WithValue(ctx, key, 2)

	// go watch(valueCtx1)
	// go watch(valueCtx2)

	// time.Sleep(8 * time.Second)

	// logger.Println("mission stop")
	// cancel()
	// // ---------------------------------------------------------------------

	// time.Sleep(1 * time.Second)
	// ---------------------------------------------------------------------
	logger = log.New(os.Stdout, "", log.Ltime)

	d := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	logger.Println("Start")
	go do(ctx)

	time.Sleep(3 * time.Second)
	// ---------------------------------------------------------------------
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logger.Println("mission", ctx.Value(key), ": stop")
			return
		default:
			var value int = ctx.Value(key).(int)
			logger.Println("mission", ctx.Value(key), ": running")
			time.Sleep(time.Duration(value) * time.Second)
		}
	}
}

func do(ctx context.Context) {
	if deadline, ok := ctx.Deadline(); ok == true {
		logger.Println("deadline: ", deadline)

	}
	for {
		select {
		case <-ctx.Done():
			logger.Println(ctx.Err())
			return
		default:
			logger.Println("do")
			time.Sleep(1 * time.Second)
		}
	}
}
