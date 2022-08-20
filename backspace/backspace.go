package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	text_list := []rune(text)
	var out []string

	for _, l := range text_list {
		if string(l) == "<" {
			if len(out) > 0 {
				out = out[:len(out)-1]
			}
		} else {
			out = append(out, string(l))
		}
	}

	out_string := strings.Join(out, "")
	if out_string != "" {
		fmt.Print(out_string)
	}
}
