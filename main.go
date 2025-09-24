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
	// The type of elements and length are both part of the array’s type.
	// By default, an array is zero-valued.

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}

func slices() {
	//Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	//An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninitialized:", s, s == nil, len(s) == 0)
}
