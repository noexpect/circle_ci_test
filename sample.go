package main

import (
	"flag"
	"fmt"
)

func check_args(args []string) string {
	result := ""
	if len(args) > 0 {
		result = args[0]
	} else {
		result = "no args"
	}
	return result
}

func main() {
	flag.Parse()
	args := flag.Args()
	result := check_args(args)
	fmt.Println(result)

}
