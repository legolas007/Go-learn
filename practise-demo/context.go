package main

import (
	"fmt"
	"context"
	"time"
)

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <- ctx.Done():
				fmt.Println("监控退出，stop")
				return
			default:
				fmt.Println("监控中")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("ok,watch stop")
	cancel()
	time.Sleep(5 * time.Second)
}
