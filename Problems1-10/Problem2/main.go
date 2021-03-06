package main

import "fmt"

/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/
func main () {
	sumEvens := fibonacci(0,1,4e6)
	fmt.Printf("Sum of even fibonnaci numbers: %v", sumEvens)
}

func fibonacci(start1, start2, limit int) int {
	if start1 > limit || start2 > limit {
		return 0
	}
	fmt.Println(start1+start2)
	sum := start1 + start2
	if sum % 2 == 0 {
		return fibonacci(start2, sum, limit) + sum
	} else {
		return fibonacci(start2, sum, limit)
	}
}