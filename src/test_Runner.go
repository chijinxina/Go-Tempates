package main

import (
	"Runner"
	"log"
	"os"
	"time"
)

//timeout规定了Runner必须在多少秒时间内处理完成
const timeout = 10 * time.Second

//main主程序入口
func main() {
	log.Println("Starting work.")

	//创建Runner并指定任务执行超时时间
	r := Runner.New(timeout)

	//绑定相应的执行任务到Runner去执行
	r.Add(createTask(), createTask(), createTask())

	//执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case Runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case Runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)

		}
	}
	log.Println("Process ended.")
}

//createTask返回一个根据id休眠执行秒数的任务函数
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
