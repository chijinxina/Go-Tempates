//Runner包管理处理任务的运行和生命周期
//类似线程池
package Runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//Runner在给定的超时时间内执行一组任务，并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	//interrupt通道报告从操作系统发送的信号
	interrupt chan os.Signal

	//Complete通道报告处理任务已经完成
	complete chan error

	//timeout报告处理任务已经超时
	timeout <-chan time.Time

	//tasks持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

//ErrTimeout会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout.")

//ErrInterrupt会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt.")

//New返回一个新的准备使用的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//Add将一个任务绑定到Runner上
//这个任务是一个接受一个int类型的id作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//gotInterrupt验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	//当中断事件被触发时发出信号
	case <-r.interrupt:
		//停止接收后续的任务信号
		signal.Stop(r.interrupt)
		return true
	//继续正常执行，不阻塞
	default:
		return false
	}
}

//run执行每一个已经注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		//检测到系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		//执行已注册的任务
		task(id)
	}
	return nil
}

//Start执行所有的任务，并监视通道事件
func (r *Runner) Start() error {
	//我们希望接收所有的中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	//用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	//当任务处理完成时发出信号
	case err := <-r.complete:
		return err
	//当任务处理程序运行超时时发出信号
	case <-r.timeout:
		return ErrTimeout
	}
}
