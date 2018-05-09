package main

import "fmt"

func main() {
	n := 600851475143
	fmt.Println(primeFactorization(n))
	fmt.Println(isPrime(6857))
}

func primeFactorization(n int) []int {
	r := []int{}

	if n <= 1 {
		return r
	}

	if n <= 3 {
		return append(r, n)
	}

	for n%2 == 0 {
		r = append(r, 2)
		n /= 2
	}

	for n%3 == 0 {
		r = append(r, 3)
		n /= 3
	}

	i := 5

	for n > 1 {
		for n%i == 0 {
			r = append(r, i)
			n /= i
		}

		for n%(i+2) == 0 {
			r = append(r, i+2)
			n /= i + 2
		}

		i += 6
	}

	return r
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5

	for i*i < n {
		// i = 6k - 1 (5 = 6 * 1 - 1, k = 1)
		// i + 2 = 6k - 1 + 2 = 6k + 1
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
