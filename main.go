package main

import (
	"fmt"

	"github.com/maximotejeda/Exercise/parser"
)

func main() {
	fmt.Println(parser.ParseSTR("maximo(maximo){config{[(])]}}"))
}
