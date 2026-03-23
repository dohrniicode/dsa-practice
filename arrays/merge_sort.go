package arrays

func MergeSort(input []int) []int {
	inputLength := len(input)

	if inputLength <= 1 {
		return input
	}

	midpoint := inputLength / 2
	firstHalf := input[0:midpoint]
	firstHalf = MergeSort(firstHalf)

	secondHalf := input[midpoint:]
	secondHalf = MergeSort(secondHalf)

	firstHalfPtr := 0
	secondHalfPtr := 0
	var result []int
	for firstHalfPtr < len(firstHalf) && secondHalfPtr < len(secondHalf) {
		if firstHalf[firstHalfPtr] < secondHalf[secondHalfPtr] {
			result = append(result, firstHalf[firstHalfPtr])
			firstHalfPtr += 1
		} else {
			result = append(result, secondHalf[secondHalfPtr])
			secondHalfPtr += 1
		}
	}
	if firstHalfPtr < len(firstHalf) {
		result = append(result, firstHalf[firstHalfPtr:]...)
	} else if secondHalfPtr < len(secondHalf) {
		result = append(result, secondHalf[secondHalfPtr:]...)
	}

	return result
}
