#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include <immintrin.h> //AVX: -mavx
/*
void avx_float32_dot(const size_t n, float *x, float *y,float*res)
{
    static const size_t groupsize = 16;
    const size_t end = n / groupsize;
    __m512 *vx = (__m512 *)x;
    __m512 *vy = (__m512 *)y;
    __m512 vsum = {0};
    for(size_t i=0; i<end; ++i) {
      //vsum = _mm512_add_ps(vsum, _mm512_mul_ps(vx[i], vy[i]));
      vsum = _mm512_fmadd_ps(vx[i], vy[i], vsum);
    }
    __attribute__((aligned(32))) float t[16] = {0};
    _mm512_store_ps(t, vsum);
    for(size_t i=end*groupsize; i<n; ++i) {
      t[i&7] += x[i]*y[i];
    }
    *res= t[0] + t[1] + t[2] + t[3] + t[4] + t[5] + t[6] + t[7] + t[8] + t[9] + t[10] + t[11] + t[12] + t[13] + t[14] + t[15];
}
*/
void avxFloat32Dot(const size_t n, float *x, float *y, float*res)
{
    static const size_t groupsize = 8;
    const size_t end = n / groupsize;
    __m256 vsum={0};
    for(size_t i=0; i<end*groupsize; i+=groupsize) {
      //vsum = _mm256_add_ps(vsum, _mm256_mul_ps(vx[i], vy[i]));

      vsum = _mm256_fmadd_ps(_mm256_loadu_ps(x+i), _mm256_loadu_ps(y+i), vsum);
    }
    __attribute__((aligned(32))) float t[8] = {0};
    _mm256_store_ps(t, vsum);
    *res= t[0] + t[1] + t[2] + t[3] + t[4] + t[5] + t[6] + t[7];
    for(size_t i=end*groupsize; i<n; i++) {
      *res += x[i]*y[i];
    }
}

void avxFloat64Dot(const size_t n, double *x, double *y,double*res)
{
    static const size_t groupsize = 4;
    const size_t end = n / groupsize;
    __m256d vsum = {0};
    for(size_t i=0; i<end*groupsize; i+=groupsize) {
      //vsum = _mm256_add_pd(vsum, _mm256_mul_pd(vx[i], vy[i]));
      vsum = _mm256_fmadd_pd(_mm256_loadu_pd(x+i), _mm256_loadu_pd(y+i),vsum);
    }
    __attribute__((aligned(32))) double t[4] = {0};
    _mm256_store_pd(t, vsum);
    *res= t[0] + t[1] + t[2] + t[3];
    for(size_t i=end*groupsize; i<n; i++) {
      *res += x[i]*y[i];
    }
}

/*
void avx_float32_addto(const size_t n, float *x, float *y,float*res)
{
    static const size_t groupsize = 8;
    const size_t end = n / groupsize;
    __m256 *vx = (__m256 *)x;
    __m256 *vy = (__m256 *)y;
    __m256 *vsum = (__m256 *)res;
    __m256 tmp;
    for(size_t i=0; i<end; ++i) {
      tmp = _mm256_add_ps(vx[i], vy[i]);
      vsum[i] =tmp;
    }

    for(size_t i=end*groupsize; i<n; ++i) {
      res[i] = x[i]+y[i];
    }
}

void avx_float64_addto(const size_t n, double *x, double *y,double*res)
{
    static const size_t groupsize = 4;
    const size_t end = n / groupsize;
    __m256d *vx = (__m256d *)x;
    __m256d *vy = (__m256d *)y;
    __m256d *vsum = (__m256d *)res;
    __m256d tmp;
    for(size_t i=0; i<end; ++i) {
      tmp =_mm256_add_pd(vx[i], vy[i]);
      vsum[i] = tmp;
    }
    for(size_t i=end*groupsize; i<n; ++i) {
      res[i] = x[i]+y[i];
    }
}
*/