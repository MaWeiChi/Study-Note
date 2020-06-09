package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func doForever(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logger.Println(ctx.Err())
			return
		default:
			logger.Println("doForever")
			time.Sleep(1 * time.Second)
		}
	}
}

// func do1second(ctx context.Context) {

// 	for {
// 		fmt.Println(1)
// 		time.Sleep(1 * time.Second)
// 	}

// }

func main() {
	logger = log.New(os.Stdout, "", log.Ltime)
	// 建立一個timeout context,  3秒後沒返回就發出超時
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()
	logger.Println("start")
	go doForever(ctx)
	//go do1second(ctx)

	time.Sleep(8 * time.Second)
	logger.Println("end")
}
