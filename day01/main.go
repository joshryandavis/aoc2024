package main

import (
	util "aoc2024go/utils"
	_ "embed"
	"sort"
	"strings"
)

//go:embed data.txt
var data string

func main() {
	input := strings.Split(strings.TrimSpace(data), "\n")

	// part 1
	var left []int
	var right []int
	for _, line := range input {
		n := strings.SplitN(line, "   ", 2)
		left = append(left, util.ToInt(n[0]))
		right = append(right, util.ToInt(n[1]))
	}
	sort.Ints(left)
	sort.Ints(right)
	var distance int
	for i, leftN := range left {
		rightN := right[i]
		d := leftN - rightN
		if d < 0 {
			d = -d
		}
		distance += d
	}
	println("distance:", distance)

	// part 2
	similarity := 0
	for _, leftN := range left {
		count := 0
		for _, rightN := range right {
			if leftN == rightN {
				count++
			}
		}
		similarity += leftN * count
	}
	println("similarity:", similarity)
}
