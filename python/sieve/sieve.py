class Sieve:
    prime_numbers = []
    limit = 200000000
    arr = [0] * limit
    def __init__(self) -> None:
        for i in  range(2,self.limit):
            if self.arr[i]!=0:
                continue
            self.prime_numbers.append(i)
            # print(i)
            for j in range(i,self.limit):
                if i*j > self.limit-1:
                    break
                self.arr[i*j]+=1
    def nth_prime(self, n: int) -> int:
        return self.prime_numbers[n]
