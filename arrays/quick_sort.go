package arrays

import (
	"math/rand"
	"slices"
)

func QuickSort(input []int) []int {
	inputLength := len(input)

	if (inputLength <= 1) {
		return input
	}

	pivot := rand.Intn(inputLength)

	i := 0
	j :=  inputLength - 1

	for !(i == pivot && j == pivot) {
		if (i != pivot) {
			if (input[i] < input[pivot]) {
				i += 1
			} else {
				// save, delete input[i]
				swap := input[i]
				input = slices.Delete(input, i, i+1)
				// update indices, no need to update i as slice to the right shifts left and i points to next elem
				j -= 1
				pivot -= 1
				// insert input[i] at j+1
				input = slices.Insert(input, j+1, swap)
			}
		}
		if (j != pivot) {
			if (input[j] > input[pivot]) {
				j -= 1
			} else {
				// save, delete input[j]
				swap := input[j]
				input = slices.Delete(input, j, j+1)
				// insert input[j] at i
				input = slices.Insert(input, i, swap)
				// update indices
				pivot += 1
				i += 1
			}
		}
	}
	left := append(QuickSort(input[:pivot]), input[pivot])
	right := []int{}
	if (pivot != inputLength) {
		right = QuickSort(input[pivot + 1:])
	}
	return append(left, right...)
}