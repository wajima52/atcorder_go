package main

import (
	"sort"
	"strconv"
)

func main() {

}

func validInput(inputs []string) bool {
	for _, input := range inputs {
		if _, err := strconv.Atoi(input); err != nil {
			return false
		}
	}
	if len(inputs) < 1 || len(inputs) > 100000 {
		return false
	}
	sum := 0
	for _, input := range inputs {
		sum += len(input)
	}
	if sum > 100000 {
		return false
	}
	return true
}

func SortInputs(inputs []string) []string {
	sort.SliceStable(inputs, func(i, j int) bool {
		//inputsをintに変換して比較
		iInt, _ := strconv.Atoi(inputs[i])
		jInt, _ := strconv.Atoi(inputs[j])

		if iInt == jInt {
			return len(inputs[i]) > len(inputs[j])
		}
		return iInt < jInt
	})
	return inputs
}
