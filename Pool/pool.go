//本示例包使用有缓冲的通道实现资源池，来管理可以在任意数量的goroutine之间共享及独立使用的资源（如共享数据库连接或者内存缓冲池常用）
//如果goroutine需要从资源池中获取资源，它可以从Pool中申请，使用完成后归还到资源池中。
package Pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

//包Pool管理用户定义的一组资源，该资源必须实现io.Closer接口
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed表示请求（Acquire)了一个已经关闭了的资源池
var ErrPoolClosed = errors.New("Pool has been closed.")

//New创建一个用来管理资源的池，这个池需要一个可以分配新资源的函数，并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

//Acquire从资源池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	//检查是否有空闲的资源
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

		// 因为没有空闲的资源可用，多以需要提供一个新的资源
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

//Release将一个使用后的资源放回到池中去
func (p *Pool) Release(r io.Closer) {
	// 保证本操作和close操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	// If the pool is closed, discard the resource.
	if p.closed {
		r.Close()
		return
	}

	select {
	// Attempt to place the new resource on the queue.
	case p.resources <- r:
		log.Println("Release:", "In Queue")

		// If the queue is already at cap we close the resource.
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close会让资源池停止工作，并且关闭所有现有的资源
func (p *Pool) Close() {
	// 保证本操作与Release操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	//如果Pool已经关闭，什么也不做
	if p.closed {
		return
	}

	// 将资源池关闭
	p.closed = true

	// 在清空通道资源之前将通道关闭
	// 如果不这么做会发生死锁
	close(p.resources)

	//遍历关闭资源
	for r := range p.resources {
		r.Close()
	}
}
