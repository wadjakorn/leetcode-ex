package main

func findEQSumPair(input []int, x int) []int {
	total := len(input) - 1
	for outer := 0; outer <= total; outer++ {
		for inner := 0; inner <= total && input[outer] != input[inner]; inner++ {
			if input[outer]+input[inner] == x {
				return []int{input[outer], input[inner]}
			}
		}
	}
	return []int{}
}
