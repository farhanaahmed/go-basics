package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	//helloWorld()
	//testForLoop()
	//arrays()
	//slicesExample()
	//mapsExample()
	//call plus function
	//res := plus(1, 2)
	//fmt.Println("1+2 =", res)
	// call plusPlus function
	//sum := plusPlus(1, 2, 3)
	//fmt.Println("1+2+3 =", sum)
	// call multipleReturn function
	//a, b := multipleReturn()
	//fmt.Println("a =", a, "b =", b)
	//fmt.Println(a)
	//fmt.Println(b)
	//If you only want a subset of the returned values, use the blank identifier _
	//_, b := multipleReturn()
	//fmt.Println(b)
	//a, _ := multipleReturn()
	//fmt.Println(a)
	// Variadic Functions
	sum(1, 2, 3)
	sum(1, 2, 3, 4)
	num := []int{1, 2, 3, 4, 5}
	sum(num...)
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
func mapsExample() {
	// Maps are Go’s built-in associative data type
	// Sometimes called hashes or dicts in other languages
	// declare and initialize a new map in the same line with this syntax
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	// To create an empty map, use the builtin make: make(map[key-type]val-type)
	m := make(map[string]int)
	fmt.Println("empty:", m)
	// Set key/value pairs using typical name[key] = val syntax
	m["key1"] = 7
	m["key2"] = 13
	// Printing a map with will show all of its key/value pairs
	fmt.Println("map:", m)
	// Get a value for a key with name[key]
	value1 := m["key1"]
	fmt.Println("value1:", value1)
	// If the key doesn’t exist, the zero value of the value type is returned
	value3 := m["key3"]
	fmt.Println("value3:", value3)
	// len returns the number of key/value pairs when called on a map
	fmt.Println("len:", len(m))
	// Delete a key/value pair with delete
	delete(m, "key2")
	fmt.Println("map:", m)
	// remove all key/value pairs from a map
	clear(m)
	fmt.Println("map:", m)
	// check if a key exists in a map
	// if key exists, prs will be true, otherwise false
	// Here we didn’t need the value itself, so we ignored it with the blank identifier _
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	// without the blank identifier, we will get the value of the key
	prs2 := m["k1"]
	fmt.Println("prs2:", prs2)

	// check if a key exists in a map
	// val, exists are only available inside the if statement
	if val, exists := m["a2"]; exists {
		fmt.Printf("Key 'a2' exists with value: %d\n", val)
	} else {
		fmt.Println("Key 'a2' does not exist in the map")
	}

	// check if two maps are equal
	a := map[string]int{"one": 1, "two": 2}
	b := map[string]int{"one": 1, "two": 2}
	if maps.Equal(a, b) {
		fmt.Println("a == b")
	}
}

// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.

func plus(a int, b int) int {
	return a + b
}

// omit the type name for the like-typed parameters up to the final parameter that declares the type.

func plusPlus(a, b, c int) int {
	return a + b + c
}

// multiple return values
// Go has built-in support for multiple return values
// The return statement can return multiple values
// For example, to return both result and error values from a function.

func multipleReturn() (int, int) {
	return 1, 2
}

// Variadic Functions
// They can be called with any number of trailing arguments
// A function that will take an arbitrary number of ints as arguments.

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
