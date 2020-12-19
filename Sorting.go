package main

import (
	"fmt"
	"sort"
)

//Sorting by Functions
type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
//Sorting by Functions

func main() {

	//Sort

	/* strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s) */

	//sort


	//Sorting by Functions

	fruits := []string{"pessego", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
	//Sorting by Functions
}