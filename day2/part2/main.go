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

// X == lose
// Y == draw
// Z == win``

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

func chooseMove(game_outcome int, player_opponent int) int {
	player_you := 0
	switch game_outcome {
	case 1: //lose
		if player_opponent == 1 {
			player_you = 3
		}
		if player_opponent == 2 {
			player_you = 1
		}
		if player_opponent == 3 {
			player_you = 2
		}
	case 2: //draw
		player_you = player_opponent
	case 3: //win
		if player_opponent == 1 {
			player_you = 2
		}
		if player_opponent == 2 {
			player_you = 3
		}
		if player_opponent == 3 {
			player_you = 1
		}
	}
	return player_you
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
		game_outcome := decrypt(line_list[1])
		player_opponent := decrypt(line_list[0])
		player_you := chooseMove(game_outcome, player_opponent)
		score_total += scoreRound(player_you, player_opponent)
	}
	fmt.Println(score_total)
}
