package main

import "fmt"

func main() {
	from := 100
	to := 999
	found := -1
	for i := to; i > from-1; i-- {
		for j := to; j > i-1; j-- {
			p := i * j
			if p < found {
				break
			}
			if isPalindromic(p) {
				found = p
			}
		}
	}

	if found != -1 {
		fmt.Println("Found", found)
	} else {
		fmt.Printf("No such palindromic in this range [%d, %d]", from, to)
	}
}

func isPalindromic(n int) bool {
	return reverseNumber(n) == n
}

func reverseNumber(n int) int {
	var r int
	for n > 0 {
		r = (r * 10) + (n % 10)
		n /= 10
	}

	return r
}

// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
