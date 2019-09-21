package main

import "fmt"

type Test struct {
	TestField bool
}

func main() {
	test := Test{}
	fmt.Println(test.TestField)
}
