package main

import (
	"fmt"
	etcdserver "github.com/chijinxina/Go-Tempates/EtcdClientV3/etcd_server"
	"os"
	"os/signal"
	"path"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	pwd, _ := os.Getwd()
	dir := path.Join(pwd, "EtcdClientV3", "etcd_server")
	etcdServer := etcdserver.New(dir)
	if err := etcdServer.StartServer(); err != nil {
		fmt.Errorf("start etcd server failed: %s", err)
	} else {
		fmt.Println("start etcd server successfully")
	}

	defer etcdServer.Close()

	sig := <-sigChan
	fmt.Printf("receive %s signal, exiting...", sig)
}
