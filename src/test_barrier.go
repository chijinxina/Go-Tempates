package main

import (
	"barrier"
	"fmt"
	"sync"
)

func main() {
	b := barrier.NewBarrierCond(3) //新建栅栏 3个协程

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1) //添加协程

		go func(index int) {
			defer wg.Done() //协程运行结束 通知wg
			fmt.Printf("Goroutine: %d print A\n", index)
			b.Wait()
			fmt.Printf("Goroutine: %d print B\n", index)
			b.Wait()
			fmt.Printf("Goroutine: %d print C\n", index)
		}(i)
	}

	wg.Wait() //等待所有协程运行结束
}
