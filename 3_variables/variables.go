package main

import "fmt"

func main() {
	var name string = "ram"
	fmt.Println(name)

	// go lang will infer type of this variable automatically
	var name2 = "go lang"
	fmt.Println(name2)

	var isAdult = true
	fmt.Println(isAdult)

	var age int = 30
	fmt.Println(age)

	// shorthand syntax
	name3 := "raghu"
	fmt.Println(name3)

	// we can't always use shorthand for ex. in this case
	var str string
	str = "go lang"
	fmt.Println(str)

	var price1 float32 = 59.1
	var price2 = 50.3
	price3 := 40.4

	fmt.Println(price1, price2, price3)
}
