package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

var (
	proto = "tcp"
	addr = "127.0.0.1:10086"
	InheritFdPrefix = "LISTEN_FD_INHERIT"

	allListenFds *sync.Map
)

func test1() {
	allListenFds = &sync.Map{}

	key := fmt.Sprintf("%s_%s_%s", InheritFdPrefix, proto, addr)
	val := os.Getenv(key)
	fmt.Println(val)
	//for val != "" {
	//	fd, err := strconv.Atoi(val)
	//
		//fmt.Println(fd)
		//fmt.Println(err)
		//break
		//if err != nil {
		//	break
		//}
		//file := os.NewFile(uintptr(fd), "listener")
		//ln, err := net.FileListener(file)
		//if err != nil {
		//	_ = file.Close()
		//	break
		//}
		//allListenFds.Store(key, ln)
		//fmt.Println("in for")
		//fmt.Println(ln)
	//}
	// not inherit, create new
	ln, err := net.Listen(proto, addr)
	if err != nil {
		fmt.Println("net.Listen fail: ", err)
		return
	}
	defer func() { _ = ln.Close() }()


	allListenFds.Store(key, ln)
	fmt.Println(allListenFds)
	fmt.Println("out for")
	fmt.Println(ln)
}

func main() {
	test1()
}
