package calc

import "testing"

func TestCalculateSum(t *testing.T) {
	result := CalculateSum(5)

	if result != 15 {
		t.Errorf("CalculateSum(5) = %v, Expected 15", result)
	}
}

// BenchmarkCalculateSum benchmarks the CalculateSum function
func BenchmarkCalculateSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSum(100)
	}
}
