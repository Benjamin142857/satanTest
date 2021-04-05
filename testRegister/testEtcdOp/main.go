package main

import (
	"context"
	"fmt"
	"time"

	etcd "go.etcd.io/etcd/clientv3"
)

var etcdServerPool []string = []string{
	"benjamin142857.ticp.vip:56800",
}


func testPut(k, v string) {
	client, err := etcd.New(etcd.Config{
		Endpoints:   etcdServerPool,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("client conn fail: ", err)
		return
	}
	defer func(){ _ = client.Close() }()
	fmt.Println("client conn succ.")

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	put, err := client.Put(ctx, k, v)
	cancel()
	if err != nil {
		fmt.Println("put fail: ", err)
		return
	}
	fmt.Println(put)
}

func testGet(k string) (vS []string) {
	client, err := etcd.New(etcd.Config{
		Endpoints:   etcdServerPool,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("client conn fail: ", err)
		return
	}
	defer func(){ _ = client.Close() }()
	fmt.Println("client conn succ.")

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	get, err := client.Get(ctx, k)
	cancel()
	if err != nil {
		fmt.Println("put fail: ", err)
		return
	}
	for _, kvItem := range get.Kvs {
		vS = append(vS, string(kvItem.Value))
	}
	return vS
}

func testDel(k string) {
	client, err := etcd.New(etcd.Config{
		Endpoints:   etcdServerPool,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("client conn fail: ", err)
		return
	}
	defer func(){ _ = client.Close() }()
	fmt.Println("client conn succ.")

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	response, err := client.Delete(ctx, k)
	cancel()
	if err != nil {
		fmt.Println("put fail: ", err)
		return
	}
	fmt.Println(response.Deleted)
}

func testWatch(ctx context.Context, k string) {
	client, err := etcd.New(etcd.Config{
		Endpoints:   etcdServerPool,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("client conn fail: ", err)
		return
	}
	defer func(){ _ = client.Close() }()
	fmt.Println("client conn succ.")

	wChan := client.Watch(ctx, k)
	for {
		select {
		case <- ctx.Done():
			fmt.Println("ctx timeout")
			return
		case wResp := <- wChan:
			for idx, ev := range wResp.Events {
				fmt.Printf("[%v] ev.Type=%v, ev.Key=%v, ev.Value=%v\n", idx, ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	testWatch(ctx, "goTest")
	cancel()



	//for i:=0; i<20; i++ {
	//	go testPut("goTest", "111")
	//}
	//time.Sleep(10 * time.Second)
}


