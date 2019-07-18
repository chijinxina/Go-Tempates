package main

import (
	"context"
	"fmt"
	etcd "github.com/chijinxina/Go-Tempates/EtcdClientV3/etcd_server"
)

func main() {
	ctx := context.Background()
	etcdClient, err := etcd.CreateClient("10.224.16.114", 2379)
	if err != nil {
		fmt.Errorf("create etcd client failed: %s", err)
	}
	defer etcdClient.Close()

	if _, err := etcdClient.Put(ctx, "/test/chi1", "value111"); err != nil {
		fmt.Errorf("put {%s: %s} to etcd failed: %s", "test", "value1", err)
	} else {
		fmt.Println("put kv to etcd success: ")
	}

	if res, err := etcdClient.Get(ctx, "/chijinxin/dangxue"); err != nil {
		fmt.Errorf("get %v error", "test1")
	} else {
		fmt.Println("success: ", res)
	}

	select {}
}
