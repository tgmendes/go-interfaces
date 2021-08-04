package main

import "fmt"

type Barker interface {
	Bark() string
}

type TailWagger interface {
	WagTail() bool
}

type WagBarker interface {
	Barker
	TailWagger
}

type Cat struct{}

func (c Cat) WagTail() bool {
	return false
}

type Dog struct{}

func (d Dog) Bark() string {
	return "woof"
}

func (d Dog) WagTail() bool {
	return true
}

func WagAndBark(wb WagBarker) {
	fmt.Println(wb.Bark())
	Wag(wb)
}

func Wag(w TailWagger) {
	if w.WagTail() {
		fmt.Println("I wagged my tail!")
		return
	}
	fmt.Println("I did not wag my tail")
}

func main() {
	d := Dog{}
	c := Cat{}

	WagAndBark(d)
	// WagAndBark(c) -> this doesn't work
	Wag(c)
}
