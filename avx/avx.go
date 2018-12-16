//+build amd64,cgo

package avx

/*
#cgo CFLAGS: -mavx -std=c99
#cgo LDFLAGS: -lm
#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include <immintrin.h> //AVX: -mavx

float avx_sum32(const size_t n, float *x) {
    static const size_t single_size = 8;
	const size_t end = n / single_size;
	__m256 vz={0};
    __m256 *vx = (__m256 *)x;
    for(size_t i=0; i<end; ++i) {
      vz = _mm256_add_ps(vx[i], vz);
	}
	float*z=(float*)&vz;
	for(size_t i=end*single_size;i<n;i++) {
		z[i&7]+=x[i];
	}
	return z[0]+z[1]+z[2]+z[3]+z[4]+z[5]+z[6]+z[7];
}

float avx_dot32(const size_t n, float *x, float *y)
{
    static const size_t single_size = 8;
    const size_t end = n / single_size;
    __m256 *vx = (__m256 *)x;
    __m256 *vy = (__m256 *)y;
    __m256 vsum = {0};
    for(size_t i=0; i<end; ++i) {
      vsum = _mm256_add_ps(vsum, _mm256_mul_ps(vx[i], vy[i]));
    }
    __attribute__((aligned(32))) float t[8] = {0};
    _mm256_store_ps(t, vsum);
    for(size_t i=end*single_size; i<n; ++i) {
      t[i&7] += x[i]*y[i];
    }
    return t[0] + t[1] + t[2] + t[3] + t[4] + t[5] + t[6] + t[7];
}

float avx_sum64(const size_t n, double *x) {
    static const size_t single_size = 4;
	const size_t end = n / single_size;
	__m256d vz={0};
    __m256d *vx = (__m256d *)x;
    for(size_t i=0; i<end; ++i) {
      vz = _mm256_add_pd(vx[i], vz);
	}
	double*z=(double*)&vz;
	for(size_t i=end*single_size;i<n;i++) {
		z[i&3]+=x[i];
	}
	return z[0]+z[1]+z[2]+z[3];
}

float avx_dot64(const size_t n, double *x, double *y)
{
    static const size_t single_size = 4;
    const size_t end = n / single_size;
    __m256d *vx = (__m256d *)x;
    __m256d *vy = (__m256d *)y;
    __m256d vsum = {0};
    for(size_t i=0; i<end; ++i) {
      vsum = _mm256_add_pd(vsum, _mm256_mul_pd(vx[i], vy[i]));
    }
    __attribute__((aligned(32))) double t[4] = {0};
    _mm256_store_pd(t, vsum);
    for(size_t i=end*single_size; i<n; ++i) {
      t[i&3] += x[i]*y[i];
    }
    return t[0] + t[1] + t[2] + t[3];
}

*/
import (
	"C"
)

func Sum32(a []float32) float32 {
	return float32(C.avx_sum32(C.ulong(len(a)), (*C.float)(&a[0])))
}

func Dot32(size int, x, y []float32) float32 {
	dot := C.avx_dot32((C.size_t)(size), (*C.float)(&x[0]), (*C.float)(&y[0]))
	return float32(dot)
}

func Sum64(a []float64) float64 {
	return float64(C.avx_sum64(C.ulong(len(a)), (*C.double)(&a[0])))
}

func Dot64(size int, x, y []float64) float64 {
	dot := C.avx_dot64((C.size_t)(size), (*C.double)(&x[0]), (*C.double)(&y[0]))
	return float64(dot)
}
