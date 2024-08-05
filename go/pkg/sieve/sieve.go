package sieve

type NthPrime struct{}

type Sieve interface {
	NthPrime(n int64) int64
}

// fill func marks the multiple of the prime number and skips previews marked multiples 
func fill(primes []bool, n int64) []bool {
	for p := int64(2); p*p <= n; p++ {
		if primes[p] {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}
	return primes
}

// markNumbers return a list of primenumber within low and high range 
func markNumbers(p []bool, n int64, l chan []int64) {
	primes := fill(p, n)
	var primeNumbers []int64
	for p := int64(2); p <= n; p++ {
		if primes[p] {
			primeNumbers = append(primeNumbers, p)
		}
	}

	l <- primeNumbers

	defer close(l)
}

// NthPrime returns the prime number 
func (p NthPrime) NthPrime(position int64) int64 {
	if position == 0 {
		return 2
	}

	n := position * position
	ch := make(chan []int64, n)
	primes := make([]bool, n+1)

	for i := int64(2); i <= n; i++ {
		primes[i] = true
	}

	go markNumbers(primes, n, ch)
	var primeNumbers []int64
	for nums := range ch {
		primeNumbers = append(primeNumbers, nums[position])
	}

	return primeNumbers[0]
}

func NewSieve() Sieve {
	return &NthPrime{}
}
