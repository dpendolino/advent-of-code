package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	elves := make([]int, 0)
	lines := make([]string, 0)
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elf := 0
	maxCalories := 0
	for _, line := range lines {
		// fmt.Println(elves)
		if line == "" {
			elf++
		} else {
			tempInt, _ := strconv.Atoi(line)
			if len(elves) > elf {
				elves[elf] += tempInt
			} else {
				elves = append(elves, tempInt)
			}
			// fmt.Println(len(elves), elf, elves)
			if maxCalories == 0 {
				maxCalories = elves[elf]
			} else if maxCalories < elves[elf] {
				maxCalories = elves[elf]
			}
		}
	}
	fmt.Printf("maxCalories=%d", maxCalories)
}
