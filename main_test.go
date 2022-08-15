package main

import (
	"testing"
)

type testCase struct {
	expected int
	result int
}
func TestFactorial(t *testing.T) {
	//refactoring tests
	testCases := []testCase {
		{expected: 1, result: factorial(1)},
		{expected: 2, result: factorial(2)},
		{expected: 3 * 2 * 1, result: factorial(3)},
		{expected: 5 * 4 * 3 * 2 * 1, result: factorial(5)},

	}
	for i := range testCases {
		if testCases[i].expected != testCases[i].result{
			t.Errorf("failed expected: %d got %d", testCases[i].expected, testCases[i].result)
		}
	}
/*	expected := 1
	result := factorial(1)

	if expected != result {
		//failed here it is going to keep on to the next test
		t.Errorf("failed expected: %d got %d", expected, result)
	}

	expected = 2
	result = factorial(2)

	if expected != result {
		t.Errorf("failed expected: %d got %d", expected, result)
	}

	expected = 3 * 2 * 1
	result = factorial(3)

	if expected != result {
		t.Errorf("failed expected: %d got %d", expected, result)
	}*/
}