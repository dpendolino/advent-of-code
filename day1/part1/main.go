package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var readings []int
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		num, _ := strconv.Atoi(scanner.Text())
		readings = append(readings, num)
	}

	fmt.Printf("%v\n", readings)

	for i, r := range readings {
		if i >= len(readings)-1 {
			break
		}
		if r < readings[i+1] {
			count++
		}
	}

	fmt.Printf("%d", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
