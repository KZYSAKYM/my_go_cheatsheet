package main

import (
	"fmt"
)

func type_generic_simple(v interface{}) {
	var type_name string
	_, isInt := v.(int)
	if (isInt) {
		type_name = "int"
	}
	_, isString := v.(string)
	if (isString) {
		type_name = "string"
	}
	_, isBool := v.(bool)
	if (isBool) {
		type_name = "bool"
	}
	fmt.Println("type:", type_name)
}

func type_generic_switch(v interface{}) {
	var type_name string
	switch val := v.(type) {
	case int:
		type_name = "int"
		fmt.Println("val:", val)
	case string:
		type_name = "string"
		fmt.Println("val:", val)
	case uint:
		type_name = "uint"
		fmt.Println("val:", val)
	case bool:
		type_name = "bool"
		fmt.Println("val:", val)
	}
	fmt.Println("type:", type_name)
}

func main() {
	var (
		integer = -90
		unsigned = 100
		str = "foo"
		boolean = true
	)
	type_generic_simple(integer)
	type_generic_simple(unsigned)
	type_generic_simple(str)
	type_generic_simple(boolean)
	type_generic_switch(integer)
	type_generic_switch(unsigned)
	type_generic_switch(str)
	type_generic_switch(boolean)
}
