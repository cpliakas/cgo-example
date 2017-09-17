
build: include/HelloWorld.c
	gcc -c -o lib/HelloWorld.o include/HelloWorld.c
	gcc -dynamiclib -o lib/libHelloWorld.dylib lib/HelloWorld.o
	go build

clean:
	rm -f cgo-example lib/*
