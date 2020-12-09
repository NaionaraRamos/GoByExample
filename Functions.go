package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

//Multiple Return Values
func vals() (int, int) {
	return 4, 8
}

//Variadic Functions
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}


//Closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//Recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func main(){
	fmt.Println(plus(1, 2))
	fmt.Println(plusPlus(1, 2, 3))

	//Multiple Return Values
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
	//Multiple Return Values

	//Variadic Functions
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
	//Variadic Functions


	//Closures
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	//Closures


	//Recursion
	fmt.Println(fact(6))
	//Recursion
}