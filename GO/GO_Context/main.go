package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logger *log.Logger
var key string = "name"
var key2 string = "name2"

func main() {
	logger = log.New(os.Stdout, "", log.Ltime)
	// 建立一個cancel context
	ctx, cancel := context.WithCancel(context.Background())

	// 建立數個withValue context, 繼承於ctx,  並給值
	valueCtx := context.WithValue(ctx, key, 1)
	valueCtx2 := context.WithValue(ctx, key2, 2)
	go watch(valueCtx)
	go watch2(valueCtx2)

	time.Sleep(4 * time.Second)

	logger.Println("任務停止")
	// 發出取消
	cancel()

	// 確保工作結束
	time.Sleep(1 * time.Second)
}

func watch(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			//接收到取消訊號
			logger.Println("任務", ctx.Value(key), ":任務停止...")
			return
		default:
			//取出值
			var value int = ctx.Value(key).(int)
			logger.Println("任務", ctx.Value(key), ":工作中")
			time.Sleep(time.Duration(value) * time.Second)
		}
	}
}

func watch2(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			//接收到取消訊號
			logger.Println("任務", ctx.Value(key2), ":任務停止...")
			return
		default:
			//取出值
			var value int = ctx.Value(key2).(int)
			logger.Println("任務", ctx.Value(key2), ":工作中")
			time.Sleep(time.Duration(value) * time.Second)
		}
	}
}
