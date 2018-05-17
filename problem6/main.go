package main

import "fmt"

func main() {
	n := 100

	sum := n * (n + 1) / 2
	sumOfSquares := n * (2*n + 1) * (n + 1) / 6
	squareOfSum := sum * sum

	fmt.Println(squareOfSum - sumOfSquares)
}

// The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385

// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)^2 = 55^2 = 3025

// Hence the difference between the sum of the squares of the first ten
// natural numbers and the square of the sum is 3025 − 385 = 2640.

// Find the difference between the sum of the squares of the first one
// hundred natural numbers and the square of the sum.

// NOTE
// Differences method

// Table
// n       0   1   2   3   4
// f(n)    0   1   5  14  30
// d1	       1   4   9  16
// d2	           3   5   7
// d3	               2   2

// => f(n) = a.n^3 + b.n^2 + c.n + d

// f(0) = 0  => d = 0
// f(1) = 1  => a + b + c = 1 (1)
// f(2) = 5  => 8a + 4b + 2c = 5 (2)
// f(3) = 14 => 27a + 9b + 3c = 14 (3)

// from (1) => c = 1 - a - b
// replace c to (2) => 8a + 4b + 2(1 - a - b) = 5
// <=> 6a + 2b = 3 => b = (3 - 6a)/2 = 3/2 - 3a
// =>  c = 1 - a - (3/2 - 3a) = -1/2 + 2a
// replace b, c to (3) => 27a + 9(3/2 - 3a) + 3(-1/2 + 2a) = 14
// <=> 6a = 2
// <=> a = 1/3
// =>  b = 3/2 - 3(1/3) = 3/2 - 1 = 1/2
// =>  c = 1 - 1/3 - 1/2 = (6 - 2 - 3)/6 = 1/6

// =>  f(n) = 1/3.n^3 + 1/2.n^2 + 1/6.n = n/6(2.n^2 + 3n + 1) = n/6(2n + 1)(n + 1)
