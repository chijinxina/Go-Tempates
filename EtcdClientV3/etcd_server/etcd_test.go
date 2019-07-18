package etcd_server

import (
	"context"
	"fmt"
	"os"
	"testing"
)

const (
	testKey   = "test"
	testValue = "test_test"
)

func TestEtcd(t *testing.T) {
	ctx := context.Background()
	dir, _ := os.Getwd()
	server := New(dir)
	err := server.StartServer()
	if err != nil {
		t.Fatalf("start etcd server failed: %s", err)
	} else {
		t.Log("start etcd server successfully")
	}

	client, err := CreateClient()
	if err != nil {
		t.Fatalf("create etcd client failed: %s", err)
	}
	defer client.Close()

	fmt.Println("PUT KV TO ETCD")
	if _, err := client.Put(ctx, testKey, testValue); err != nil {
		t.Fatalf("put {%s: %s} to etcd failed: %s", testKey, testValue, err)
	} else {
		t.Logf("put {%s: %s} to etcd successfully", testKey, testValue)
	}
	if resp, err := client.Get(ctx, testKey); err != nil {
		t.Fatalf("get %s value failed: %s", testKey, err)
	} else if resp.Count == 1 && string(resp.Kvs[0].Value) == testValue {
		t.Logf("get %s value successfully", testKey)
	} else {
		t.Fatalf("get %s value failed, resp: %+v", testKey, resp)
	}
	if err := server.Close(); err != nil {
		t.Fatalf("close etcd server failed: %s", err)
	} else {
		t.Log("etcd server has closed")
	}
}
