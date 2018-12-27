// Package float offer same operations on float32 and float64 with some AVX2 operations if cgo enabled
package float

//go:generate go run _tools/tmpl/main.go -i -data float.gen.go.tmpldata float.gen.go.tmpl

//go:generate go run _tools/tmpl/main.go -i -data float.gen.go.tmpldata float_cgo.go.tmpl

//go:generate go run _tools/tmpl/main.go -i -data float.gen.go.tmpldata float_nocgo.go.tmpl

//go:generate go run _tools/tmpl/main.go -i -data float.gen.go.tmpldata float_test.go.tmpl

//go:generate go run _tools/tmpl/main.go -i -data float.gen.go.tmpldata powell.go.tmpl

//go:generate go run _tools/tmpl/main.go -i -data float.gen.go.tmpldata powell_test.go.tmpl

//go:generate go run _tools/tmpl/main.go -i -data float.gen.go.tmpldata brent.go.tmpl

//go:generate go run _tools/tmpl/main.go -i -data float.gen.go.tmpldata brent_test.go.tmpl

//go:generate clang -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -O3 -mavx -mavx2 -mavx512f -mfma -S -oC/avx.s C/avx.c

//go:generate c2goasm -a C/avx.s avx_amd64.s
