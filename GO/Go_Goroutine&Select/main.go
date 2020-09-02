package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	BSSCtx, BSSCancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer BSSCancel()
	interrupt := false

	go func() {
		select {
		case <-BSSCtx.Done():
			interrupt = true
		}
	}()

	for i := 0; i < 5; i++ {
		if interrupt {
			break
		}
		fmt.Println(i)
		time.Sleep(time.Duration(1) * time.Second)
	}

}
