package main

import (
	"go-practice-playground/basic/container"
	contextDemo "go-practice-playground/basic/context-demo"
	"go-practice-playground/basic/encrypt"
	gotour "go-practice-playground/basic/go-tour"
	"go-practice-playground/basic/goruntine"
)

func main() {
	// test simple examples of go tour
	gotour.Handler.Run()

	// test context usage
	contextDemo.Handler.Run()

	// test goruntine usage
	goruntine.Handler.Run()

	// test encrypt usage
	encrypt.Handler.Run()

	// test container type (List, Heap, Ring)
	container.RunHeap()
}
