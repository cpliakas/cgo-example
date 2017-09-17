package main

// #include "HelloWorld.h"
import "C"

import "fmt"

func main() {
	helloWorld()
}

// helloWorld calls the hello_world() function defined in the HelloWorld.c
// file, converts the character array to a native Go string, and prints it.
func helloWorld() {
	ptr := C.hello_world()
	fmt.Println(C.GoString(ptr))
}
