package main

import (
	"fmt"
	"init_func/cmd"
)

var age = test()

func test() int {
	fmt.Println("main test...")
	return 90
}

func init() {
	fmt.Println("main init...")
}

func main() {
	fmt.Println("main...")
	fmt.Printf("Age:%v,Name:%v\n", cmd.Age, cmd.Name)
}
