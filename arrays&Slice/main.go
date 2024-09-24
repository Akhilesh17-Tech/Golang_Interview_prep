package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello, World!")

	//creating an array with fixed size value type
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr := array

	array[0] = 10000
	arr[0] = 2000

	fmt.Println(array)
	fmt.Println(arr)

	// slice is reference type in go
	slice1 := []int{1, 2, 3, 4}
	slice2 := slice1
	slice1[0] = 100
	slice2[1] = 200

	fmt.Println(slice1)
	fmt.Println(slice2)

	// for loop range function
	for index, value := range array {
		fmt.Println(index, value)
	}

	// Declare and Initialize an Array with Default Values
	var array1 [5]int
	fmt.Println("array1:", reflect.TypeOf(array1).Kind(), array1)

	// Declare and Initialize an Array with Specific Values
	var array2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("array2:", array2)

	// Declare and Initialize an Array with Specific Values Using `...`
	var array3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("array3:", array3)

	// Short Variable Declaration with Array
	array4 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array4:", array4)

	// Partial Initialization
	var array5 = [5]int{1, 2}
	fmt.Println("array5:", array5)

	// Initialize Specific Elements
	var array6 = [5]int{1: 10, 3: 30}
	fmt.Println("array6:", array6)

	sl := make([]int, 5, 10)
	fmt.Println("sl:", reflect.TypeOf(sl).Kind(), sl)
	sl = append(sl, 10)
	fmt.Println("sl:", cap(sl), sl)

	var arrs = new([5]int)
	var arrs1 = make([]int, 5)
	fmt.Println(*arrs)
	fmt.Println(arrs1)
}
