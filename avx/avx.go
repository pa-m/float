//+build amd64,cgo

package avx

/*
#cgo CFLAGS: -mavx -std=c99
#cgo LDFLAGS: -lm
#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include <immintrin.h> //AVX: -mavx

float avx_float32_sum(const size_t n, float *x) {
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

float avx_float32_dot(const size_t n, float *x, float *y)
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

void avx_float32_addto(const size_t n, float*dst,float *x, float *y){
    static const size_t single_size = 8;
    const size_t end = n / single_size;
    __m256 *vdst = (__m256 *)dst;
    __m256 *vx = (__m256 *)x;
    __m256 *vy = (__m256 *)y;
    for(size_t i=0; i<end; ++i) {
      vdst[i] = _mm256_add_ps(vx[i], vy[i]);
    }
    for(size_t i=end*single_size; i<n; ++i) {
      dst[i] = x[i]+y[i];
    }
}

void avx_float32_subto(const size_t n, float*dst,float *x, float *y){
    static const size_t single_size = 8;
    const size_t end = n / single_size;
    __m256 *vdst = (__m256 *)dst;
    __m256 *vx = (__m256 *)x;
    __m256 *vy = (__m256 *)y;
    for(size_t i=0; i<end; ++i) {
      vdst[i] = _mm256_sub_ps(vx[i], vy[i]);
    }
    for(size_t i=end*single_size; i<n; ++i) {
      dst[i] = x[i]-y[i];
    }
}

void avx_float32_multo(const size_t n, float*dst,float *x, float *y){
    static const size_t single_size = 8;
    const size_t end = n / single_size;
    __m256 *vdst = (__m256 *)dst;
    __m256 *vx = (__m256 *)x;
    __m256 *vy = (__m256 *)y;
    for(size_t i=0; i<end; ++i) {
      vdst[i] = _mm256_mul_ps(vx[i], vy[i]);
    }
    for(size_t i=end*single_size; i<n; ++i) {
      dst[i] = x[i]*y[i];
    }
}

void avx_float32_scaleto(const size_t n, float*dst,float *x, float scale){
    static const size_t single_size = 8;
    const size_t end = n / single_size;
    __m256 *vdst = (__m256 *)dst;
    __m256 *vx = (__m256 *)x;
    __m256 vy = _mm256_broadcast_ss(&scale);
    for(size_t i=0; i<end; ++i) {
      vdst[i] = _mm256_mul_ps(vx[i], vy);
    }
    for(size_t i=end*single_size; i<n; ++i) {
      dst[i] = x[i]*scale;
    }
}



float avx_float64_sum(const size_t n, double *x) {
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

float avx_float64_dot(const size_t n, double *x, double *y)
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

void avx_float64_addto(const size_t n, double *dst,double *x, double *y)
{
    static const size_t single_size = 4;
    const size_t end = n / single_size;
    __m256d *vdst =(__m256d *)dst;
    __m256d *vx = (__m256d *)x;
    __m256d *vy = (__m256d *)y;
    for(size_t i=0; i<end; ++i) {
      vdst[i] = _mm256_add_pd(vx[i], vy[i]);
    }
    for(size_t i=end*single_size; i<n; ++i) {
      dst[i] = x[i]+y[i];
    }
}
void avx_float64_subto(const size_t n, double *dst,double *x, double *y)
{
    static const size_t single_size = 4;
    const size_t end = n / single_size;
    __m256d *vdst =(__m256d *)dst;
    __m256d *vx = (__m256d *)x;
    __m256d *vy = (__m256d *)y;
    for(size_t i=0; i<end; ++i) {
      vdst[i] = _mm256_sub_pd(vx[i], vy[i]);
    }
    for(size_t i=end*single_size; i<n; ++i) {
      dst[i] = x[i]-y[i];
    }
}
void avx_float64_multo(const size_t n, double *dst,double *x, double *y)
{
    static const size_t single_size = 4;
    const size_t end = n / single_size;
    __m256d *vdst =(__m256d *)dst;
    __m256d *vx = (__m256d *)x;
    __m256d *vy = (__m256d *)y;
    for(size_t i=0; i<end; ++i) {
      vdst[i] = _mm256_mul_pd(vx[i], vy[i]);
    }
    for(size_t i=end*single_size; i<n; ++i) {
      dst[i] = x[i]*y[i];
    }
}
void avx_float64_scaleto(const size_t n, double *dst,double *x, double scale)
{
    static const size_t single_size = 4;
    const size_t end = n / single_size;
    __m256d *vdst =(__m256d *)dst;
    __m256d *vx = (__m256d *)x;
    __m256d vy = _mm256_broadcast_sd(&scale);
    for(size_t i=0; i<end; ++i) {
      vdst[i] = _mm256_mul_pd(vx[i], vy);
    }
    for(size_t i=end*single_size; i<n; ++i) {
      dst[i] = x[i]*scale;
    }
}
*/
import (
	"C"
)

type F32 struct{}
type F64 struct{}

func (F32) Sum(a []float32) float32 {
	return float32(C.avx_float32_sum(C.ulong(len(a)), (*C.float)(&a[0])))
}

func (F32) Dot(x, y []float32) float32 {
	dot := C.avx_float32_dot((C.size_t)(len(x)), (*C.float)(&x[0]), (*C.float)(&y[0]))
	return float32(dot)
}

func (F32) AddTo(dst, x, y []float32) {
	C.avx_float32_addto((C.size_t)(len(dst)), (*C.float)(&dst[0]), (*C.float)(&x[0]), (*C.float)(&y[0]))
}

func (F32) SubTo(dst, x, y []float32) {
	C.avx_float32_subto((C.size_t)(len(dst)), (*C.float)(&dst[0]), (*C.float)(&x[0]), (*C.float)(&y[0]))
}

func (F32) MulTo(dst, x, y []float32) {
	C.avx_float32_multo((C.size_t)(len(dst)), (*C.float)(&dst[0]), (*C.float)(&x[0]), (*C.float)(&y[0]))
}
func (F32) ScaleTo(dst, x []float32, scale float32) {
	C.avx_float32_scaleto((C.size_t)(len(dst)), (*C.float)(&dst[0]), (*C.float)(&x[0]), (C.float)(scale))
}

func (F64) Sum(a []float64) float64 {
	return float64(C.avx_float64_sum(C.ulong(len(a)), (*C.double)(&a[0])))
}

func (F64) Dot(x, y []float64) float64 {
	dot := C.avx_float64_dot((C.size_t)(len(x)), (*C.double)(&x[0]), (*C.double)(&y[0]))
	return float64(dot)
}

func (F64) AddTo(dst, x, y []float64) {
	C.avx_float64_addto((C.size_t)(len(dst)), (*C.double)(&dst[0]), (*C.double)(&x[0]), (*C.double)(&y[0]))
}

func (F64) SubTo(dst, x, y []float64) {
	C.avx_float64_subto((C.size_t)(len(dst)), (*C.double)(&dst[0]), (*C.double)(&x[0]), (*C.double)(&y[0]))
}

func (F64) MulTo(dst, x, y []float64) {
	C.avx_float64_multo((C.size_t)(len(dst)), (*C.double)(&dst[0]), (*C.double)(&x[0]), (*C.double)(&y[0]))
}
func (F64) ScaleTo(dst, x []float64, scale float64) {
	C.avx_float64_scaleto((C.size_t)(len(dst)), (*C.double)(&dst[0]), (*C.double)(&x[0]), (C.double)(scale))
}
