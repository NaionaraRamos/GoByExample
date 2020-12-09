package main

import (
	"fmt"
	"time"
)

func main(){
	/* i := 1
	for i < 3{
		fmt.Println(i)
		i = i + 1
	}

	for j := 6; j <= 9; j++{
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	} */

	i := 2
	fmt.Println("Escreva", i, "como")
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Fim de semana...")
	default:
		fmt.Println("Dia de semana...")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Antes do meio dia...")
	default:
		fmt.Println("Depois do meio-dia...")
	}

	oQueSou := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Bool")
		case int:
			fmt.Println("Int")
		default:
			fmt.Printf("Não dá pra saber tipo %T\n", t)
		}
	}

	oQueSou(true)
	oQueSou(1)
	oQueSou("tipo")
}