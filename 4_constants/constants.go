package main

import "fmt"

// we can declare/use variables & constants outside main function (except shorthand syntax variables)

// can be used
// const age = 20
// var name string = "randomName"

// can't be used here
// num := 20

func main() {
	// constans: once assigned can't be changed
	const name string = "go-lang"
	const age = 20
	// age = 30 // not allowed
	fmt.Println(name, age)

	// declaring and assigning to multiple constants simultaneously

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)

}
