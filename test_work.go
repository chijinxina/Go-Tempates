package main

import (
	"github.com/chijinxina/Go-Tempates/Work"
	"log"
	"sync"
	"time"
)

//展示如何使用work包创建一个goroutine池并完成工作

//names提供一组用来显示的名字
var names = []string{
	"Steve",
	"Bob",
	"Mary",
	"Therese",
	"Jason",
}

//namePrinter使用特定的方式打印名字
type namePrinter struct {
	name string
}

//Task实现Worker接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(1 * time.Second)
}

//main主程序入口
func main() {
	//使用2个goroutine来创建工作池
	p := Work.New(100)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		//迭代names切片
		for _, name := range names {
			//创建一个namePrinter并提供指定的名字
			np := namePrinter{
				name: name,
			}
			go func() {
				//将任务提交执行，当Run返回时我们就知道任务已经完成
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	//让工作池停止工作，等待所有现有的工作完成
	p.Shutdown()
}
