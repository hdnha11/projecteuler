package main

import (
	"fmt"
	"math"
)

func main() {
	var k float64 = 20
	var n float64 = 1

	primes := primesTo(k)

	for _, p := range primes {
		var a float64 = 1
		if p*p <= k {
			// a <= logp(k) => a <= log10(k) / log10(p) => a = floor(log10(k) / log10(p))
			a = math.Floor(math.Log10(k) / math.Log10(p))
		}
		n *= math.Pow(p, a)
	}

	fmt.Println("Primes:", primes)
	fmt.Printf("Found: %.0f\n", n)
}

func primesTo(n float64) []float64 {
	primes := make([]float64, 0)

	if n >= 2 {
		primes = append(primes, 2)
	}
	if n >= 3 {
		primes = append(primes, 3)
	}

	var i float64 = 5
	for {
		if i > n {
			break
		}
		primes = append(primes, i)

		if i+2 > n {
			break
		}
		primes = append(primes, i+2)

		i += 6
	}

	return primes
}

// 2520 is the smallest number that can be divided by each of the numbers
// from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of
// the numbers from 1 to 20?

// NOTE
// 2
// 3
// 4 = 2.2
// 5
// 6 = 2.3
// 7
// 8 = 2.2.2
// 9 = 3.3
// 10 = 2.5 -> 2^3.3^2.5.7 = 2520
// 11
// 12 = 2.2.3
// 13
// 14 = 2.7
// 15 = 3.5
// 16 = 2.2.2.2
// 17
// 18 = 2.3.3
// 19
// 20 = 2.2.5 -> 2^4.3^2.5.7.11.13.17.19 = 232792560
// => find a[i] which p[i]^a[i] <= k (p is prime && p <= k)
//    n = n * p[i]^a[i]
