package main

import (
	"testing"
)

// TestSumExerciseNums calls SumExercise with [1, 2, 3, 4, 5] nums array
// checking for a valid sum
func TestSumExerciseNums(t *testing.T) {
	nums := []int64{1, 2, 3, 4, 5}
	expectedResult := int64(15)

	result := SumExercise(nums)

	if result != expectedResult {
		t.Errorf("\nInput: [1, 2, 3, 4, 5]\nResult: %d\nExpected result: %d\n", result, expectedResult)
	}
}

func TestSumExerciseEmpty(t *testing.T) {
	nums := []int64{}
	expectedResult := 0

	result := SumExercise(nums)

	if result != 0 {
		t.Errorf("\nInput: []\nResult: %d\nExpected result: %d\n", result, expectedResult)
	}
}
