package main

import (
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
