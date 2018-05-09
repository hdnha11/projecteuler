package main

import "fmt"

func main() {
	target := 999

	fmt.Println(sum(target))
}

func sum(t int) int {
	return sumDevisibleBy(3, t) + sumDevisibleBy(5, t) - sumDevisibleBy(15, t)
}

func sumDevisibleBy(n, t int) int {
	p := t / n

	return n * (p * (p + 1)) / 2
}

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
