package arrays

import "strconv"

func IntToString(input []int) []string {
	output := make([]string, len(input))
	for i, el := range input {
		output[i] = strconv.Itoa(el)
	}
	return output
}

func Int64ToString(input []int64) []string {
	output := make([]string, len(input))
	for i, el := range input {
		output[i] = strconv.FormatInt(el, 10)
	}
	return output
}
