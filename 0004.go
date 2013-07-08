/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91  99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import "fmt"
import "strconv"

func isPalindrome(x int) bool {
    firstHalf := strconv.Itoa(x / 1000)
    secondHalf := strconv.Itoa(x - ((x / 1000) * 1000))
    runes := []rune(secondHalf)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    secondHalf = string(runes)

    if firstHalf == secondHalf {
        return true
    }
    return false
}

func main() {
    i := 998001
    for ; i >= 111*111 ; i-- {
        if isPalindrome(i) {
            for j := 999; j >= 111; j-- {
                if (i / j < 1000 && i % j == 0) {
                    fmt.Printf("The largest palindrome made from the product of two 3-digit numbers (%d x %d) is %d\n", j, i/j, i)
                    return
                }
            }
        }
    }
}
