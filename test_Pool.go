package main

import (
	"github.com/chijinxina/Go-Tempates/Pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25 //要使用的携程数量
	pooledResources = 2  //资源池中资源的数量（资源池容量）
)

//dbConnection模拟要共享的资源（数据库连接池）
type dbConnection struct {
	ID int32
}

//Close实现了io.Closer接口，以便dbConnection可以被资源池管理
//Close用来完成任意资源的释放管理
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

//idCounter用来给每个数据库连接分配一个独一无二的id
var idCounter int32

//createConnection是一个工厂函数，当需要一个新连接时，资源池会调用这个工厂函数构造新的数据库连接
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

//main主程序入口
func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	//创建用来管理数据库连接的内存池
	p, err := Pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	//使用数据库连接池来完成查询
	for query := 0; query < maxGoroutines; query++ {
		//每个协程需要自己复制一份要查询值的副本，不然所有查询都会共享同一查询变量
		go func(q int) {
			performQueries(q, p)
			//Done用来通知main当前协程工作已完成
			wg.Done()
		}(query)
	}

	//等待协程结束
	wg.Wait()
	//关闭资源池
	log.Println("Shutdown Program.")
	p.Close()
}

//performQueries用来测试数据库连接
func performQueries(query int, pool *Pool.Pool) {
	//从资源池中请求一个数据库连接
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	//将该链接释放回资源池中
	defer pool.Release(conn)
	//用等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
