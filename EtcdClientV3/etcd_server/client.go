package etcd_server

import (
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func CreateClient(host string, port int) (*clientv3.Client, error) {
	client, err := clientv3.New(
		clientv3.Config{
			Endpoints:   []string{fmt.Sprintf("%s:%d", host, port)},
			DialTimeout: 5 * time.Second,
		},
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}
