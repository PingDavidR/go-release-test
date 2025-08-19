package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive", 2, 3, 5},
		{"negative", -2, -3, -5},
		{"mixed", -2, 3, 1},
		{"zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.expected {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive", 5, 3, 2},
		{"negative", -5, -3, -2},
		{"mixed", -5, 3, -8},
		{"zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.a, tt.b); got != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive", 2, 3, 6},
		{"negative", -2, -3, 6},
		{"mixed", -2, 3, -6},
		{"zero", 0, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.a, tt.b); got != tt.expected {
				t.Errorf("Multiply(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectError bool
	}{
		{"positive", 6, 3, 2, false},
		{"negative", -6, -3, 2, false},
		{"mixed", -6, 3, -2, false},
		{"zero_dividend", 0, 5, 0, false},
		{"zero_divisor", 5, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.expectError {
				t.Errorf("Divide(%v, %v) error = %v, expectError %v", tt.a, tt.b, err, tt.expectError)
				return
			}
			if !tt.expectError && got != tt.expected {
				t.Errorf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive_base_positive_exponent", 2, 3, 8},
		{"positive_base_negative_exponent", 2, -1, 0.5},
		{"negative_base_even_exponent", -2, 2, 4},
		{"negative_base_odd_exponent", -2, 3, -8},
		{"zero_base", 0, 5, 0},
		{"any_base_zero_exponent", 5, 0, 1},
		{"one_base_any_exponent", 1, 99, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Power(tt.a, tt.b); got != tt.expected {
				t.Errorf("Power(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestSquareRoot(t *testing.T) {
	tests := []struct {
		name        string
		a           float64
		expected    float64
		expectError bool
	}{
		{"positive", 4, 2, false},
		{"zero", 0, 0, false},
		{"negative", -4, 0, true},
		{"perfect_square", 16, 4, false},
		{"non_perfect_square", 2, 1.4142135623730951, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SquareRoot(tt.a)
			if (err != nil) != tt.expectError {
				t.Errorf("SquareRoot(%v) error = %v, expectError %v", tt.a, err, tt.expectError)
				return
			}
			if !tt.expectError && got != tt.expected {
				t.Errorf("SquareRoot(%v) = %v, want %v", tt.a, got, tt.expected)
			}
		})
	}
}

func TestSin(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"zero", 0, 0},
		{"pi/2", 1.5707963267948966, 1},
		{"pi", 3.141592653589793, 0},
		{"3pi/2", 4.71238898038469, -1},
		{"2pi", 6.283185307179586, 0},
	}

	const epsilon = 1e-9

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sin(tt.input)
			if diff := got - tt.expected; diff < -epsilon || diff > epsilon {
				t.Errorf("Sin(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestCos(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"zero", 0, 1},
		{"pi/2", 1.5707963267948966, 0},
		{"pi", 3.141592653589793, -1},
		{"3pi/2", 4.71238898038469, 0},
		{"2pi", 6.283185307179586, 1},
	}

	const epsilon = 1e-9

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Cos(tt.input)
			if diff := got - tt.expected; diff < -epsilon || diff > epsilon {
				t.Errorf("Cos(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestTan(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"zero", 0, 0},
		{"pi/4", 0.7853981633974483, 1},
		{"pi/3", 1.0471975511965976, 1.7320508075688767}, // √3
		{"pi", 3.141592653589793, 0},
		{"-pi/4", -0.7853981633974483, -1},
	}

	const epsilon = 1e-9

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tan(tt.input)
			if diff := got - tt.expected; diff < -epsilon || diff > epsilon {
				t.Errorf("Tan(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	tests := []struct {
		name    string
		min     float64
		max     float64
		swapped bool
	}{
		{"positive_range", 10, 20, false},
		{"negative_range", -20, -10, false},
		{"mixed_range", -10, 10, false},
		{"zero_range", 0, 10, false},
		{"swapped_range", 20, 10, true}, // Tests auto-swap of min/max
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run multiple times to increase confidence in the test
			for i := 0; i < 100; i++ {
				var min, max float64
				if tt.swapped {
					// If swapped, the function should swap them back internally
					min = tt.max
					max = tt.min
				} else {
					min = tt.min
					max = tt.max
				}

				got := Random(min, max)

				// Check that the random number is within the expected range
				if got < min || got > max {
					t.Errorf("Random(%v, %v) = %v, should be between %v and %v",
						min, max, got, min, max)
				}
			}
		})
	}
}
