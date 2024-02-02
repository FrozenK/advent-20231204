package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type card struct {
	wn        []string
	nl        []string
	num       int
	intersect int
	ID        int
}

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
	hasWiningTicket := false
	var cards []card
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		wn, nl := parseLine(line)
		cards = append(cards, card{
			wn: wn,
			nl: nl,
			ID: y,
		})
		count := countIntersection(wn, nl)
		cards[y-1].intersect = count
		if count > 0 {
			hasWiningTicket = true
			points := math.Exp2(float64(count - 1))
			sum += int(points)
		}
		y++
	}

	k := 0
	for hasWiningTicket {

		c := cards[k]

		if c.intersect > 0 {
			for i := c.ID; i < c.ID+c.intersect; i++ {
				if i >= y {
					break
				}

				cards = append(cards, cards[i])
			}
		}

		if len(cards) == k+1 {
			hasWiningTicket = false
		}
		k++
	}

	fmt.Printf("Sum = %d\n", sum)
	fmt.Printf("Total = %d\n", len(cards))
}
