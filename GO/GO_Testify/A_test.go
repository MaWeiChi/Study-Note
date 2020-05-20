package a

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

// 計算並返回 x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}
