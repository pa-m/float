package float

import (
	"fmt"
	"testing"

	avx "./avx"
	math32 "github.com/chewxy/math32"
	"gonum.org/v1/gonum/blas/blas32"
	"gorgonia.org/vecf32"
)

func dot_go(a, b []float32) (s float32) {
	for i, ai := range a {
		s += ai * b[i]
	}
	return
}
func dot_blas(a, b []float32) (s float32) {
	s = blas32.Dot(len(a), blas32.Vector{1, a}, blas32.Vector{1, b})
	return
}

type dotvecf32 []float32

func (tmp *dotvecf32) dot(a, b []float32) (s float32) {
	copy(*tmp, a)
	vecf32.Mul(*tmp, b)
	return vecf32.Sum(*tmp)
}

var testAB struct{ a, b []float32 }

func init() {
	testAB.a = make([]float32, 100000)
	testAB.b = make([]float32, 100000)
	for i := range testAB.a {
		testAB.a[i] = float32(i)
		testAB.b[i] = float32(100000 - i)
	}
	return
}

func ExampleDotGo() {
	a, b := testAB.a, testAB.b
	s := dot_go(a, b)
	fmt.Println(s)
	// Output:
	// 1.6666857e+14
}
func BenchmarkDotGo(b *testing.B) {
	x, y := testAB.a, testAB.b
	var actual float32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual = dot_go(x, y)
	}
	expected := float32(1.6666857e+14)
	if math32.Abs(expected-actual) > expected/1e6 {
		b.Errorf("expected:%g actual:%g diff:%g", expected, actual, expected-actual)
	}
}
func BenchmarkDotBlas(b *testing.B) {
	x, y := testAB.a, testAB.b
	var actual float32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual = dot_blas(x, y)
	}
	expected := float32(1.6666857e+14)
	if math32.Abs(expected-actual) > expected/1e4 {
		b.Errorf("expected:%g actual:%g diff:%g", expected, actual, expected-actual)
	}

}
func BenchmarkDotvecf32(b *testing.B) {
	x, y := testAB.a, testAB.b
	var actual float32
	tmp := make(dotvecf32, len(x))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual = tmp.dot(x, y)
	}
	expected := float32(1.6666857e+14)
	if math32.Abs(expected-actual) > expected/1e4 {
		b.Errorf("expected:%g actual:%g diff:%g", expected, actual, expected-actual)
	}

}

func BenchmarkDot_avx(b *testing.B) {
	x, y := testAB.a, testAB.b
	var actual float32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual = avx.Dot32(len(x), x, y)
	}
	expected := float32(1.6666857e+14)
	if math32.Abs(expected-actual) > expected/1e4 {
		b.Errorf("expected:%g actual:%g diff:%g", expected, actual, expected-actual)
	}

}
