class Sieve:
    def __init__(self):
        self.prime_numbers = [2]  # Initialize with the first prime
        self.limit = 3  # Start with a small limit
        self.sieve = []

    def expand_sieve(self, new_limit):
        # Implementation of sieve expansion

    def nth_prime(self, n: int) -> int:
        while n >= len(self.prime_numbers):
            self.expand_sieve(min(self.limit * 2, 10**8))  # Double the limit, up to a maximum
        return self.prime_numbers[n]
