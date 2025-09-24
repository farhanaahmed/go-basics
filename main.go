package main

import "fmt"

func main() {
	//helloWorld()
	//testForLoop()
	arrays()
	//slices()
}

func helloWorld() {
	fmt.Println("hello world")
}

func testForLoop() {
	for i := range 10 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func arrays() {
	var a [5]int
	fmt.Println("emp:", a)
}

func slices() {
	//Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	//An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninitialized:", s, s == nil, len(s) == 0)
}
