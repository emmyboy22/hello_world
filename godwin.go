package main

import "fmt"

func main2() {
	var name string
	var car string 
	var clothe string
	
	car = "camery"
	name = "emeka"
	clothe = "agbada"


	var age int
	var numKid int
	var shoes int

	age = 45
	numKid = 9
	shoes = 4

	fmt.Fprintf("")



	func addition(num4, num5 int) int {
		answer := num4 + num5
		return answer
	}
	
	func substraction(num1, num2 int) int  {
		answer := num1 - num2
		return answer
	}
	
	func divide(num1, num2 int) int  {
		answer := num1 / num2
		return answer
	}
	
	func main()  {
		answer := addition(60, 60)
		fmt.Println(answer)
	
		anything := substraction(100,50)
		fmt.Println(anything)
	
		anything = divide(200, 100)
		fmt.Println(anything)
	
	
	}
	
	
	

}