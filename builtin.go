// Package pointer provides functions that returns pointer to basic type.
package pointer

//go:generate go run generate.go -type bool true false

//go:generate go run generate.go -type uint 0 1 ^uint(0)
//go:generate go run generate.go -type uint8 0 1 ^uint8(0)
//go:generate go run generate.go -type uint16 0 1 ^uint16(0)
//go:generate go run generate.go -type uint32 0 1 ^uint32(0)
//go:generate go run generate.go -type uint64 0 1 ^uint64(0)

//go:generate go run generate.go -type int -- -(^0>>1)-1 -1 0 1 (^0)>>1
//go:generate go run generate.go -type int8 -- -(1<<7) 0 1 ^int8(0)>>1
//go:generate go run generate.go -type int16 -- -(1<<15) -1 0 1 ^int16(0)>>1
//go:generate go run generate.go -type int32 -- -(1<<31) -1 0 1 ^int32(0)>>1
//go:generate go run generate.go -type int64 -- -(1<<63) -1 0 1 ^int64(0)>>1

// see go.dev/ref/spec: Floating-point/Imaginary/Rune/String literals
//go:generate go run generate.go -type float32 0. 72.40 6.67428e-11
//go:generate go run generate.go -type float64 0. 72.40 6.67428e-11
//go:generate go run generate.go -type complex64 0i 0123i 2.71828i 6.67428-11i
//go:generate go run generate.go -type complex128 0i 0123i 2.71828i 6.67428-11i
//go:generate go run generate.go -type byte 0 'a' '\007' '\t'
//go:generate go run generate.go -type rune 0 'a' 'ä' '本' '\t' '\007'
//go:generate go run generate.go -type string "\"\"" "\"hello, world!\\n\"" "\"日本語\""
