package main

import (
	"fmt"
	"os"
)

func testOsEnv() {
	for _, env := range os.Environ() {
		//if strings.HasPrefix(env, grace.InheritFdPrefix) {
			fmt.Printf("env %s\n", env)
		//}
	}
}

func main() {
	testOsEnv()
}