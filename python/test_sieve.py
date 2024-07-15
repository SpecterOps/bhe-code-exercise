class Sieve:
    def __init__(self):
        self.prime_numbers = [2]  # Initialize with the first prime
        self.limit = 3  # Start with a small limit
        self.sieve = []

    def expand_sieve(self, new_limit):
        """Expand the sieve to a new limit."""
        self.sieve = [True] * new_limit
        for prime in self.prime_numbers:
            for i in range(prime * prime, new_limit, prime):
                self.sieve[i] = False
        
        for i in range(self.limit, new_limit):
            if self.sieve[i]:
                self.prime_numbers.append(i)
        
        self.limit = new_limit

    def nth_prime(self, n: int) -> int:
        """Return the nth prime number (0-indexed)."""
        while n >= len(self.prime_numbers):
            self.expand_sieve(min(self.limit * 2, 10**8))  # Double the limit, up to a maximum
        return self.prime_numbers[n]

def get_nth_prime(n: int) -> int:
    """API function to get the nth prime number."""
    sieve = Sieve()
    return sieve.nth_prime(n)
