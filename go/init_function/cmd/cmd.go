package cmd

import (
	"fmt"
)

var (
	Age  int = test()
	Name string
)

func init() {
	fmt.Println("cmd init...")
	Age = 100
	Name = "XXX"
}

func test() int {
	fmt.Println("cmd test...")
	return 10
}
