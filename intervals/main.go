package main

/*
 You are given a series of intervals and some of the intervals overlap.
 Reduce the list of intervals such that no intervals overlap with each other
 by joining them.
 Ex: [(2,4), (3,5), (6,8),(7,9)] => [(2,5), (6,9)]
*/

import (
	"fmt"
	"math/rand"
	"sort"
)

type intervals []struct {
	Start, End int
}

func main() {

	test3 := intervals{
		{2, 4},
		{6, 8},
		{1, 3},
		{5, 7},
	}

	fmt.Println(test3)

	sort.SliceStable(test3, func(i int, j int) bool {
		return test3[i].Start < test3[j].Start
	})

	fmt.Println(test3)
	test3 = quickSortIntervals(test3)

	finalResult := closeIntervals(test3)
	fmt.Println(finalResult)
}

// function that eliminates overlapping intervals
func closeIntervals(arr intervals) intervals {

	var squishedIntervals intervals
	squishedIntervals = append(squishedIntervals, arr[0])
	for i := 1; i < len(arr); i++ {
		if squishedIntervals[len(squishedIntervals)-1].End < arr[i].Start {
			squishedIntervals = append(squishedIntervals, arr[i])
		} else {
			squishedIntervals[len(squishedIntervals)-1].End = arr[i].End
		}
	}
	return squishedIntervals
}

// optional custom quickSort
func quickSortIntervals(arr intervals) intervals {
	if len(arr) < 2 {
		return arr
	}

	left := 0
	right := len(arr) - 1
	pivot := rand.Int() % len(arr)
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i].Start < arr[right].Start {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quickSortIntervals(arr[:left])
	quickSortIntervals(arr[left+1:])
	return arr
}

// OG quicksort
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left := 0
	right := len(arr) - 1
	pivot := rand.Int() % len(arr)
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	return arr
}
