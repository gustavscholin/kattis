package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	nbr := 1
	fmt.Println(nbr)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		nbr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if nbr == 99 {
			break
		}

		switch nbr % 3 {
		case 0:
			nbr += rand.Intn(2) + 1
		case 1:
			nbr += 2
		case 2:
			nbr += 1
		}

		fmt.Println(nbr)
		if nbr == 99 {
			break
		}
	}
}
