package main

import (
	"reflect"
	"testing"
)

func Test_validInput(t *testing.T) {
	type args struct {
		inputs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "正常な場合",
			args: args{
				inputs: []string{"1", "2", "3", "0004"},
			},
			want: true,
		},
		{
			name: "inputに数値ではない文字が含まれる場合",
			args: args{
				inputs: []string{"1", "2", "3", "0004a"},
			},
			want: false,
		},
		{
			name: "inputの長さが1未満の場合",
			args: args{
				inputs: []string{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validInput(tt.args.inputs); got != tt.want {
				t.Errorf("validInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortInputs(t *testing.T) {
	type args struct {
		inputs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "数字が正常に並んでいることを確認",
			args: args{
				inputs: []string{"2", "1", "01", "1", "0010"},
			},
			want: []string{"01", "1", "1", "2", "0010"},
		},
		{
			name: "入力例２の値を確認",
			args: args{
				inputs: []string{
					"1111111111111111111111",
					"00011111111111111111111",
					"000000111111111111111111",
					"0000000001111111111111111",
					"00000000000011111111111111",
					"000000000000000111111111111",
				},
			},
			want: []string{
				"000000000000000111111111111",
				"00000000000011111111111111",
				"0000000001111111111111111",
				"000000111111111111111111",
				"00011111111111111111111",
				"1111111111111111111111",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortInputs(tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortInputs() = %v, want %v", got, tt.want)
			}
		})
	}
}
