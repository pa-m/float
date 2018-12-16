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

func TestAddTo32(t *testing.T) {
	var avx32 F32
	expected, actual, x, y := make([]float32, 8759), make([]float32, 8759), make([]float32, 8759), make([]float32, 8759)
	for i := range x {
		x[i], y[i] = rand.Float32(), rand.Float32()
		expected[i] = x[i] + y[i]
	}
	avx32.AddTo(actual, x, y)
	for i := range x {
		if abs32(expected[i]-actual[i]) > .01 {
			t.Errorf("expected[%d] %g actual %g", i, expected, actual)
		}
	}
}

func TestSubTo32(t *testing.T) {
	var avx32 F32
	expected, actual, x, y := make([]float32, 8759), make([]float32, 8759), make([]float32, 8759), make([]float32, 8759)
	for i := range x {
		x[i], y[i] = rand.Float32(), rand.Float32()
		expected[i] = x[i] - y[i]
	}
	avx32.SubTo(actual, x, y)
	for i := range x {
		if abs32(expected[i]-actual[i]) > .01 {
			t.Errorf("expected[%d] %g actual %g", i, expected, actual)
		}
	}
}

func TestMulTo32(t *testing.T) {
	var avx32 F32
	expected, actual, x, y := make([]float32, 8759), make([]float32, 8759), make([]float32, 8759), make([]float32, 8759)
	for i := range x {
		x[i], y[i] = rand.Float32(), rand.Float32()
		expected[i] = x[i] * y[i]
	}
	avx32.MulTo(actual, x, y)
	for i := range x {
		if abs32(expected[i]-actual[i]) > .01 {
			t.Errorf("expected[%d] %g actual %g", i, expected, actual)
		}
	}
}

func TestScaleTo32(t *testing.T) {
	var avx32 F32
	expected, actual, x, scale := make([]float32, 8759), make([]float32, 8759), make([]float32, 8759), rand.Float32()
	for i := range x {
		x[i] = rand.Float32()
		expected[i] = x[i] * scale
	}
	avx32.ScaleTo(actual, x, scale)
	for i := range x {
		if abs32(expected[i]-actual[i]) > .01 {
			t.Errorf("expected[%d] %g actual %g", i, expected, actual)
		}
	}
}

func TestAddTo64(t *testing.T) {
	var avx64 F64
	expected, actual, x, y := make([]float64, 8759), make([]float64, 8759), make([]float64, 8759), make([]float64, 8759)
	for i := range x {
		x[i], y[i] = rand.Float64(), rand.Float64()
		expected[i] = x[i] + y[i]
	}
	avx64.AddTo(actual, x, y)
	for i := range x {
		if abs64(expected[i]-actual[i]) > .01 {
			t.Errorf("expected[%d] %g actual %g", i, expected, actual)
		}
	}
}

func TestSubTo64(t *testing.T) {
	var avx64 F64
	expected, actual, x, y := make([]float64, 8759), make([]float64, 8759), make([]float64, 8759), make([]float64, 8759)
	for i := range x {
		x[i], y[i] = rand.Float64(), rand.Float64()
		expected[i] = x[i] - y[i]
	}
	avx64.SubTo(actual, x, y)
	for i := range x {
		if abs64(expected[i]-actual[i]) > .01 {
			t.Errorf("expected[%d] %g actual %g", i, expected, actual)
		}
	}
}

func TestMulTo64(t *testing.T) {
	var avx64 F64
	expected, actual, x, y := make([]float64, 8759), make([]float64, 8759), make([]float64, 8759), make([]float64, 8759)
	for i := range x {
		x[i], y[i] = rand.Float64(), rand.Float64()
		expected[i] = x[i] * y[i]
	}
	avx64.MulTo(actual, x, y)
	for i := range x {
		if abs64(expected[i]-actual[i]) > .01 {
			t.Errorf("expected[%d] %g actual %g", i, expected, actual)
		}
	}
}

func TestScaleTo64(t *testing.T) {
	var avx64 F64
	expected, actual, x, scale := make([]float64, 8759), make([]float64, 8759), make([]float64, 8759), rand.Float64()
	for i := range x {
		x[i] = rand.Float64()
		expected[i] = x[i] * scale
	}
	avx64.ScaleTo(actual, x, scale)
	for i := range x {
		if abs64(expected[i]-actual[i]) > .01 {
			t.Errorf("expected[%d] %g actual %g", i, expected, actual)
		}
	}
}

func TestSum32(t *testing.T) {
	var avx32 F32
	var expected, actual float32
	x, y := make([]float32, 8759), make([]float32, 8759)
	for i := range x {
		x[i], y[i] = rand.Float32(), rand.Float32()
		expected += x[i]
	}
	actual = avx32.Sum(x)
	if abs32(expected-actual) > .01 {
		t.Errorf("expected %g actual %g", expected, actual)
	}
}
func TestSum64(t *testing.T) {
	var avx64 F64
	var expected, actual float64
	x, y := make([]float64, 8759), make([]float64, 8759)
	for i := range x {
		x[i], y[i] = rand.Float64(), rand.Float64()
		expected += x[i]
	}
	actual = avx64.Sum(x)
	if abs64(expected-actual) > .01 {
		t.Errorf("expected %g actual %g", expected, actual)
	}
}
func BenchmarkSum32(b *testing.B) {
	var avx32 F32
	a := make([]float32, 8760)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avx32.Sum(a)
	}
}
func BenchmarkSum64(b *testing.B) {
	var avx64 F64
	a := make([]float64, 8760)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avx64.Sum(a)
	}
}
func TestDot32(t *testing.T) {
	var avx32 F32
	var expected, actual float32
	x, y := make([]float32, 8759), make([]float32, 8759)
	for i := range x {
		x[i], y[i] = rand.Float32(), rand.Float32()
		expected += x[i] * y[i]
	}
	actual = avx32.Dot(x, y)
	if abs32(expected-actual) > .01 {
		t.Errorf("expected %g actual %g", expected, actual)
	}
}
func TestDot64(t *testing.T) {
	var avx64 F64
	var expected, actual float64
	x, y := make([]float64, 8759), make([]float64, 8759)
	for i := range x {
		x[i], y[i] = rand.Float64(), rand.Float64()
		expected += x[i] * y[i]
	}
	actual = avx64.Dot(x, y)
	if abs64(expected-actual) > .01 {
		t.Errorf("expected %g actual %g", expected, actual)
	}
}

func BenchmarkDot32(b *testing.B) {
	var avx32 F32
	a := make([]float32, 8760)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avx32.Dot(a, a)
	}
}
func BenchmarkDot64(b *testing.B) {
	var avx64 F64
	a := make([]float64, 8760)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avx64.Dot(a, a)
	}
}
