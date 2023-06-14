package main

import (
	"fmt"
)

func calculator(num3 float32, sign string, num4 float32)float32{
	var answer float32
	if sign == "+" {
		answer = num3 + num4
	}else if sign == "-" {
		answer = num3 - num4
	}else if sign == "/" {
		answer = num3 / num4
	}else if sign == "*" {
		answer = num3 * num4
	}else {
		fmt.Println("please input a valid operator")
	}
	return answer
}

func main()  {
	result := calculator(100, "*", 10)
	fmt.Println(result)

	result = calculator(100, "/", 20)
	fmt.Println(result)

	result = calculator(50, "+", 40)
	fmt.Println(result)

	result = calculator(100, "-", 80)
	fmt.Println(result)
}
