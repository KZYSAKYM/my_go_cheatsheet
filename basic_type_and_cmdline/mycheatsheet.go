// In go, all source file belong to some package
// main package is special
// We can define our own package like myownpkg
package main

/* We can import external package by import statement.
 * We can access attributes of them like <package>.<attribute>.
 */
import (
	"flag" //line argparse
	"fmt"
	"math"
	"os"
)

/*
 * @flags: type, ptr, opt, default, usage
 */
func initializeFlag(flags [][]interface{}) bool {
	// set usage
	flag.Usage = func() {
		print("usage: this is my usage message\n")
		flag.PrintDefaults()
	}
	for i := 0; i < len(flags); i++ {
		typename := flags[i][0].(string)
		opt := flags[i][2].(string)
		usage := flags[i][4].(string)
		if typename == "bool" {
			ptr := flags[i][1].(*bool)
			defval := flags[i][3].(bool)
			flag.BoolVar(ptr, opt, defval, usage)
		} else if typename == "string" {
			ptr := flags[i][1].(*string)
			defval := flags[i][3].(string)
			flag.StringVar(ptr, opt, defval, usage)
		} else if typename == "int" {
			ptr := flags[i][1].(*int)
			defval := flags[i][3].(int)
			flag.IntVar(ptr, opt, defval, usage)
		} else {
			return false
		}
	}
	flag.Parse()
	return true
}

func main() {
	var (
		normal_int     = int(-1)
		normal_int_alt = -1
		normal_uint    = uint(0)
		int8_t         = int8(1)
		int16_t        = int16(2)
		int32_t        = int32(3)
		int64_t        = int64(4)
		uint8_t        = uint8(5)
		byte_t         = byte(6)
		uint16_t       = uint16(7)
		uint32_t       = uint32(8)
		uint64_t       = uint64(9)
		boolean        = true
		strings        = "string"
		normal_float   = 5.34
		float_t        = float32(3.141592)
		double_t       = float64(4.52)
		zero           = 0.0
		inf            = 1.0 / zero
		ninf           = -1.0 / zero
		NaN            = zero / zero
		complex128_t   = 6.364e-54i
		complex64_t    = complex64(1E8i)
		// we can extract real part by real() and imagine part by imag()
		utf32_char = '龍'
		raw_string = `raw string\n`
		array      = [][]string{
			{"normal_int", "int", fmt.Sprintf("%v", normal_int)},
			{"normal_int_alt", "int", fmt.Sprintf("%v", normal_int_alt)},
			{"normal_uint", "uint", fmt.Sprintf("%v", normal_uint)},
			{"int8_t", "int8", fmt.Sprintf("%v", int8_t)},
			{"int16_t", "int16", fmt.Sprintf("%v", int16_t)},
			{"int32_t", "int32", fmt.Sprintf("%v", int32_t)},
			{"int64_t", "int64", fmt.Sprintf("%v", int64_t)},
			{"uint8_t", "uint8", fmt.Sprintf("%v", uint8_t)},
			{"byte_t", "byte", fmt.Sprintf("%v", byte_t)},
			{"uint16_t", "uint16", fmt.Sprintf("%v", uint16_t)},
			{"uint32_t", "uint32", fmt.Sprintf("%v", uint32_t)},
			{"uint64_t", "uint64", fmt.Sprintf("%v", uint64_t)},
			{"boolean", "boolean", fmt.Sprintf("%v", boolean)},
			{"string", "string", fmt.Sprintf("%v", strings)},
			{"normal_float", "float64", fmt.Sprintf("%v", normal_float)},
			{"float_t", "float64", fmt.Sprintf("%v", float_t)},
			{"double_t", "float32", fmt.Sprintf("%v", double_t)},
			{"inf", "float64", fmt.Sprintf("%v", inf)},
			{"ninf", "float64", fmt.Sprintf("%v", ninf)},
			{"NaN", "float64", fmt.Sprintf("%v", NaN)},
			{"complex128_t", "complex128", fmt.Sprintf("%v", complex128_t)},
			{"complex64_t", "complex64", fmt.Sprintf("%v", complex64_t)},
			{"utf32_char", "rune", fmt.Sprintf("%v", utf32_char)},
			{"raw_string", "raw", fmt.Sprintf("%v", raw_string)},
			{"end", "end", fmt.Sprintf("%v", false)},
		}
		interface_t = []interface{}{
			5,
			1.09,
			"hoge",
			false,
			"We can set any typed value",
		}
	)

	fmt.Printf("var name :\t\t\t\ttype name :\t\t\t\tvalue\n")
	for i := 0; array[i][0] != "end"; i++ {
		fmt.Printf("%s :\t\t\t\t%s :\t\t\t\t%s\n", array[i][0], array[i][1], array[i][2])
	}
	// the scope of i is in for block.
	for i := 0; i < len(interface_t); i++ {
		fmt.Printf("interface_t :\t\t\t\tinterface{} :\t\t\t\t%v\n", interface_t[i])
	}

	fmt.Printf("int8: \n\tmax: %d \n\tmin: %d\n", math.MaxInt8, math.MinInt8)

	fmt.Printf("int16: \n\tmax: %d \n\tmin: %d\n", math.MaxInt16, math.MinInt16)

	fmt.Printf("int32: \n\tmax: %d \n\tmin: %d\n", math.MaxInt32, math.MinInt32)

	fmt.Printf("int64: \n\tmax: %d \n\tmin: %d\n", math.MaxInt64, math.MinInt64)

	fmt.Printf("uint8: \n\tmax: %d \n\tmin: %d\n", math.MaxUint8, 0)

	fmt.Printf("uint16: \n\tmax: %d \n\tmin: %d\n", math.MaxUint16, 0)

	fmt.Printf("uint32: \n\tmax: %d \n\tmin: %d\n", math.MaxUint32, 0)

	fmt.Printf("uint64: \n\tmax: %d \n\tmin: %d\n", uint64(math.MaxUint64), 0)

	fmt.Printf("float32: \n\tmax: %f\n\tmin: %f\n", math.MaxFloat32, -math.MaxFloat32)

	fmt.Printf("float64: \n\tmax: %f\n\tmin: %f\n", math.MaxFloat64, -math.MaxFloat64)

	// parse cmdline opts
	// prepare cmdline opt vars
	var (
		is_debug bool
		message  string
		value    int
	)
	flags := [][]interface{}{
		{"bool", &is_debug, "debug", false, "enable debug message"},
		{"string", &message, "message", "", "append your message into debug message"},
		{"int", &value, "value", 0, "set numeric val for debug message"},
		//{"uint", &value, "value", 0, "raise error in initializeFlag"},
	}
	if !initializeFlag(flags) {
		print("Error: invalid typename found\n")
		os.Exit(-1)
	}
	if is_debug {
		print("\nThis is a debug message outputted to stderr\n")
		print("\tYour message : ", message)
		print("\tYour value   : ", value)
	}
}
