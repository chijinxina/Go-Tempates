package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

//并发应用场景结构体
type Scenario struct {
	Name        string
	Description []string
	Example     []string
	RunExample  func()
}

//并发场景1：简单并发执行任务
var s1 = &Scenario{
	"s1",
	[]string{"简单并发执行任务"},
	[]string{"比如并发的请求后端的某一个接口"},
	RunScenario1,
}

//并发场景2：持续一定时间的高并发模型
var s2 = &Scenario{
	"s2",
	[]string{"持续一定时间的高并发模型"},
	[]string{"在规定时间内，持续地高并发请求后端服务，防止服务死循环"},
	RunScenario2,
}

//并发场景3：基于大数据量的并发任务模型，goroutine worker pool
var s3 = &Scenario{
	"s3",
	[]string{"基于大数据量的并发任务模型，goroutine worker pool"},
	[]string{"比如技术支持要给某个客户删除几个TB/GB的文件"},
	RunScenario3,
}

//并发场景4：等待异步任务的执行结果（goroutine + select + channel）
var s4 = &Scenario{
	"s4",
	[]string{"等待异步任务的执行结果（goroutine + select + channel）"},
	[]string{"比如事件循环"},
	RunScenario4,
}

//并发场景5：定时反馈结果（Ticker）
var s5 = &Scenario{
	"s5",
	[]string{"定时反馈结果（Ticker）"},
	[]string{"比如测试上传接口的性能，要实时给出指标：吞吐量,IOPS,成功率等"},
	RunScenario5,
}

var Scenarios []*Scenario

//包初始化
func init() {
	Scenarios = append(Scenarios, s1) //并发场景1：简单并发执行任务
	Scenarios = append(Scenarios, s2) //并发场景2：持续一定时间的高并发模型
	Scenarios = append(Scenarios, s3) //并发场景3：基于大数据量的并发任务模型，goroutine worker pool
	Scenarios = append(Scenarios, s4) //并发场景4：等待异步任务的执行结果（goroutine + select + channel）
	Scenarios = append(Scenarios, s5) //并发场景5：定时反馈结果（Ticker）
}

//主程序入口 常用的并发和同步场景
func main() {
	if len(os.Args) == 1 {
		fmt.Println("请选择使用场景：")
		for _, sc := range Scenarios {
			fmt.Printf("场景：%s, ", sc.Name)
			fmt.Sprintf("场景描述： %s \n", sc.Description)
		}
		return
	}

	for _, arg := range os.Args[1:] {
		sc := matchScenario(arg)
		if sc != nil {
			fmt.Printf("场景描述：%s \n", sc.Description)
			fmt.Printf("场景举例：%s \n", sc.Example)
			sc.RunExample()
		}
	}
}

func matchScenario(name string) *Scenario {
	for _, sc := range Scenarios {
		if sc.Name == name {
			return sc
		}
	}
	return nil
}

var doSomething = func(i int) string {
	time.Sleep(time.Millisecond * time.Duration(10))
	fmt.Printf("Goroutine %d do things ... \n", i)
	return fmt.Sprintf("Goroutine %d", i)
}
var takeSomething = func(res string) string {
	time.Sleep(time.Millisecond * time.Duration(10))
	tmp := fmt.Sprintf("Take result from %s... \n", res)
	fmt.Printf(tmp)
	return tmp
}

//并发场景1执行体：简单并发执行任务
func RunScenario1() {
	count := 10
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			doSomething(index)
		}(i)
	}
	wg.Wait()
}

//并发场景2执行体：持续一定时间的高并发模型
func RunScenario2() {

	timeout := time.Now().Add(time.Second + time.Duration(10))
	n := runtime.NumCPU()

	//在Go语言中,有一种特殊的struct{}类型的channel
	//它不能被写入任何数据,只有通过close()函数进行关闭操作,才能进行输出操作
	//struct类型的channel不占用任何内存
	//等待某任务的结束
	waitForAll := make(chan struct{})
	done := make(chan struct{})
	concurrentCount := make(chan struct{}, n)

	for i := 0; i < n; i++ {
		concurrentCount <- struct{}{}
	}

	go func() {
		for time.Now().Before(timeout) {
			<-done
			concurrentCount <- struct{}{}
		}
		waitForAll <- struct{}{}
	}()

	var wg sync.WaitGroup

	wg.Wait()
}
