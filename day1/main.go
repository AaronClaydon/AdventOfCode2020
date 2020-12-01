package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

func main() {
	target := 2020

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []int

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
	}

	file.Close()

	p1err, p1num1, p1num2 := part1(numbers, target)

	if p1err != nil {
		log.Fatalln(p1err)
	}

	log.Printf("Found pair (%d,%d) multiplied is %d\n", p1num1, p1num2, (p1num1 * p1num2))

	p2err, p2num1, p2num2, p2num3 := part2(numbers, target)

	if p2err != nil {
		log.Fatalln(p2err)
	}

	log.Printf("Found triplet (%d,%d,%d) multiplied is %d", p2num1, p2num2, p2num3, (p2num1 * p2num2 * p2num3))
}

// Brute force with no optimisations but who cares
//   could remove comibations we've already tried
func part1(numbers []int, target int) (error, int, int) {
	for _, num1 := range numbers {
		for _, num2 := range numbers {
			if num1+num2 == target {
				return nil, num1, num2
			}
		}
	}

	return errors.New("Could not find pair"), -1, -1
}

func part2(numbers []int, target int) (error, int, int, int) {
	// any number more than 1000 isn't likely to be part of the answer
	//   filter it from our search space
	numbersFiltered := []int{}
	for _, num := range numbers {
		if num < 1000 {
			numbersFiltered = append(numbersFiltered, num)
		}
	}

	searchHits := 0
	for _, num1 := range numbersFiltered {
		for _, num2 := range numbersFiltered {
			for _, num3 := range numbersFiltered {
				searchHits += 1
				if num1+num2+num3 == target {
					// log.Println(searchHits)
					return nil, num1, num2, num3
				}
			}
		}
	}

	return errors.New("Could not find triplet"), -1, -1, -1
}
