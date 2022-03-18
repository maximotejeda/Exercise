package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/maximotejeda/Exercise/parser"
)

// can be called with "go run main.go < containers.json"
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	file := []string{}
	text := ""

	for scanner.Scan() {
		file = append(file, scanner.Text())
	}

	text = strings.Join(file, "")
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, " ", "")

	fmt.Println(parser.ParseSTR(text))
	fmt.Println(parser.Rparser(text, nil, true))
}
