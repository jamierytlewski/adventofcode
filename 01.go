package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs/01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var start []int
	var end []int

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		lines := strings.Fields(string(line))
		startInt, err := strconv.Atoi(lines[0])
		if err != nil {
			log.Fatal(err)
		}
		start = append(start, startInt)
		endInt, err := strconv.Atoi(lines[1])
		if err != nil {
			log.Fatal(err)
		}
		end = append(end, endInt)
	}
	sort.Ints(start)
	sort.Ints(end)

	var distance int = 0
	for i := range start {
		distance += int(math.Abs(float64(start[i] - end[i])))
	}
	fmt.Println(distance)
}
