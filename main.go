package main

// #cgo CFLAGS: -Iinclude
// #cgo LDFLAGS: -Llib -lHelloWorld
// #include "HelloWorld.h"
import "C"

import "fmt"

func main() {
	helloWorld()
}

func helloWorld() {
	ptr := C.hello_world()
	fmt.Println(C.GoString(ptr))
}
