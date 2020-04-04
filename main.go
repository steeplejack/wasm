package main

import (
	"strconv"
	"syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, err := strconv.Atoi(value1)
	if err != nil {
		panic(err)
	}

	int2, err := strconv.Atoi(value2)
	if err != nil {
		panic(err)
	}

	println(int1 + int2)
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
	return js.ValueOf(int1 + int2)
}

func subtract(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, err := strconv.Atoi(value1)
	if err != nil {
		panic(err)
	}

	int2, err := strconv.Atoi(value2)
	if err != nil {
		panic(err)
	}

	println(int1 - int2)
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)
	return js.ValueOf(int1 - int2)
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
