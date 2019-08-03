package main

import (
	"testing"
)

func Test_mySum(t *testing.T) {
	type test struct {
		name   string
		data   []int
		answer int
	}

	tests := []test{
		{"2+3=5", []int{2, 3}, 5},
		{"4+7=11", []int{4, 7}, 11},
		{"5+9=14", []int{5, 9}, 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySum(tt.data...); got != tt.answer {
				t.Errorf("mySum() = %v, want %v", got, tt.answer)
			}
		})
	}
}
