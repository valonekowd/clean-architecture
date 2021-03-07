package main

import "fmt"

type A interface{}

type B struct {
	a *A
}

func (b *B) Q() bool {
	return b.a == nil
}

func main() {
	var b *B

	fmt.Println(b.Q())
}
