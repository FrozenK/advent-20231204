package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func parseLine(line string) ([]string, []string) {
	split := strings.Split(line, ":")
	card := strings.Split(split[1], "|")
	return strings.Split(card[0], " "), strings.Split(card[1], " ")
}

func countIntersection(wn []string, nl []string) int {
	count := 0
	for _, w := range wn {
		for _, n := range nl {
			if n != "" && n == w {
				count++
			}
		}
	}
	return count
}

func main() {
	f, err := os.Open("input2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	y := 1
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		wn, nl := parseLine(line)
		count := countIntersection(wn, nl)
		if count > 0 {
			points := math.Exp2(float64(count - 1))
			sum += int(points)
		}
		y++
	}
	fmt.Printf("Sum = %d\n", sum)
}
