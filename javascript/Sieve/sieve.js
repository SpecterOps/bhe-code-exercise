class Sieve {
  NthPrime(n) {
    let potentialPrime = 2;
    let numbers = [];

    //loop that determines how big the numbers array should be
    for (let i = 0; i <= n*2 ; i++){

      // 2 is the first prime and is the only even prime; it's a special case
      if(potentialPrime == 2){
        numbers.push(potentialPrime);
      // 2 and 3 are the only adjacent primes
        potentialPrime = potentialPrime + 1;
      } else if (potentialPrime == 3){
        numbers.push(potentialPrime);
      // prime numbers will never be even, so to only test odd numbers, we iterate by 2
        potentialPrime = potentialPrime+2;
      }
      else{
      // all prime numbers can be expressed in the following equations - by using these equations, we're filtering out more odd composite numbers than we would otherwise
        potentialPrime = ((6*(i-1)) - 1);
        numbers.push(potentialPrime);
        potentialPrime = ((6*(i-1)) + 1);
        numbers.push(potentialPrime);
      }
        } 
      //since the above equations sometimes include composite numbers, we can filter to return only primes
        function isPrime(potentialPrime){
        //this loop takes the potential prime and divides it by all of the numbers that came before it - since it involves iterating through the entire array, it doesn't scale well
          for(let j = 2; potentialPrime > j; j++){
            if(potentialPrime % j === 0){
              return false;
            }
          }
          return potentialPrime;
        }

        let primeNumbers = numbers.filter(isPrime);
        return primeNumbers[n];
        
  }

}

module.exports = Sieve;