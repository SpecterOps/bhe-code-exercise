package sieve

import (
	"fmt"
	"math"
)

type Sieve interface {
	NthPrime(n int64) int64
}

type Prime struct{}

func NewSieve() Sieve {
	return &Prime{}
}

// method to satisfy interface
func (p *Prime) NthPrime(n int64) int64 {
	// checks if n == 0 || n == 1 since i calculate a limit with multiplication
	if n == 0 {
		return int64(2);
	} else if(n == 1) {
		return int64(3);
	} else if (n < 0) {
		fmt.Println("Number must be a positive number."); // error for negative n
		return -1;
	}

	// i actually found this on stackoverflow for finding the max size. 
	// i don't know much about logarithms other than it is like exponents.
	limit := int(float64(n) * (math.Log(float64(n)) + math.Log(math.Log(float64(n))))) + 3;
	// i orginally used this but it was inconsistent and wouldn't always work as n got bigger
	// 18 is arbitrary. limit is used for initiating array
	// limit := n*18;

	// initialize array and set each index to true
	nums := make([]bool, limit + 1)
	for i := range nums {
		nums[i] = true;
	}

	// square root is used because all numbers below sqrt catch all numbers up to the square root
	sqrt := int(math.Sqrt(float64(limit)));

	// loop through nums array.
	for i := 2; i <= sqrt; i++ {
		// if true then start marking other num j indexes as false
		// set j = i*i because all nums up to i^2 are already marked
		// j+=i because you are comparing multiples of i
		if nums[i] {
			for j := i*i; j <= limit; j += i {
				nums[j] = false;
			}
		}
	}

	// count is used for keeping track of prime num count
	count := int64(0); 

	// iterate over array and once count == n return nth prime
	for i := 2; i <= limit || count == n; i++ {
		if nums[i] {
			if(count == n) {
				return int64(i);
			}
			count++;
		}
	}

	// this is for if prime isn't found at index in last for loop.
	fmt.Println("Something went wrong.");
	return -1;
}