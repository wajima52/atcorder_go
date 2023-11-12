package main

import "strconv"

func main() {

}

func validInput(inputs []string) bool {
	//	inputが全て数値でない場合はfalseを返す
	for _, input := range inputs {
		if _, err := strconv.Atoi(input); err != nil {
			return false
		}
	}
	//	inputsの長さは1以上100000以下である
	if len(inputs) < 1 || len(inputs) > 100000 {
		return false
	}
	//	inputsの長さの合計は100000以下である
	sum := 0
	for _, input := range inputs {
		sum += len(input)
	}
	if sum > 100000 {
		return false
	}
	return true
}
