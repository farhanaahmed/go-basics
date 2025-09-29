package main

import (
	"fmt"
	"slices"
)

func main() {
	//helloWorld()
	//testForLoop()
	//arrays()
	//slicesExample()
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

	// set value at index 4
	a[4] = 100
	fmt.Println("set:", a)

	// get value at index 4
	fmt.Println("get:", a[4])

	// get length of array
	fmt.Println("len:", len(a))

	// declare and initialize array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// declare and initialize array without said length
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//If you specify the index with :, the elements in between will be zeroed.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// 2D array : i -> row, j -> column
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// initialize 2D array in one line,
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}

func slicesExample() {
	//Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	var s []string
	//An uninitialized slice equals to nil and has length 0.
	fmt.Println("uninitialized:", s, s == nil, len(s) == 0)
	// To create a slice with non-zero length, use the builtin make.
	//By default a new slice’s capacity is equal to its length.
	s = make([]string, 3)
	fmt.Println("using make:", s, "len:", len(s), "cap:", cap(s))
	//if we know the slice is going to grow ahead of time, it’s possible to pass a capacity explicitly as an additional parameter to make.
	// The capacity can be greater than or equal to the length.
	ss := make([]string, 3, 5)
	fmt.Println("using make:", ss, "len:", len(ss), "cap:", cap(ss))
	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	// len returns the length of the slice as expected.
	fmt.Println("len:", len(s))
	// append appends elements to the end of the slice.
	// append returns the new slice and modifies the original one.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	// supports a “slice” operator with the syntax slice[low:high].
	// The returned slice will contain all the elements from index low up to but not including index high.
	l := s[2:5]
	fmt.Println("low-high:", l)
	// This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("low-high:", l)
	//  This slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("low-high", l)
	// declare and initialize slice in one line
	t := []string{"g", "h", "i"}
	fmt.Println("t", t)
	// check if two slices are equal
	t2 := []string{"g", "h", "i"}
	fmt.Println("t2", t2)
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}
	// 2D slice
	//The length of the inner slices can vary, unlike the multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
