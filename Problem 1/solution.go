package main

import "fmt"

func main() {

	var index, sum int = 0, 0

	for index < 1000 {
		if index%3 == 0 || index%5 == 0 {
			sum = sum + index
		}
		index = index + 1

	}
	fmt.Printf("Sum of all the multiples of 3 or 5 below 1000 is %v", sum)
}
