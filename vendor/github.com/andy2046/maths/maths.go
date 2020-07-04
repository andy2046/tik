package maths

import (
	"math"
)

type (
	// String represents type string.
	String string
	// Int represents type int.
	Int int
	// Int8 represents type int8.
	Int8 int8
	// Int16 represents type int16.
	Int16 int16
	// Int32 represents type int32.
	Int32 int32
	// Int64 represents type int64.
	Int64 int64
	// Uint represents type uint.
	Uint uint
	// Uint8 represents type uint8.
	Uint8 uint8
	// Uint16 represents type uint16.
	Uint16 uint16
	// Uint32 represents type uint32.
	Uint32 uint32
	// Uint64 represents type uint64.
	Uint64 uint64
	// Byte represents type byte.
	Byte byte
	// Rune represents type rune.
	Rune rune
	// Float32 represents type float32.
	Float32 float32
	// Float64 represents type float64.
	Float64 float64
)

var (
	// StringVar represents instance of string.
	StringVar = new(String)
	// IntVar represents instance of int.
	IntVar = new(Int)
	// Int8Var represents instance of int8.
	Int8Var = new(Int8)
	// Int16Var represents instance of int16.
	Int16Var = new(Int16)
	// Int32Var represents instance of int32.
	Int32Var = new(Int32)
	// Int64Var represents instance of int64.
	Int64Var = new(Int64)
	// UintVar represents instance of uint.
	UintVar = new(Uint)
	// Uint8Var represents instance of uint8.
	Uint8Var = new(Uint8)
	// Uint16Var represents instance of uint16.
	Uint16Var = new(Uint16)
	// Uint32Var represents instance of uint32.
	Uint32Var = new(Uint32)
	// Uint64Var represents instance of uint64.
	Uint64Var = new(Uint64)
	// ByteVar represents instance of byte.
	ByteVar = new(Byte)
	// RuneVar represents instance of rune.
	RuneVar = new(Rune)
	// Float32Var represents instance of float32.
	Float32Var = new(Float32)
	// Float64Var represents instance of float64.
	Float64Var = new(Float64)
)

// Range creates a range progressing from zero up to, but not including end.
func Range(end int) []struct{} {
	return make([]struct{}, end)
}

// Max returns the largest in numerics.
func (Float32) Max(numerics ...float32) float32 {
	n := float64(numerics[0])
	for i := 1; i < len(numerics); i++ {
		n = math.Max(float64(numerics[i]), n)
	}
	return float32(n)
}

// Min returns the smallest in numerics.
func (Float32) Min(numerics ...float32) float32 {
	n := float64(numerics[0])
	for i := 1; i < len(numerics); i++ {
		n = math.Min(float64(numerics[i]), n)
	}
	return float32(n)
}

// Max returns the largest in numerics.
func (Float64) Max(numerics ...float64) float64 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		n = math.Max(numerics[i], n)
	}
	return n
}

// Min returns the smallest in numerics.
func (Float64) Min(numerics ...float64) float64 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		n = math.Min(numerics[i], n)
	}
	return n
}

// Max returns the largest in numerics.
func (Rune) Max(numerics ...rune) rune {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (Rune) Min(numerics ...rune) rune {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Max returns the largest in numerics.
func (Byte) Max(numerics ...byte) byte {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (Byte) Min(numerics ...byte) byte {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Max returns the largest in numerics.
func (Uint64) Max(numerics ...uint64) uint64 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (Uint64) Min(numerics ...uint64) uint64 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Max returns the largest in numerics.
func (Uint32) Max(numerics ...uint32) uint32 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (Uint32) Min(numerics ...uint32) uint32 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Max returns the largest in numerics.
func (Uint16) Max(numerics ...uint16) uint16 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (Uint16) Min(numerics ...uint16) uint16 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Max returns the largest in numerics.
func (Uint8) Max(numerics ...uint8) uint8 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (Uint8) Min(numerics ...uint8) uint8 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Max returns the largest in numerics.
func (Uint) Max(numerics ...uint) uint {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (Uint) Min(numerics ...uint) uint {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Max returns the largest in numerics.
func (Int64) Max(numerics ...int64) int64 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (Int64) Min(numerics ...int64) int64 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Max returns the largest in numerics.
func (Int32) Max(numerics ...int32) int32 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (Int32) Min(numerics ...int32) int32 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Max returns the largest in numerics.
func (Int16) Max(numerics ...int16) int16 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (Int16) Min(numerics ...int16) int16 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Max returns the largest in numerics.
func (Int8) Max(numerics ...int8) int8 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (Int8) Min(numerics ...int8) int8 {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Max returns the largest in numerics.
func (Int) Max(numerics ...int) int {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (Int) Min(numerics ...int) int {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Max returns the largest in numerics.
func (String) Max(numerics ...string) string {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > n {
			n = numerics[i]
		}
	}
	return n
}

// Min returns the smallest in numerics.
func (String) Min(numerics ...string) string {
	n := numerics[0]
	for i := 1; i < len(numerics); i++ {
		if numerics[i] < n {
			n = numerics[i]
		}
	}
	return n
}

// Ternary for no ternary operation in Go.
func Ternary(expr bool, trueVal, falseVal interface{}) interface{} {
	if expr {
		return trueVal
	}
	return falseVal
}
