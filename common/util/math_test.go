package util

import (
	"fmt"
	"testing"
)

func TestCalculatePercentageChange(t *testing.T) {
	c1, _ := CalculatePercentageChange("0.01", "0.02")
	t.Log(fmt.Sprintf("%.2f%%", c1))

	c2, _ := CalculatePercentageChange("0.02", "0.01")
	t.Log(fmt.Sprintf("%.2f%%", c2))

}
