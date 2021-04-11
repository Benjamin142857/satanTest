package protocol

import (
	"fmt"
	"testing"
)

func test1(b byte) {

}

func test2(i int32) {

}

func TestTypeValue(t *testing.T) {
	//t.Run("normal", func(t *testing.T) {
	//	a := 1
	//	switch a {
	//	case Byte:
	//		fmt.Println("Byte")
	//	case Int:
	//		fmt.Println("Int")
	//	default:
	//		fmt.Println("none")
	//	}
	//})
	//a := 1
	//switch a {
	//case Byte:
	//	fmt.Println("Byte")
	//case Int:
	//	fmt.Println("Int")
	//default:
	//	fmt.Println(a)
	//	fmt.Println("none")
	//}

	//for i:=0; i<1000; i++ {
	//	for j:=1; j<1000; j++ {
	//		if res:=i*j + i/j; res==1000 {
	//			fmt.Println(res)
	//		}
	//	}
	//}
	//
	qq := []string{"a", "b"}
	var tt interface{} = 1
	_, ok := tt.(int)
	if !ok {
		fmt.Println("type error")
	}
	fmt.Println(qq[1])
}
