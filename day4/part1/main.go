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

	power_consumption := make([]string, 0)
	gamma, epsilon := make([]string, 0), make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		power_consumption = append(power_consumption, scanner.Text())
	}
	for i := 0; i < len(power_consumption[0]); i++ {
		one, zero := 0, 0
		for _, reading := range power_consumption {
			if string(reading[i]) == "0" {
				zero++
			} else {
				one++
			}
		}

		if zero > one {
			epsilon = append(epsilon, "1")
			gamma = append(gamma, "0")
		} else {
			epsilon = append(epsilon, "0")
			gamma = append(gamma, "1")
		}
	}
	gamma_int, err := strconv.ParseInt(strings.Join(gamma, ""), 2, 32)
	epsilon_int, err := strconv.ParseInt(strings.Join(epsilon, ""), 2, 32)
	fmt.Printf("gamma: %v\n", gamma)
	fmt.Printf("epsilon: %v\n", epsilon)
	fmt.Printf("%v\n", gamma_int*epsilon_int)
	fmt.Printf("power_consumption: %v\n", power_consumption)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
