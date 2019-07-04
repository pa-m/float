// Code generated by float.gen.go.tmpl. DO NOT EDIT.

package float

import (
	"fmt"
	"sort"

	math32 "github.com/chewxy/math32"

	math64 "math"
)

type F32 float32
type F32s []float32

var f32 F32

func (F32) NaN() float32 { return math32.NaN() }

func (F32) IsNaN(x float32) bool { return math32.IsNaN(x) }

func (F32) Inf(sgn int) float32 { return math32.Inf(sgn) }

func (F32) IsInf(x float32, sgn int) bool { return math32.IsInf(x, sgn) }

func (F32) Iif(cond bool, a, b float32) float32 {
	if cond {
		return a
	}
	return b
}

// Len (sort.Interface)
func (a F32s) Len() int { return len(a) }

// Less  (sort.Interface)
func (a F32s) Less(i, j int) bool { return a[i] < a[j] }

// Swap  (sort.Interface)
func (a F32s) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Sort  (sort.Interface)
func (a F32s) Sort() {
	sort.Sort(F32s(a))
}

// Search  (sort.Interface)
func (a F32s) Search(x float32) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

// ArgSort ...
func (F32) ArgSort(a []float32) []int {
	ind := make([]int, len(a))
	for i := range a {
		ind[i] = i
	}
	sort.Slice(ind, func(i, j int) bool { return a[ind[i]] < a[ind[j]] })
	return ind
}

// Median ...
func (F32) Median(a []float32) float32 {
	var sorted []float32
	if sort.IsSorted(F32s(a)) {
		sorted = a
	} else {
		sorted = make([]float32, len(a))
		copy(sorted, a)
		sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	}
	return sorted[int(len(sorted)/2)]
}

// Mean for []float32
func (F32) Mean(a []float32) float32 {
	return f32.Sum(a) / float32(len(a))
}

// Abs for []float32
func (F32) Abs(a float32) float32 {
	if a < float32(0) {
		return -a
	}
	return a
}

// MaxIdx ...
func (F32) MaxIdx(a []float32) int {
	mi := -1
	for i, v := range a {
		if mi < 0 || v > a[mi] {
			mi = i
		}
	}
	return mi
}

// MinIdx ...
func (F32) MinIdx(a []float32) int {
	mi := -1
	for i, v := range a {
		if mi < 0 || v < a[mi] {
			mi = i
		}
	}
	return mi
}

// Max ...
func (F32) Max(a []float32) float32 {
	mi := f32.MaxIdx(a)
	if mi < 0 {
		return math32.NaN()
	}
	return a[mi]
}

// Min ...
func (F32) Min(a []float32) float32 {
	mi := f32.MinIdx(a)
	if mi < 0 {
		return math32.NaN()
	}
	return a[mi]
}

// Alloc []float32
func (F32) Alloc(n int, ptrs ...*F32s) {
	c := n * len(ptrs)
	a := make(F32s, c)
	for i, ptr := range ptrs {
		*ptr = a[n*i : n*i+n]
	}
}

// EqualWithinAbs returns true if a and b have an absolute
// difference of less than tol.
func (F32) EqualWithinAbs(a, b, tol float32) bool {
	return a == b || math32.Abs(a-b) <= tol
}

// CheckFloat ...
func (F32) CheckFloat(msg string, ai float32) {
	if math32.IsInf(ai, -1) {
		panic(fmt.Errorf("-inf at %s", msg))

	}
	if math32.IsInf(ai, 1) {
		panic(fmt.Errorf("+inf at %s", msg))

	}
	if math32.IsNaN(ai) {
		panic(fmt.Errorf("nan at %s", msg))
	}

}

// Scale ...
func (f32 F32) Scale(scale float32, a []float32) {
	f32.ScaleTo(a, a, scale)
}

// Add ...
func (f32 F32) Add(dst, a []float32) {
	f32.AddTo(dst, dst, a)
}

// Mul ...
func (f32 F32) Mul(dst, a []float32) {
	f32.MulTo(dst, dst, a)
}

// Sign ...
func (F32) Sign(x float32) float32 {
	if math32.Signbit(x) {
		return -1
	}
	return 1
}

// Reduce ...
func (F32) Reduce(a []float32, f func(carry, item float32) float32, init float32) float32 {
	r := init
	for _, v := range a {
		r = f(r, v)
	}
	return r
}

// Square ...
func (F32) Square(x float32) float32 { return x * x }

// vanillaAddTo perform naive dst = a+b for []float32
func (F32) vanillaAddTo(dst, a, b []float32) {
	l := len(a)
	b = b[:l]
	for i := 0; i < l; i++ {
		dst[i] = a[i] + b[i]
	}
}

// unrolledAddTo perform dst = a+b with loop unrolling for []float32
func (F32) unrolledAddTo(dst, a, b []float32) {
	const groupsize = 8

	l := len(a)
	b = b[:l]
	end := (l / groupsize) * groupsize
	var i int
	for i = 0; i < end; i += groupsize {

		dst[i+0] = a[i+0] + b[i+0]

		dst[i+1] = a[i+1] + b[i+1]

		dst[i+2] = a[i+2] + b[i+2]

		dst[i+3] = a[i+3] + b[i+3]

		dst[i+4] = a[i+4] + b[i+4]

		dst[i+5] = a[i+5] + b[i+5]

		dst[i+6] = a[i+6] + b[i+6]

		dst[i+7] = a[i+7] + b[i+7]

	}
	for i = end; i < l; i++ {
		dst[i] = a[i] + b[i]
	}

}

// vanillaSubTo perform naive dst = a - b for []float32
func (F32) vanillaSubTo(dst, a, b []float32) {
	l := len(a)
	b = b[:l]
	for i := 0; i < l; i++ {
		dst[i] = a[i] - b[i]
	}
}

// unrolledSubTo perform dst = a - b with loop unrolling for []float32
func (F32) unrolledSubTo(dst, a, b []float32) {
	const groupsize = 8
	l := len(a)
	b = b[:l]
	end := (l / groupsize) * groupsize
	var i int
	for i = 0; i < end; i += groupsize {

		dst[i+0] = a[i+0] - b[i+0]

		dst[i+1] = a[i+1] - b[i+1]

		dst[i+2] = a[i+2] - b[i+2]

		dst[i+3] = a[i+3] - b[i+3]

		dst[i+4] = a[i+4] - b[i+4]

		dst[i+5] = a[i+5] - b[i+5]

		dst[i+6] = a[i+6] - b[i+6]

		dst[i+7] = a[i+7] - b[i+7]

	}
	for i = end; i < l; i++ {
		dst[i] = a[i] - b[i]
	}
}

// vanillaMulTo perform naive dst = a * b for []float32
func (F32) vanillaMulTo(dst, a, b []float32) {
	l := len(a)
	b = b[:l]
	for i := 0; i < l; i++ {
		dst[i] = a[i] * b[i]
	}
}

// unrolledMulTo perform dst = a * b with loop unrolling for []float32
func (F32) unrolledMulTo(dst, a, b []float32) {
	const groupsize = 8
	l := len(a)
	b = b[:l]
	end := (l / groupsize) * groupsize
	var i int
	for i = 0; i < end; i += groupsize {

		dst[i+0] = a[i+0] * b[i+0]

		dst[i+1] = a[i+1] * b[i+1]

		dst[i+2] = a[i+2] * b[i+2]

		dst[i+3] = a[i+3] * b[i+3]

		dst[i+4] = a[i+4] * b[i+4]

		dst[i+5] = a[i+5] * b[i+5]

		dst[i+6] = a[i+6] * b[i+6]

		dst[i+7] = a[i+7] * b[i+7]

	}
	for i = end; i < l; i++ {
		dst[i] = a[i] * b[i]
	}
}

// vanillaScaleTo perform naive dst = a * scale for []float32
func (F32) vanillaScaleTo(dst, a []float32, scale float32) {
	l := len(a)
	for i := 0; i < l; i++ {
		dst[i] = a[i] * scale
	}
}

// unrolledScaleTo perform dst = a * scale with loop unrolling for []float32
func (F32) unrolledScaleTo(dst, a []float32, scale float32) {
	const groupsize = 8
	l := len(a)
	end := (l / groupsize) * groupsize
	var i int
	for i = 0; i < end; i += groupsize {

		dst[i+0] = a[i+0] * scale

		dst[i+1] = a[i+1] * scale

		dst[i+2] = a[i+2] * scale

		dst[i+3] = a[i+3] * scale

		dst[i+4] = a[i+4] * scale

		dst[i+5] = a[i+5] * scale

		dst[i+6] = a[i+6] * scale

		dst[i+7] = a[i+7] * scale

	}
	for i = end; i < l; i++ {
		dst[i] = a[i] * scale
	}
}

// vanillaSum for []float32
func (F32) vanillaSum(a []float32) float32 {
	var acc float32
	l := len(a)
	for i := 0; i < l; i++ {
		acc += a[i]
	}
	return acc
}

// unrolledSum for []float32
func (F32) unrolledSum(a []float32) float32 {
	var acc float32
	const groupsize = 8

	l := len(a)
	end := (l / groupsize) * groupsize
	var i int
	for i = 0; i < end; i += groupsize {
		acc += +a[i+0] + a[i+1] + a[i+2] + a[i+3] + a[i+4] + a[i+5] + a[i+6] + a[i+7]
	}
	for i = end; i < l; i++ {
		acc += a[i]
	}
	return acc
}

// vanillaDot for []float32
func (F32) vanillaDot(a, b []float32) float32 {
	var acc float32
	l := len(a)
	b = b[:l]
	for i := 0; i < l; i++ {
		acc += a[i] * b[i]
	}
	return acc
}

// unrolledDot for []float32
func (F32) unrolledDot(a, b []float32) float32 {
	var acc float32
	const groupsize = 8
	l := len(a)
	b = b[:l]
	end := (l / groupsize) * groupsize
	var i int
	for i = 0; i < end; i += groupsize {
		acc += +a[i+0]*b[i+0] + a[i+1]*b[i+1] + a[i+2]*b[i+2] + a[i+3]*b[i+3] + a[i+4]*b[i+4] + a[i+5]*b[i+5] + a[i+6]*b[i+6] + a[i+7]*b[i+7]
	}
	for i = end; i < l; i++ {
		acc += a[i] * b[i]
	}
	return acc
}

type F64 float64
type F64s []float64

var f64 F64

func (F64) NaN() float64 { return math64.NaN() }

func (F64) IsNaN(x float64) bool { return math64.IsNaN(x) }

func (F64) Inf(sgn int) float64 { return math64.Inf(sgn) }

func (F64) IsInf(x float64, sgn int) bool { return math64.IsInf(x, sgn) }

func (F64) Iif(cond bool, a, b float64) float64 {
	if cond {
		return a
	}
	return b
}

// Len (sort.Interface)
func (a F64s) Len() int { return len(a) }

// Less  (sort.Interface)
func (a F64s) Less(i, j int) bool { return a[i] < a[j] }

// Swap  (sort.Interface)
func (a F64s) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Sort  (sort.Interface)
func (a F64s) Sort() {
	sort.Sort(F64s(a))
}

// Search  (sort.Interface)
func (a F64s) Search(x float64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

// ArgSort ...
func (F64) ArgSort(a []float64) []int {
	ind := make([]int, len(a))
	for i := range a {
		ind[i] = i
	}
	sort.Slice(ind, func(i, j int) bool { return a[ind[i]] < a[ind[j]] })
	return ind
}

// Median ...
func (F64) Median(a []float64) float64 {
	var sorted []float64
	if sort.IsSorted(F64s(a)) {
		sorted = a
	} else {
		sorted = make([]float64, len(a))
		copy(sorted, a)
		sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	}
	return sorted[int(len(sorted)/2)]
}

// Mean for []float64
func (F64) Mean(a []float64) float64 {
	return f64.Sum(a) / float64(len(a))
}

// Abs for []float64
func (F64) Abs(a float64) float64 {
	if a < float64(0) {
		return -a
	}
	return a
}

// MaxIdx ...
func (F64) MaxIdx(a []float64) int {
	mi := -1
	for i, v := range a {
		if mi < 0 || v > a[mi] {
			mi = i
		}
	}
	return mi
}

// MinIdx ...
func (F64) MinIdx(a []float64) int {
	mi := -1
	for i, v := range a {
		if mi < 0 || v < a[mi] {
			mi = i
		}
	}
	return mi
}

// Max ...
func (F64) Max(a []float64) float64 {
	mi := f64.MaxIdx(a)
	if mi < 0 {
		return math64.NaN()
	}
	return a[mi]
}

// Min ...
func (F64) Min(a []float64) float64 {
	mi := f64.MinIdx(a)
	if mi < 0 {
		return math64.NaN()
	}
	return a[mi]
}

// Alloc []float64
func (F64) Alloc(n int, ptrs ...*F64s) {
	c := n * len(ptrs)
	a := make(F64s, c)
	for i, ptr := range ptrs {
		*ptr = a[n*i : n*i+n]
	}
}

// EqualWithinAbs returns true if a and b have an absolute
// difference of less than tol.
func (F64) EqualWithinAbs(a, b, tol float64) bool {
	return a == b || math64.Abs(a-b) <= tol
}

// CheckFloat ...
func (F64) CheckFloat(msg string, ai float64) {
	if math64.IsInf(ai, -1) {
		panic(fmt.Errorf("-inf at %s", msg))

	}
	if math64.IsInf(ai, 1) {
		panic(fmt.Errorf("+inf at %s", msg))

	}
	if math64.IsNaN(ai) {
		panic(fmt.Errorf("nan at %s", msg))
	}

}

// Scale ...
func (f64 F64) Scale(scale float64, a []float64) {
	f64.ScaleTo(a, a, scale)
}

// Add ...
func (f64 F64) Add(dst, a []float64) {
	f64.AddTo(dst, dst, a)
}

// Mul ...
func (f64 F64) Mul(dst, a []float64) {
	f64.MulTo(dst, dst, a)
}

// Sign ...
func (F64) Sign(x float64) float64 {
	if math64.Signbit(x) {
		return -1
	}
	return 1
}

// Reduce ...
func (F64) Reduce(a []float64, f func(carry, item float64) float64, init float64) float64 {
	r := init
	for _, v := range a {
		r = f(r, v)
	}
	return r
}

// Square ...
func (F64) Square(x float64) float64 { return x * x }

// vanillaAddTo perform naive dst = a+b for []float64
func (F64) vanillaAddTo(dst, a, b []float64) {
	l := len(a)
	b = b[:l]
	for i := 0; i < l; i++ {
		dst[i] = a[i] + b[i]
	}
}

// unrolledAddTo perform dst = a+b with loop unrolling for []float64
func (F64) unrolledAddTo(dst, a, b []float64) {
	const groupsize = 4

	l := len(a)
	b = b[:l]
	end := (l / groupsize) * groupsize
	var i int
	for i = 0; i < end; i += groupsize {

		dst[i+0] = a[i+0] + b[i+0]

		dst[i+1] = a[i+1] + b[i+1]

		dst[i+2] = a[i+2] + b[i+2]

		dst[i+3] = a[i+3] + b[i+3]

	}
	for i = end; i < l; i++ {
		dst[i] = a[i] + b[i]
	}

}

// vanillaSubTo perform naive dst = a - b for []float64
func (F64) vanillaSubTo(dst, a, b []float64) {
	l := len(a)
	b = b[:l]
	for i := 0; i < l; i++ {
		dst[i] = a[i] - b[i]
	}
}

// unrolledSubTo perform dst = a - b with loop unrolling for []float64
func (F64) unrolledSubTo(dst, a, b []float64) {
	const groupsize = 4
	l := len(a)
	b = b[:l]
	end := (l / groupsize) * groupsize
	var i int
	for i = 0; i < end; i += groupsize {

		dst[i+0] = a[i+0] - b[i+0]

		dst[i+1] = a[i+1] - b[i+1]

		dst[i+2] = a[i+2] - b[i+2]

		dst[i+3] = a[i+3] - b[i+3]

	}
	for i = end; i < l; i++ {
		dst[i] = a[i] - b[i]
	}
}

// vanillaMulTo perform naive dst = a * b for []float64
func (F64) vanillaMulTo(dst, a, b []float64) {
	l := len(a)
	b = b[:l]
	for i := 0; i < l; i++ {
		dst[i] = a[i] * b[i]
	}
}

// unrolledMulTo perform dst = a * b with loop unrolling for []float64
func (F64) unrolledMulTo(dst, a, b []float64) {
	const groupsize = 4
	l := len(a)
	b = b[:l]
	end := (l / groupsize) * groupsize
	var i int
	for i = 0; i < end; i += groupsize {

		dst[i+0] = a[i+0] * b[i+0]

		dst[i+1] = a[i+1] * b[i+1]

		dst[i+2] = a[i+2] * b[i+2]

		dst[i+3] = a[i+3] * b[i+3]

	}
	for i = end; i < l; i++ {
		dst[i] = a[i] * b[i]
	}
}

// vanillaScaleTo perform naive dst = a * scale for []float64
func (F64) vanillaScaleTo(dst, a []float64, scale float64) {
	l := len(a)
	for i := 0; i < l; i++ {
		dst[i] = a[i] * scale
	}
}

// unrolledScaleTo perform dst = a * scale with loop unrolling for []float64
func (F64) unrolledScaleTo(dst, a []float64, scale float64) {
	const groupsize = 4
	l := len(a)
	end := (l / groupsize) * groupsize
	var i int
	for i = 0; i < end; i += groupsize {

		dst[i+0] = a[i+0] * scale

		dst[i+1] = a[i+1] * scale

		dst[i+2] = a[i+2] * scale

		dst[i+3] = a[i+3] * scale

	}
	for i = end; i < l; i++ {
		dst[i] = a[i] * scale
	}
}

// vanillaSum for []float64
func (F64) vanillaSum(a []float64) float64 {
	var acc float64
	l := len(a)
	for i := 0; i < l; i++ {
		acc += a[i]
	}
	return acc
}

// unrolledSum for []float64
func (F64) unrolledSum(a []float64) float64 {
	var acc float64
	const groupsize = 4

	l := len(a)
	end := (l / groupsize) * groupsize
	var i int
	for i = 0; i < end; i += groupsize {
		acc += +a[i+0] + a[i+1] + a[i+2] + a[i+3]
	}
	for i = end; i < l; i++ {
		acc += a[i]
	}
	return acc
}

// vanillaDot for []float64
func (F64) vanillaDot(a, b []float64) float64 {
	var acc float64
	l := len(a)
	b = b[:l]
	for i := 0; i < l; i++ {
		acc += a[i] * b[i]
	}
	return acc
}

// unrolledDot for []float64
func (F64) unrolledDot(a, b []float64) float64 {
	var acc float64
	const groupsize = 4
	l := len(a)
	b = b[:l]
	end := (l / groupsize) * groupsize
	var i int
	for i = 0; i < end; i += groupsize {
		acc += +a[i+0]*b[i+0] + a[i+1]*b[i+1] + a[i+2]*b[i+2] + a[i+3]*b[i+3]
	}
	for i = end; i < l; i++ {
		acc += a[i] * b[i]
	}
	return acc
}
