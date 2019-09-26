package tests

import (
	"Test2/src/mathstuff"
	"fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	expectedResult := ((1.0 + 2.0 + 3.0) / (3.0))
	actualResult := mathstuff.Average([]float64{1.0, 2.0, 3.0})

	if expectedResult != actualResult {
		t.Errorf("Expected: %.2f got: %.2f", expectedResult, actualResult)
	} else {
		t.Logf("Expected: %.2f got: %.2f", expectedResult, actualResult)
	}
}

func TestMedian(t *testing.T) {
	expected1 := 5.0
	actual1 := mathstuff.Median([]float64{1.0, 10.1, 30.0, 2.0, 5.0})

	if fmt.Sprintf("%.2f", expected1) != fmt.Sprintf("%.2f", actual1) {
		t.Errorf("Expected: %.2f got: %.2f", expected1, actual1)
	} else {
		t.Logf("Expected: %.2f got: %.2f", expected1, actual1)
	}

	expected2 := 505.05
	actual2 := mathstuff.Median([]float64{1000.0, 10.1, 35550.0, 1111.0, 5.0, 1.0})

	if fmt.Sprintf("%.2f", expected2) != fmt.Sprintf("%.2f", actual2) {
		t.Errorf("Expected: %.2f got: %.2f", expected2, actual2)
	} else {
		t.Logf("Expected: %.2f got: %.2f", expected2, actual2)
	}
}

func TestMode(t *testing.T) {
	expected1 := []float64{1.0}
	actual1 := mathstuff.Mode([]float64{1.0, 10.1, 30.0, 2.0, 5.0, 1.0, 1.0, 1.0, 1.0})

	if !FloatArrayEquals(expected1, actual1) {
		t.Errorf("Expected: %.2f got: %.2f", expected1, actual1)
	} else {
		t.Logf("Expected: %.2f got: %.2f", expected1, actual1)
	}

	expected2 := []float64{1.0, 2.0}
	actual2 := mathstuff.Mode([]float64{1.0, 10.1, 30.0, 2.0, 2.0, 2.0, 2.0, 5.0, 1.0, 1.0, 1.0, 1.0, 2.0})

	if !FloatArrayEquals(expected2, actual2) {
		t.Errorf("Expected: %.2f got: %.2f", expected2, actual2)
	} else {
		t.Logf("Expected: %.2f got: %.2f", expected2, actual2)
	}
}

func TestRounding(t *testing.T) {
	expected1 := float64(1.33)
	actual1 := mathstuff.Average([]float64{1.334})
	if expected1 != actual1 {
		t.Errorf("Expected: %.2f got: %.2f", expected1, actual1)
	} else {
		t.Logf("Expected: %.2f got: %.2f", expected1, actual1)
	}

	expected2 := float64(1.34)
	actual2 := mathstuff.Average([]float64{1.335})
	if expected2 != actual2 {
		t.Errorf("Expected: %.2f got: %.2f", expected2, actual2)
	} else {
		t.Logf("Expected: %.2f got: %.2f", expected2, actual2)
	}
}

func FloatArrayEquals(a []float64, b []float64) bool {
	if len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if fmt.Sprintf("%.2f", v) != fmt.Sprintf("%.2f", b[i]) {
				return false
			}
		}
		return true
	}
}
