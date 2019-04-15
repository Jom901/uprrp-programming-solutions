package main

/*
Given a list of integers from 1 to n that's missing a number, find
the missing number. Assume that the list may not be in order.
*/
import "fmt"

func main() {
	listOfNumbers := []int{1, 2, 3, 5, 4, 6, 8}
	fmt.Println(findMissingNumber(listOfNumbers))
}

func findMissingNumber(listOfNumbers []int) int {
	sum := 0
	n := (len(listOfNumbers) + 1)
	for x := 0; x < n-1; x++ {
		sum += listOfNumbers[x]
	}
	actualSum := (n * (n + 1)) / 2
	return actualSum - sum
}
