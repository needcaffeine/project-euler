/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"
import "math"

const n int64 = 600851475143

func isPrime(x int64) bool {
    var i int64 = 2;

    // 1 and 0 are not prime, neither are negative numbers.
    if x <= 1 {
        return false
    }

    // A number is a prime if it is divisible only by 1 and itself.
    for ; i*i <= x; i++ {
        if x % i == 0 {
            return false
        }
    }
    return true
}

func main() {
    // The largest factor of the number is going to be its square root.
    var i int64 = int64(math.Sqrt(float64(n)))

    for ; i > 1; i-- {
        if n % i == 0 && isPrime(i) {
            fmt.Printf("Largest prime factor of %d = %d\n", n, i)
            return
        }
    }

    // If we got here, n is prime.
    fmt.Printf("Largest prime factor of %d = %d\n", n, n)
}
