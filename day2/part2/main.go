package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	// file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	depth, position, aim := 0, 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		words := strings.Fields(scanner.Text())
		direction := words[0]
		magnitude, _ := strconv.Atoi(words[1])
		switch direction {
		case "forward":
			position += magnitude
			depth = depth + (aim * magnitude)
		case "backward":
			position -= magnitude
		case "up":
			aim -= magnitude
		case "down":
			aim += magnitude
		}
	}

	fmt.Printf("%v\n", depth)
	fmt.Printf("%v\n", position)
	fmt.Printf("%v\n", position*depth)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
