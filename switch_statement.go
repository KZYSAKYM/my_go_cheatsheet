package main

import (
	"fmt"
	"unsafe"
)

const (
	EVEN = 0x0
	ODD  = 0x1
)

func switch_numeric(v int) {
	switch v {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("fall through")
		fallthrough
	case 2, 3:
		fmt.Println("numeric", v)
	default:
		fmt.Println("Enter in default")
	}
}

func switch_strings(v string) {
	switch v {
	case "0":
		fmt.Println("zero")
	case "1":
		fmt.Println("fall through")
		fallthrough
	case "2", "3":
		fmt.Println("numeric", v)
	default:
		fmt.Println("Enter in default")
	}
}

func is_even(v int) (ret bool) {
	switch {
	case v % 2 == EVEN:
		ret = true
	case v % 2 == ODD:
		ret = false
	}
	return ret
}

func is_chara(v string) (ret bool) {
	ret = false
	switch length := len(v); length {
	case 0:
		fmt.Println("empty string")
	case 1:
		fmt.Println("character")
		ret = true
	default:
		fmt.Println("strings")
	}
	return ret
}

func str_idx(i int,v string) int {
	size   := int(unsafe.Sizeof(v[0]))
	return (i / size)
}

func main() {
	var (
		numeric = []int{0, 1, 2, 3}
		strings = []string{"0", "1", "2", "3"}
		msg string
	)

	for i,v := range numeric {
		fmt.Println("call switch_numeric: index:", i, "value:", v)
		switch_numeric(v)
		if is_even(v) {
			msg = "even"
		} else {
			msg = "odd"
		}
		fmt.Println("this is ", msg)
	}

	for i, v := range strings {
		fmt.Println("call switch_string: index:", str_idx(i, v), "value:", v)
		switch_strings(v)
		if is_chara(v) {
			msg = ""
		} else {
			msg = "not"
		}
		fmt.Println("this is ", msg, "character")
	}
}

