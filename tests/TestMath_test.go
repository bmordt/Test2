package tests

import (
	"Test2/src/mathstuff"
	"testing"
)

func TestAverage(t *testing.T) {
	expectedResult := ((1.0 + 2.0 + 3.0) / (3.0))
	actualResult := mathstuff.Average([]float64{1.0, 2.0, 3.0})

	if expectedResult != actualResult {
		t.Errorf("Expected: %f got: %f", expectedResult, actualResult)
	} else {
		t.Logf("Expected: %f got: %f", expectedResult, actualResult)
	}
}

func TestMedian(t *testing.T) {
	expected1 := 5.0
	actual1 := mathstuff.Median([]float64{1.0, 10.1, 30.0, 2.0, 5.0})

	if expected1 != actual1 {
		t.Errorf("Expected: %f got: %f", expected1, actual1)
	} else {
		t.Logf("Expected: %f got: %f", expected1, actual1)
	}

	expected2 := 505.05
	actual2 := mathstuff.Median([]float64{1000.0, 10.1, 35550.0, 1111.0, 5.0, 1.0})

	if expected2 != actual2 {
		t.Errorf("Expected: %f got: %f", expected2, actual2)
	} else {
		t.Logf("Expected: %f got: %f", expected2, actual2)
	}
}

func TestMode(t *testing.T) {
	expected1 := []float64{1.0}
	actual1 := mathstuff.Mode([]float64{1.0, 10.1, 30.0, 2.0, 5.0, 1.0, 1.0, 1.0, 1.0})

	if !FloatArrayEquals(expected1, actual1) {
		t.Errorf("Expected: %f got: %f", expected1, actual1)
	} else {
		t.Logf("Expected: %f got: %f", expected1, actual1)
	}

	expected2 := []float64{1.0, 2.0}
	actual2 := mathstuff.Mode([]float64{1.0, 10.1, 30.0, 2.0, 2.0, 2.0, 2.0, 5.0, 1.0, 1.0, 1.0, 1.0, 2.0})

	if !FloatArrayEquals(expected2, actual2) {
		t.Errorf("Expected: %f got: %f", expected2, actual2)
	} else {
		t.Logf("Expected: %f got: %f", expected2, actual2)
	}
}

func FloatArrayEquals(a []float64, b []float64) bool {
	if len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}
}
