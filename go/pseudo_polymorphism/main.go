package main

import (
	"fmt"
)

type state interface {
	monitor(m *machine)
}

type normal struct{}

func (*normal) monitor(m *machine) {
	fmt.Println("normal monitor")
	m.num++
}

type failover struct{}

func (*failover) monitor(m *machine) {
	fmt.Println("failover monitor")
	m.num++
}

type machine struct {
	state
	num int
}

func NewMachine() *machine {
	return &machine{
		state: new(normal),
	}
}

func main() {
	m := NewMachine()
	m.monitor(m)
	m.state = new(failover)
	m.monitor(m)
	fmt.Println(m.num)
}
