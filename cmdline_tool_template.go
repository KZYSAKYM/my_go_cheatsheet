package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	BOOL_OPT   = 1 + iota
	STRING_OPT
	INT_OPT
)

// TODO: LOW_PRIO:  Support LISENCE Identifier
const (
	PG_NAME = "Please input program name"
	USAGE   = "Please input usage"
	VERSION = "0.0.1"
)

// prepare cmdline opt vars
var (
	is_help     bool
	is_version  bool
	is_verbose  bool
)

func show_help() {
	flag.Usage()
	os.Exit(-1)
}

func show_version() {
	fmt.Println(VERSION)
}

func verbose_info(fmt_str string, args []interface{}) {
	if is_verbose {
		fmt.Println(fmt_str, args)
	}
}

/*
 * @flags: type, ptr, opt, default, usage
 * @return: ret bool: is process success
 *          err int:  error code
 */
func initializeFlag(flags [][]interface{}) (ret bool, err int) {
	ret = true
	err = 0
	// set usage
	flag.Usage = func() {
		show_version()
		print("usage: ", PG_NAME, USAGE, "\n")
		flag.PrintDefaults()
	}
	for i := 0; i < len(flags); i++ {
		typename := flags[i][0].(int)
		opt := flags[i][2].(string)
		usage := flags[i][4].(string)
		switch typename {
		case BOOL_OPT:
			ptr := flags[i][1].(*bool)
			defval := flags[i][3].(bool)
			flag.BoolVar(ptr, opt, defval, usage)
			break
		case STRING_OPT:
			ptr := flags[i][1].(*string)
			defval := flags[i][3].(string)
			flag.StringVar(ptr, opt, defval, usage)
			break
		case INT_OPT:
			ptr := flags[i][1].(*int)
			defval := flags[i][3].(int)
			flag.IntVar(ptr, opt, defval, usage)
			break
		default:
			ret = false
			err = -1
			break
		}
	}
	flag.Parse()
	return
}

func main() {
	flags := [][]interface{}{
		{
			BOOL_OPT, &is_verbose, "-verbose",
			false, "show verbose info message",
		},
		{
			BOOL_OPT, &is_help, "-help", false,
			"show usage",
		},
		{
			BOOL_OPT, &is_version, "-version", false,
			"show version",
		},
	}
	ret, err := initializeFlag(flags)
	if is_help {
		show_help()
	}
	if is_version {
		show_version()
		os.Exit(-1)
	}
	fmt.Println("ret:", ret)
	fmt.Println("err:", err)
}
