//+build amd64,cgo

package avx

import (
	"math/rand"
	"testing"
)

func abs32(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}
func abs64(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func TestSum32(t *testing.T) {
	var expected, actual float32
	x, y := make([]float32, 8759), make([]float32, 8759)
	for i := range x {
		x[i], y[i] = rand.Float32(), rand.Float32()
		expected += x[i]
	}
	actual = Sum32(x)
	if abs32(expected-actual) > .01 {
		t.Errorf("expected %g actual %g", expected, actual)
	}
}
func TestSum64(t *testing.T) {
	var expected, actual float64
	x, y := make([]float64, 8759), make([]float64, 8759)
	for i := range x {
		x[i], y[i] = rand.Float64(), rand.Float64()
		expected += x[i]
	}
	actual = Sum64(x)
	if abs64(expected-actual) > .01 {
		t.Errorf("expected %g actual %g", expected, actual)
	}
}
func BenchmarkSum32(b *testing.B) {
	a := make([]float32, 8760)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum32(a)
	}
}
func BenchmarkSum64(b *testing.B) {
	a := make([]float64, 8760)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum64(a)
	}
}
func TestDot32(t *testing.T) {
	var expected, actual float32
	x, y := make([]float32, 8759), make([]float32, 8759)
	for i := range x {
		x[i], y[i] = rand.Float32(), rand.Float32()
		expected += x[i] * y[i]
	}
	actual = Dot32(len(x), x, y)
	if abs32(expected-actual) > .01 {
		t.Errorf("expected %g actual %g", expected, actual)
	}
}
func TestDot64(t *testing.T) {
	var expected, actual float64
	x, y := make([]float64, 8759), make([]float64, 8759)
	for i := range x {
		x[i], y[i] = rand.Float64(), rand.Float64()
		expected += x[i] * y[i]
	}
	actual = Dot64(len(x), x, y)
	if abs64(expected-actual) > .01 {
		t.Errorf("expected %g actual %g", expected, actual)
	}
}
func BenchmarkDot32(b *testing.B) {
	a := make([]float32, 8760)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dot32(len(a), a, a)
	}
}
func BenchmarkDot64(b *testing.B) {
	a := make([]float64, 8760)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dot64(len(a), a, a)
	}
}
