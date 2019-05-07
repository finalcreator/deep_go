package _04_Cancel_By_Context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				fmt.Printf("No.%d channel is working!\n",i)
				time.Sleep(time.Millisecond * 200)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second * 5)
}
