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

	var readings, windows []int
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		num, _ := strconv.Atoi(scanner.Text())
		readings = append(readings, num)
	}

	for i, _ := range readings {
		if i >= len(readings)-2 { // account for 3 value window
			break
		}
		windows = append(windows, readings[i]+readings[i+1]+readings[i+2])
	}
	for i, w := range windows {
		if i >= len(windows)-1 {
			break
		}
		if windows[i+1] > w {
			count++
		}
	}

	fmt.Printf("%v\n", readings)
	fmt.Printf("%v\n", windows)
	fmt.Printf("%d", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
