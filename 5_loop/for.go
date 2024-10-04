package main

import "fmt"

// for -> only construct in go for looping

func main() {
	// while loop
	i := 1
	for i <= 10 {
		// fmt.Print(i, " ")
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	// for loop
	for i := 0; i < 3; i++ {
		// fmt.Println(i)
	}

	// break and continue

	for i := 1; i <= 15; i++ {
		if i%2 != 0 {
			continue
		} else if i == 12 {
			break
		}
		// fmt.Print(i, " ")
	}

	// 1.22 range
	for i := range 10 {
		fmt.Println(i)
	}
}
