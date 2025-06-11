package main

import (
	"calc/calc"
	"fmt"
)

func main(){
	a := 2
	b := 2

	fmt.Printf("Add: %d\n", calc.Add(a, b))
	fmt.Printf("Sub: %d\n", calc.Sub(a, b))
	fmt.Printf("Mul: %d\n", calc.Mul(a, b))
	
	res, err := calc.Div(a, b)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("Div: %d\n", res)
	}
}