package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("最初に文字の数を入力してください: ")
	numInput, _ := reader.ReadString('\n')
	numInput = strings.TrimSpace(numInput)
	num, err := strconv.ParseUint(numInput, 10, 64) // 文字列を整数に変換
	if err != nil {
		fmt.Println("入力が不正です")
		return
	}
	var inputs []string
	for i := 0; uint64(i) < num; i++ {
		fmt.Printf("%d番目の文字を入力してください: ", i+1)
		strInput, _ := reader.ReadString('\n')
		strInput = strings.TrimSpace(strInput) // 改行を削除
		inputs = append(inputs, strInput)
	}
	if err := validInput(inputs); err != nil {
		fmt.Println(err)
		return
	}
	outputs := SortInputs(inputs)
	for _, output := range outputs {
		fmt.Println(output + "\n")
	}
}

func validInput(inputs []string) error {
	for _, input := range inputs {
		inputNum := new(big.Int)
		_, ok := inputNum.SetString(input, 10)
		if !ok {
			return errors.New("inputに数値ではない文字が含まれています: " + input)
		}
	}
	if len(inputs) < 1 || len(inputs) > 100000 {
		return errors.New("inputの長さが不正です")
	}
	sum := 0
	for _, input := range inputs {
		sum += len(input)
	}
	if sum > 100000 {
		return errors.New("inputの長さの合計が不正です")
	}
	return nil
}

func SortInputs(inputs []string) []string {
	sort.SliceStable(inputs, func(i, j int) bool {
		//inputsをintに変換して比較
		numI := new(big.Int)
		numJ := new(big.Int)
		numI, ok := numI.SetString(inputs[i], 10)
		if !ok {
			return false
		}
		numJ, ok = numJ.SetString(inputs[j], 10)
		if !ok {
			return false
		}

		cmp := numI.Cmp(numJ)
		if cmp == 0 {
			return len(inputs[i]) > len(inputs[j])
		}
		return cmp < 0
	})
	return inputs
}
