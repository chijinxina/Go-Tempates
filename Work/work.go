//Work包展示如何使用无缓冲的通道来创建一个gorourine池
//Work包管理一个goroutine池来管理工作
package Work

import "sync"

//Worker接口
type Worker interface {
	Task()
}

//Pool提供一个goroutine池
//这个池可以完成任何已提交的Worker任务
type Pool struct {
	works chan Worker
	wg    sync.WaitGroup
}

//New创建一个新的工作池
func New(maxGoroutine int) *Pool {
	p := Pool{
		works: make(chan Worker),
	}
	p.wg.Add(maxGoroutine)

	//每个协程执行任务
	for i := 0; i < maxGoroutine; i++ {
		go func() {
			//for range会一直阻塞循环直到从works通道中接收到一个Worker
			// 一旦通道被关闭for range就会停止循环
			for w := range p.works {
				//执行任务
				w.Task()
			}
			//Done通知任务已完成
			p.wg.Done()
		}()
	}
	return &p
}

//Run提交任务到工作池
func (p *Pool) Run(w Worker) {
	p.works <- w
}

//Shutdown等待关闭所有goroutine停止工作
func (p *Pool) Shutdown() {
	close(p.works)
	p.wg.Wait()
}
