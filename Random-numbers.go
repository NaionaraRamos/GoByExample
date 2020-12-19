package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	p := fmt.Print
	r := rand.Intn
	f := rand.Float64

	p(r(100), ", ")
	p(r(100))
	p("\n")

	p(f())
	p("\n")

	p((f() * 5) + 5, ", ")
	p((f() * 5) + 5)
	p("\n")

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	p(r1.Intn(100), ", ")
	p(r1.Intn(100))
	p("\n")

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	p(r2.Intn(100), ", ")
	p(r2.Intn(100))
	p("\n")

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)

	p(r3.Intn(100), ", ")
	p(r3.Intn(100))

}