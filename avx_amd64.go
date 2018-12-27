//+build !noasm
//+build !appengine

package float

import (
	"unsafe"
)

//go:noescape
func _avxFloat32Dot(n int, x, y, res unsafe.Pointer)

func avxFloat32Dot(x, y []float32) (res float32) {
	n := len(x)
	y = y[:n]
	_avxFloat32Dot(n, unsafe.Pointer(&x[0]), unsafe.Pointer(&y[0]), unsafe.Pointer(&res))
	return
}

//go:noescape
func _avxFloat64Dot(n int, x, y, res unsafe.Pointer)

func avxFloat64Dot(x, y []float64) (res float64) {
	n := len(x)
	y = y[:n]
	_avxFloat64Dot(n, unsafe.Pointer(&x[0]), unsafe.Pointer(&y[0]), unsafe.Pointer(&res))
	return
}

/*
//go:noescape
func _avx_float32_addto(n int, x, y, res unsafe.Pointer)

func avx_float32_addto(x, y, res []float32) {
	n := len(x)
	y = y[:n]
	res = res[:n]
	_avx_float32_addto(n, unsafe.Pointer(&x[0]), unsafe.Pointer(&y[0]), unsafe.Pointer(&res[0]))
}

//go:noescape
func _avx_float64_addto(n int, x, y, res unsafe.Pointer)

func avx_float64_addto(x, y, res []float64) {
	n := len(x)
	y = y[:n]
	res = res[:n]
	_avx_float64_addto(n, unsafe.Pointer(&x[0]), unsafe.Pointer(&y[0]), unsafe.Pointer(&res[0]))
}
*/
