package main

import "testing"

type palindromeNumberParameter struct {
	arg1     int
	expected bool
}

var palindromeNumberParameters = []palindromeNumberParameter{
	palindromeNumberParameter{121, true},
	palindromeNumberParameter{12, false},
	palindromeNumberParameter{1, true},
	palindromeNumberParameter{0, true},
	palindromeNumberParameter{123321, true},
	palindromeNumberParameter{12123, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range palindromeNumberParameters {
		if output := isPalindrome(test.arg1); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}
