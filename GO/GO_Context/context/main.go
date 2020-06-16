package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx_life := context.Background()

	ctx_work, cancel_work := context.WithCancel(ctx_life)

	ctx_work_meeting, cancel_work_meeting := context.WithCancel(ctx_work)

	ctx_work_meeting_share, _ := context.WithCancel(ctx_work_meeting)

	ctx_work_coding, _ := context.WithCancel(ctx_work)

	go work(ctx_work)
	go coding(ctx_work_coding)
	time.Sleep(2 * time.Second)

	go meeting(ctx_work_meeting)
	time.Sleep(1 * time.Second)

	go share(ctx_work_meeting_share)
	time.Sleep(1 * time.Second)
	cancel_work_meeting()

	time.Sleep(4 * time.Second)
	cancel_work()

	time.Sleep(1 * time.Millisecond)
}

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("go home")
			return
		default:
			fmt.Println("working")
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func meeting(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Print(" (meeting) ")
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func share(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Print("{shareing}")
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func coding(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Print(" (coding) ")
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}
