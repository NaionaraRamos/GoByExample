package main

import (
	"fmt"
	"strconv"
)

func main() {

	p := fmt.Println
	f := strconv.ParseFloat
	i := strconv.ParseInt
	u := strconv.ParseUint
	a := strconv.Atoi

	f1, _ := f("1.234", 64)
	p(f1)

	i1, _ := i("123", 0, 64)
	p(i1)

	i2, _ := i("0x1c8", 0, 64)
	p(i2)

	u1, _ := u("789", 0, 64)
	p(u1)

	a1, _ := a("135")
	p(a1)

	_, e := a("wat")
	p(e)
}