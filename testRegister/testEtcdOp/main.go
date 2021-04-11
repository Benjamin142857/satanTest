package main

import (
	"context"
	"fmt"
	etcd "go.etcd.io/etcd/clientv3"
	"time"
)
var etcdReqTimeout = 5*time.Second
var etcdServerPool = []string{
	"192.168.3.121:9092",
	//"benjamin142857.ticp.vip:56800",
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

func testPutLease(k, v string) {
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

	kv := etcd.NewKV(client)
	lease := etcd.NewLease(client)

	// 设置租约
	ctx1, cancel := context.WithTimeout(context.TODO(), etcdReqTimeout)
	leaseRsp, err := lease.Grant(ctx1, 5)
	if err != nil {
		fmt.Println("lease.Grant fail: ", err)
		return
	}
	cancel()
	fmt.Println("leaseId: ", leaseRsp.ID)

	// put
	ctx2, cancel := context.WithTimeout(context.TODO(), etcdReqTimeout)
	_, err = kv.Put(ctx2, k, v, etcd.WithLease(leaseRsp.ID))
	if err != nil {
		fmt.Println("kv.Put fail: ", err)
		return
	}

	// 定时续期租约
	for i:=0; i<10; i++ {
		ctx, cancel := context.WithTimeout(context.TODO(), etcdReqTimeout)
		keepRsp, err := lease.KeepAliveOnce(ctx, leaseRsp.ID)
		if err != nil {
			fmt.Println("leaseRsp.ID keepAlive fail: ", err)
			return
		}
		cancel()
		fmt.Printf("leaseId: %v, ttl: %v\n", keepRsp.ID, keepRsp.TTL)
		time.Sleep(2*time.Second)
	}
}

func testUint16(n uint16) {
	fmt.Println(n)
}

func WriteStProto() {

}

func main() {
	testUint16(123)
	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//testWatch(ctx, "goTest")
	//cancel()

	//testPutLease("goTest", "kkoouu")


	//for i:=0; i<20; i++ {
	//	go testPut("goTest", "111")
	//}
	//time.Sleep(10 * time.Second)
}


