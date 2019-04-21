package main

import (
	"fmt"
	"time"
)

var is_end bool = false

func parallel_echo(msg string) {
	if msg == "" {
		is_end = true
		return
	}
	fmt.Println(msg)
}

func main() {
	var (
		msgs = []string{
			"hoge",
			"hage",
			"foo",
			"bar",
			"",
		}
	)
	fmt.Println("start async messaging")

	for i, v := range msgs {
		fmt.Println("No.", i)
		go parallel_echo(v)
	}
	for is_end == false {
		time.Sleep(1 * time.Second)
		fmt.Printf(".")
	}
	fmt.Println("")
	fmt.Println("Done")
}
