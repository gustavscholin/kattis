package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parsePoint(s string) struct{ x, y float64 } {
	point := strings.Split(s, " ")
	x, _ := strconv.ParseFloat(point[0], 64)
	y, _ := strconv.ParseFloat(point[1], 64)
	return struct{ x, y float64 }{x, y}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nbr_points, _ := strconv.Atoi(scanner.Text())
	for nbr_points != 0 {
		points := make([]struct{ x, y float64 }, nbr_points)
		area := 0.0
		for i := 0; i < nbr_points; i++ {
			if i == 0 {
				scanner.Scan()
				points[i] = parsePoint(scanner.Text())
			}
			if i < nbr_points-1 {
				scanner.Scan()
				points[i+1] = parsePoint(scanner.Text())
				area += (points[i+1].x - points[i].x) * (points[i+1].y + points[i].y) / 2
			} else {
				area += (points[0].x - points[i].x) * (points[0].y + points[i].y) / 2
			}
		}

		var s string
		if area < 0 {
			s = fmt.Sprintf("CCW %.1f", -area)
		} else {
			s = fmt.Sprintf("CW %.1f", area)
		}
		fmt.Println(s)
		scanner.Scan()
		nbr_points, _ = strconv.Atoi(scanner.Text())
	}
}
