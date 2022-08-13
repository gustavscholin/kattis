package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func queryNames(start string, end string, genderSuitability int, names map[string]int) int {
	matches := 0
	for name, gender := range names {
		if (genderSuitability == 0 || gender == genderSuitability) && (start <= name && name < end) {
			matches++
		}
	}
	return matches
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commandParts := strings.Split(scanner.Text(), " ")
	command, _ := strconv.Atoi(commandParts[0])
	names := make(map[string]int)
	matches := 0
	for command != 0 {
		switch command {
		case 1:
			gender, _ := strconv.Atoi(commandParts[2])
			names[commandParts[1]] = gender
		case 2:
			delete(names, commandParts[1])
		case 3:
			genderSuitability, _ := strconv.Atoi(commandParts[3])
			matches = 0
			for name, gender := range names {
				if (genderSuitability == 0 || gender == genderSuitability) && (commandParts[1] <= name && name < commandParts[2]) {
					matches++
				}
			}
			fmt.Println(matches)
		}
		scanner.Scan()
		commandParts = strings.Split(scanner.Text(), " ")
		command, _ = strconv.Atoi(commandParts[0])
	}
}
