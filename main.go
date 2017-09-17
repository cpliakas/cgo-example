package main

// #include "HelloWorld.h"
import "C"

import "fmt"

func main() {
	helloWorld()
}

// helloWorld calls the hello_world() function defined in the HelloWorld.c
// file, converts the character array that is returned to a native Go
// string, and writes it to stdout.
func helloWorld() {
	ptr := C.hello_world()
	fmt.Println(C.GoString(ptr))
}
