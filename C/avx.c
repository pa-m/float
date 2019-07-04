#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include <immintrin.h> //AVX: -mavx
/*
void avxFloat32Dot(const size_t n, float *x, float *y, float*res)
{
    static const size_t groupsize = 16;
    const size_t end = n / groupsize;
    __m512 vsum={0};
    for(size_t i=0; i<end*groupsize; i+=groupsize) {
      vsum = _mm512_fmadd_ps(_mm512_loadu_ps(x+i), _mm512_loadu_ps(y+i), vsum);
    }
    __attribute__((aligned(64))) float t[16] = {0};
    _mm512_store_ps(t, vsum);
    *res= t[0] + t[1] + t[2] + t[3] + t[4] + t[5] + t[6] + t[7] +
    t[8] + t[9] + t[10] + t[11] + t[12] + t[13] + t[14] + t[15];
    for(size_t i=end*groupsize; i<n; i++) {
      *res += x[i]*y[i];
    }
}
*/
void avxFloat32Dot(const size_t n, float *x, float *y, float*res)
{
    static const int unroll=6;
    static const size_t regsize = 8; // a mm256 contains 8 float32
    static const size_t loopinc=unroll*regsize;
    const size_t end = (n / loopinc) * loopinc;
    __m256 vsum0={0},vsum1={0},vsum2={0},vsum3={0},vsum4={0},vsum5={0};
    for(size_t i=0; i<end; i+=loopinc) {
      vsum0 = _mm256_fmadd_ps(_mm256_loadu_ps(x), _mm256_loadu_ps(y), vsum0);
      x+=regsize;y+=regsize;
      vsum1 = _mm256_fmadd_ps(_mm256_loadu_ps(x), _mm256_loadu_ps(y), vsum1);
      x+=regsize;y+=regsize;
      vsum2 = _mm256_fmadd_ps(_mm256_loadu_ps(x), _mm256_loadu_ps(y), vsum2);
      x+=regsize;y+=regsize;
      vsum3 = _mm256_fmadd_ps(_mm256_loadu_ps(x), _mm256_loadu_ps(y), vsum3);
      x+=regsize;y+=regsize;
      vsum4 = _mm256_fmadd_ps(_mm256_loadu_ps(x), _mm256_loadu_ps(y), vsum4);
      x+=regsize;y+=regsize;
      vsum5 = _mm256_fmadd_ps(_mm256_loadu_ps(x), _mm256_loadu_ps(y), vsum5);
      x+=regsize;y+=regsize;
    }
    vsum0 = _mm256_add_ps(vsum0, vsum1);
    vsum2 = _mm256_add_ps(vsum2, vsum3);
    vsum4 = _mm256_add_ps(vsum4, vsum5);

    vsum0 = _mm256_add_ps(vsum0, vsum2);
    vsum0 = _mm256_add_ps(vsum0, vsum4);
    __attribute__((aligned(32))) float t[8] = {0};
    _mm256_store_ps(t, vsum0);
    *res= t[0] + t[1] + t[2] + t[3] + t[4] + t[5] + t[6] + t[7];
    for(size_t i=end; i<n; i++) {
      *res += *x * *y;
      x++;y++;
    }
}

void avxFloat64Dot(const size_t n, double *x, double *y,double*res)
{
    static const int unroll=6;
    static const size_t regsize = 4; // a mm256 contains 4 double
    const int loopinc=unroll*regsize;
    const size_t end = (n / loopinc) * loopinc;
    __m256d vsum0 = {0}, vsum1 = {0}, vsum2 = {0}, vsum3 = {0}, vsum4 = {0}, vsum5 = {0};
    if( (((unsigned)x & 0x1f)==0) && (((unsigned)y & 0x1f)==0)){
        for(size_t i=0; i<end; i+=loopinc) {
          //vsum = _mm256_add_pd(vsum, _mm256_mul_pd(vx[i], vy[i]));
          vsum0 = _mm256_fmadd_pd(_mm256_load_pd(x), _mm256_load_pd(y),vsum0);
          x+=regsize;y+=regsize;
          vsum1 = _mm256_fmadd_pd(_mm256_load_pd(x), _mm256_load_pd(y),vsum1);
          x+=regsize;y+=regsize;
          vsum2 = _mm256_fmadd_pd(_mm256_load_pd(x), _mm256_load_pd(y),vsum2);
          x+=regsize;y+=regsize;
          vsum3 = _mm256_fmadd_pd(_mm256_load_pd(x), _mm256_load_pd(y),vsum3);
          x+=regsize;y+=regsize;
          vsum4 = _mm256_fmadd_pd(_mm256_load_pd(x), _mm256_load_pd(y),vsum4);
          x+=regsize;y+=regsize;
          vsum5 = _mm256_fmadd_pd(_mm256_load_pd(x), _mm256_load_pd(y),vsum5);
          x+=regsize;y+=regsize;
        }
    }else{
        for(size_t i=0; i<end; i+=loopinc) {
          //vsum = _mm256_add_pd(vsum, _mm256_mul_pd(vx[i], vy[i]));
          vsum0 = _mm256_fmadd_pd(_mm256_loadu_pd(x), _mm256_loadu_pd(y),vsum0);
          x+=regsize;y+=regsize;
          vsum1 = _mm256_fmadd_pd(_mm256_loadu_pd(x), _mm256_loadu_pd(y),vsum1);
          x+=regsize;y+=regsize;
          vsum2 = _mm256_fmadd_pd(_mm256_loadu_pd(x), _mm256_loadu_pd(y),vsum2);
          x+=regsize;y+=regsize;
          vsum3 = _mm256_fmadd_pd(_mm256_loadu_pd(x), _mm256_loadu_pd(y),vsum3);
          x+=regsize;y+=regsize;
          vsum4 = _mm256_fmadd_pd(_mm256_loadu_pd(x), _mm256_loadu_pd(y),vsum4);
          x+=regsize;y+=regsize;
          vsum5 = _mm256_fmadd_pd(_mm256_loadu_pd(x), _mm256_loadu_pd(y),vsum5);
          x+=regsize;y+=regsize;
        }
    }
    vsum0 = _mm256_add_pd(vsum0, vsum1);
    vsum2 = _mm256_add_pd(vsum2, vsum3);
    vsum4 = _mm256_add_pd(vsum4, vsum5);

    vsum0 = _mm256_add_pd(vsum0, vsum2);
    vsum0 = _mm256_add_pd(vsum0, vsum4);
    __attribute__((aligned(32))) double t[4] = {0};
    _mm256_store_pd(t, vsum0);
    *res= t[0] + t[1] + t[2] + t[3];
    for(size_t i=end; i<n; i++) {
      *res += *x * *y;
      x++;y++;
    }
}

void avxFloat32Sum(const size_t n, float *x,float *res) {
    static const int unroll = 6;
    static const size_t regsize = 8;
    static const size_t loopinc = regsize * unroll;
    const size_t end = (n / loopinc)*loopinc;
    __m256 vsum0={0}, vsum1={0}, vsum2={0}, vsum3={0}, vsum4={0}, vsum5={0};
    for(size_t i=0; i<end; i+=loopinc) {
      vsum0 = _mm256_add_ps(_mm256_loadu_ps(x), vsum0);
      x+=regsize;
      vsum1 = _mm256_add_ps(_mm256_loadu_ps(x), vsum1);
      x+=regsize;
      vsum2 = _mm256_add_ps(_mm256_loadu_ps(x), vsum2);
      x+=regsize;
      vsum3 = _mm256_add_ps(_mm256_loadu_ps(x), vsum3);
      x+=regsize;
      vsum4 = _mm256_add_ps(_mm256_loadu_ps(x), vsum4);
      x+=regsize;
      vsum5 = _mm256_add_ps(_mm256_loadu_ps(x), vsum5);
      x+=regsize;
    }
    vsum0 = _mm256_add_ps(vsum0, vsum1);
    vsum2 = _mm256_add_ps(vsum2, vsum3);
    vsum4 = _mm256_add_ps(vsum4, vsum5);

    vsum0 = _mm256_add_ps(vsum0, vsum2);
    vsum0 = _mm256_add_ps(vsum0, vsum4);
    __attribute__((aligned(32))) float t[8] = {0};
    _mm256_store_ps(t, vsum0);
    *res= t[0] + t[1] + t[2] + t[3] + t[4] + t[5] + t[6] + t[7];
    for(size_t i=end; i<n; i++) {
      *res += *x;
      x++;
    }
}

void avxFloat64Sum(const size_t n, double *x, double*res)
{
    static const int unroll=6;
    static const size_t regsize = 4;
    const size_t end = (n / regsize) * regsize;
    __m256d vsum0 = {0}, vsum1 = {0}, vsum2 = {0}, vsum3 = {0}, vsum4 = {0}, vsum5 = {0};
    for(size_t i=0; i<end; i+=regsize) {
      vsum0 = _mm256_add_pd(_mm256_loadu_pd(x),vsum0);
      x+=regsize;
      vsum1 = _mm256_add_pd(_mm256_loadu_pd(x),vsum1);
      x+=regsize;
      vsum2 = _mm256_add_pd(_mm256_loadu_pd(x),vsum2);
      x+=regsize;
      vsum3 = _mm256_add_pd(_mm256_loadu_pd(x),vsum3);
      x+=regsize;
      vsum4 = _mm256_add_pd(_mm256_loadu_pd(x),vsum4);
      x+=regsize;
      vsum5 = _mm256_add_pd(_mm256_loadu_pd(x),vsum5);
      x+=regsize;
    }
    vsum0 = _mm256_add_pd(vsum0, vsum1);
    vsum2 = _mm256_add_pd(vsum2, vsum3);
    vsum4 = _mm256_add_pd(vsum4, vsum5);

    vsum0 = _mm256_add_pd(vsum0, vsum2);
    vsum0 = _mm256_add_pd(vsum0, vsum4);
    __attribute__((aligned(32))) double t[4] = {0};
    _mm256_store_pd(t, vsum0);
    *res= t[0] + t[1] + t[2] + t[3];
    for(size_t i=end; i<n; i++) {
      *res += *x;
      x++;
    }
}
/*
void avxFloat64Sum(const size_t n, double *x, double*res)
{
    static const int unroll=6;
    static const size_t regsize = 8;
    const size_t end = (n / regsize) * regsize;
    __m512d vsum0 = {0}, vsum1 = {0}, vsum2 = {0}, vsum3 = {0}, vsum4 = {0}, vsum5 = {0};
    for(size_t i=0; i<end; i+=regsize) {
      vsum0 = _mm512_add_pd(_mm512_loadu_pd(x),vsum0);
      x+=regsize;
      vsum1 = _mm512_add_pd(_mm512_loadu_pd(x),vsum1);
      x+=regsize;
      vsum2 = _mm512_add_pd(_mm512_loadu_pd(x),vsum2);
      x+=regsize;
      vsum3 = _mm512_add_pd(_mm512_loadu_pd(x),vsum3);
      x+=regsize;
      vsum4 = _mm512_add_pd(_mm512_loadu_pd(x),vsum4);
      x+=regsize;
      vsum5 = _mm512_add_pd(_mm512_loadu_pd(x),vsum5);
      x+=regsize;
    }
    vsum0 = _mm512_add_pd(vsum0, vsum1);
    vsum2 = _mm512_add_pd(vsum2, vsum3);
    vsum4 = _mm512_add_pd(vsum4, vsum5);

    vsum0 = _mm512_add_pd(vsum0, vsum2);
    vsum0 = _mm512_add_pd(vsum0, vsum4);
    __attribute__((aligned(32))) double t[8] = {0};
    _mm512_store_pd(t, vsum0);
    *res= t[0] + t[1] + t[2] + t[3] + t[4] + t[5] + t[6] + t[7];
    for(size_t i=end; i<n; i++) {
      *res += *x;
      x++;
    }
}
*/
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