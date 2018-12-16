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
