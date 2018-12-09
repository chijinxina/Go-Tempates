package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

//实现简单的curl功能

//main程序入口
func main() {
	//r是http响应， r.body实现了io.Reader
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	//创建文件来保存响应内容
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	//使用MultiWriter，这样就可以实现同时向文件和输出设备进行写操作
	dest := io.MultiWriter(os.Stdout, file)

	//读出响应内容，并写入到两个目的地
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
