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

func search(arr []int, l int, r int, key int) int {
	for r-l > 1 {
		m := l + (r-l)/2
		if arr[m] >= key {
			r = m
		} else {
			l = m
		}
	}
	return r
}

func lis(arr []int) int {
	n := len(arr)
	ends := make([]int, n)
	ends[0] = arr[0]

	len := 1
	for i := 1; i < n; i++ {
		if arr[i] < ends[0] {
			ends[0] = arr[i]
		} else if arr[i] > ends[len-1] {
			ends[len] = arr[i]
			len++
		} else {
			ends[search(ends, -1, len-1, arr[i])] = arr[i]
		}
	}

	return len
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nbrPorts := scanInt(scanner)

	ports := make([]int, nbrPorts)
	for i := 0; i < nbrPorts; i++ {
		port := scanInt(scanner)
		ports[i] = port
	}

	fmt.Println(lis(ports))
}
