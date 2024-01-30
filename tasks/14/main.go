package main

import (
	"fmt"
	"reflect"
)

type esketit struct{}

type opa interface {
	esk()
}

func (*esketit) esk() { fmt.Println("эшкерее") }

func main() {
	ch := make(chan int)
	esk := esketit{}
	o := opa(&esk)
	arr := []any{5, 64.89, "string", true, ch, esk, o}

	for _, arg := range arr {
		switch arg.(type) {
		case string:
			fmt.Printf("String: %s\n", arg)
		case int:
			fmt.Printf("Int32: %d\n", arg)
		case float64:
			fmt.Printf("float64: %f\n", arg)
		case bool:
			fmt.Printf("Bool: %b\n", arg)
		default:
			fmt.Println("unexpected type:", reflect.TypeOf(arg))
		}
	}
}
