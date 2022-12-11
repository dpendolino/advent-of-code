package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// A == Rock
// B == Paper
// C = Scissors

// X == rock
// Y == Paper
// Z == Scissors

func decrypt(entry string) int {
	val := -1
	switch entry {
	case "A", "X":
		val = 1
	case "B", "Y":
		val = 2
	case "C", "Z":
		val = 3
	}
	return val
}

func scoreRound(player_you int, player_opponent int) int {
	score := player_you
	if player_you == player_opponent { // tie
		score += 3
	} else if (player_you == 1 && player_opponent == 3) || (player_you == 3 && player_opponent == 2) || (player_you == 2 && player_opponent == 1) { // wins
		score += 6
	} else if (player_you == 1 && player_opponent == 2) || (player_you == 3 && player_opponent == 1) || (player_you == 2 && player_opponent == 3) { // loses
		score += 0
	}
	return score
}

func main() {
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
	score_total := 0
	for _, line := range lines {
		line_list := strings.Fields(line)
		player_you := line_list[1]
		player_opponent := line_list[0]
		score_total += scoreRound(decrypt(player_you), decrypt(player_opponent))
	}
	fmt.Println(score_total)
}
