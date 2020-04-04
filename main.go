package main

import "syscall/js"

func add(this js.Value, i []js.Value) interface{} {
	// js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
	println(js.ValueOf(i[0].Int() + i[1].Int()).String())
	return js.ValueOf(i[0].Int() + i[1].Int())
}

func subtract(this js.Value, i []js.Value) interface{} {
	println(js.ValueOf(i[0].Int() - i[1].Int()).String())
	return js.ValueOf(i[0].Int() - i[1].Int())
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
