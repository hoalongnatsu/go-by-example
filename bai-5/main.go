package main

import (
	"fmt"
)

func main() {
	var number int
	var max_range int

	for {
		fmt.Print("Enter an Integer Number: ")
		fmt.Scan(&number)
		fmt.Print("Enter the range or end value: ")
		fmt.Scan(&max_range)

		for i := 1; i <= max_range; i++ {
			fmt.Printf("%d * %d = %d\n", number, i, number*i)
		}

		fmt.Println("============")
	}
}
