package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func test1() {
	b := bytes.NewBuffer([]byte{})
	b.WriteString("我aaebaaa")
	fmt.Println(string(b.Next(3)))
	buf := b.Bytes()
	fmt.Println(string(buf))
	//fmt.Println(bytes.IndexByte(buf, 'b'))
}

func test2() {
	un32 := math.Float32bits(0.65)
	bs := make([]byte, 4)

	binary.BigEndian.PutUint32(bs, un32)

	fmt.Println(bs)

	fmt.Println(math.Float32frombits(binary.BigEndian.Uint32(bs)))

	//fmt.Println(int(byte(bs>>24)))
	//fmt.Println(int(byte(bs>>16)))
	//fmt.Println(int(byte(bs>>8)))
	//fmt.Println(int(byte(bs)))
}

func test3() {
	var a int = 34294967296

	fmt.Println(int32(a))
}

func main() {
	test3()
}
