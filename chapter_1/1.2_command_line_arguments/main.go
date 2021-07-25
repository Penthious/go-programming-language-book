package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	example1()
	example2()
	example3()
}


func example1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("example 1: ",s)
}

func example2() {
	for i, arg := range os.Args[1:] {
		fmt.Println("example 2: ", i, arg)
	}
}

func example3() {
	//@todo: page 8 Exercise 1.3 create benchmark for all 3 examples
	fmt.Println("example 3: ", strings.Join(os.Args[1:], " "))
}