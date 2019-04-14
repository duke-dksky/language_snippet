package main

import "fmt"

type A struct {
	num int
}

func (a *A) help() {
	fmt.Printf("A num: %d\n", a.num)
}

type B struct {
	A
}

func (b *B) help() {
	fmt.Printf("B num: %d\n", b.num)
}

type C struct {
	a A
}

func (c *C) help() {
	fmt.Printf("C num: %d\n", c.a.num)
}

func main() {
	b := B{
		A: A{
			num: 1,
		},
	}
	b.help()
	b.A.help()

	c := C{
		a: A{
			num: 2,
		},
	}
	c.help()
	c.a.help()
}
