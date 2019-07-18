package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())

	ctx1, _ := context.WithTimeout(context.Background(), time.Second*3)

	//监控单个goroutine
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("退出！")
				return

			default:
				fmt.Println("运行中...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx1)

	////监控单个goroutine
	//	go func(ctx context.Context) {
	//		for {
	//			select {
	//			case <-ctx.Done():
	//				fmt.Println("退出！")
	//				return
	//
	//			default:
	//				fmt.Println("运行中...")
	//		time.Sleep(2 * time.Second)
	//	}
	//}
	//	}(ctx)

	////监控多个goroutine
	//go watch(ctx, "111")
	//go watch(ctx, "222")
	//go watch(ctx, "333")

	time.Sleep(6 * time.Second)
	fmt.Println("发送停止信息")
	//cancel()

	time.Sleep(10 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name + " 退出了!")
			return
		default:
			fmt.Println(name + " 运行中...")
			time.Sleep(1 * time.Second)
		}
	}
}
