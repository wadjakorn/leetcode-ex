package main

import (
	"testing"
)

func TestFunc(t *testing.T) {
	for _, testcase := range []struct {
		name   string
		input  []int
		x      int
		expect []int
	}{
		{
			name:   "case 1",
			input:  []int{0, -1, 2, -3, 1},
			x:      -2,
			expect: []int{-3, 1},
		},
		{
			name:   "case 2",
			input:  []int{1, -2, 1, 0, 5},
			x:      0,
			expect: []int{},
		},
		{
			name:   "case 3",
			input:  []int{2, -2, 3, 6, 5},
			x:      4,
			expect: []int{-2, 6},
		},
	} {
		actualList := findEQSumPair(testcase.input, testcase.x)

		countMatch := 0
		for _, eachExpect := range testcase.expect {
			for _, eachActual := range actualList {
				if eachActual == eachExpect {
					countMatch++
				}
			}
		}

		if countMatch != len(testcase.expect) {
			t.Errorf("case %s: failed, expected \"%v\" but got \"%v\"", testcase.name, testcase.expect, actualList)
		} else {
			t.Logf("case %s: success, expected \"%v\" and got \"%v\"", testcase.name, testcase.expect, actualList)
		}
	}
}
