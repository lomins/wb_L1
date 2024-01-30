package main

import "fmt"

func main() {
	var a Action
	a.Walk()
}

type Human struct{}

type Action struct {
	h Human
}

func (a *Action) Walk() {
	fmt.Println("Action walk")
	a.h.Walk()
}

func (h *Human) Walk() {
	fmt.Println("Human walk")
}
