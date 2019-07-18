package barrier

import "sync"

// 抽象一个栅栏
type Barrier interface {
	Wait()
}

// 利用条件变量实现的栅栏
type BarrierCond struct {
	curCnt int        //当前等待的协程计数
	maxCnt int        //最大等待的协程计数
	cond   *sync.Cond //条件变量
}

//条件变量栅栏 接口函数实现
func (b *BarrierCond) Wait() {
	b.cond.L.Lock() //临界区 上锁

	if b.curCnt--; b.curCnt > 0 { // curCnt > 0 阻塞
		b.cond.Wait()
	} else { // curCnt = 0 唤醒协程
		b.cond.Broadcast()
		b.curCnt = b.maxCnt
	}
	b.cond.L.Unlock() //临界区结束 解锁
}

func NewBarrierCond(maxCnt int) Barrier {
	mutex := new(sync.Mutex)    //互斥锁
	cond := sync.NewCond(mutex) //条件变量
	return &BarrierCond{curCnt: maxCnt, maxCnt: maxCnt, cond: cond}
}
