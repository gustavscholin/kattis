package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func scanInt(s *bufio.Scanner) int {
	s.Scan()
	scanned, _ := strconv.Atoi(s.Text())
	return scanned
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	time := scanInt(scanner)
	nbrDishes := scanInt(scanner)

	dishTimes := make([]int, nbrDishes)
	for i := 0; i < nbrDishes; i++ {
		dish := scanInt(scanner)
		dishTimes[i] = dish
	}

	count := make([]int, time+1)
	count[0] = 1

	for i := 1; i <= time; i++ {
		for _, dishTime := range dishTimes {
			if i >= dishTime {
				count[i] += count[i-dishTime]
			}
		}
	}

	fmt.Println(count[time])
}
