package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nbr_commands int
	fmt.Scanln(&nbr_commands)

	volume := 7

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < nbr_commands; i++ {
		scanner.Scan()
		command := scanner.Text()
		if command == "Skru op!" && volume < 10 {
			volume++
		} else if command == "Skru ned!" && volume > 0 {
			volume--
		}
	}
	fmt.Println(volume)
}
