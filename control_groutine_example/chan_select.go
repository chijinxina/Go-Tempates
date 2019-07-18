package main

import (
	"fmt"
	"time"
)

func main() {

	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("退出！")
				return
			default:
				fmt.Println("运行...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("通知退出")
	stop <- true
	time.Sleep(10 * time.Second)
}
