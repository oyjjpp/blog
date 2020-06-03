package decimal

import "testing"

func TestDiv(t *testing.T) {
	num1 := 0.3
	num2 := 0.05
	number1 := NewFromFloat(num1)
	number2 := NewFromFloat(num2)
	rs, bo := number1.Div(number2).Float64()
	t.Log(rs, bo)
}
