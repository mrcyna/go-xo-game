package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	PC   = "PC"
	USER = "USER"
	MARS = "MARS"
	NONE = "NONE"
)

func main() {
	fmt.Println("*** Welcome to XO game ***")

RESET_GAME:
	rand.Seed(time.Now().Unix())

	board := make([]string, 9)
	for i := 0; i < len(board); i++ {
		board[i] = " "
	}

	turn := turn()

	for {
		displayBoard(board)

		finished, winner := finished(board)
		if finished {
			switch winner {
			case PC:
				fmt.Println("PC is the winner!")
			case USER:
				fmt.Println("You are the winner!")
			case MARS:
				fmt.Println("Game is finished But no one is the winner!")
			}

		PLAY_AGAIN_QUESTION:
			fmt.Println("Do you want to play agian? Y/n")
			var command string
			fmt.Scanf("%s\n", &command)

			switch {
			case command == "y" || command == "Y":
				goto RESET_GAME
			case command == "n" || command == "N":
				fmt.Println("Good Bye!")
				os.Exit(0)
			default:
				fmt.Println("Invalid command")
				goto PLAY_AGAIN_QUESTION
			}
		}

		if turn == PC {
			board[pcMove(board)] = "X"
			turn = USER
		} else {
		USER_INPUT:
			fmt.Println("Now It's your turn, Enter your number:")
			var choose int
			fmt.Scanf("%d\n", &choose)

			choose--

			if choose < 0 || choose > 9 {
				fmt.Println("Invalid input! (You have to enter a number between 0~9)")
				goto USER_INPUT
			}

			if board[choose] != " " {
				fmt.Println("Invalid input! (You have to enter the number of an empty cell)")
				goto USER_INPUT
			}

			board[choose] = "O"
			turn = PC
		}
	}
}

func finished(b []string) (bool, string) {
	if b[0] != " " && b[0] == b[1] && b[1] == b[2] {
		if b[0] == "X" {
			return true, PC
		} else {
			return true, USER
		}
	}

	if b[3] != " " && b[3] == b[4] && b[4] == b[5] {
		if b[3] == "X" {
			return true, PC
		} else {
			return true, USER
		}
	}

	if b[6] != " " && b[6] == b[7] && b[7] == b[8] {
		if b[6] == "X" {
			return true, PC
		} else {
			return true, USER
		}
	}

	if b[0] != " " && b[0] == b[3] && b[3] == b[6] {
		if b[0] == "X" {
			return true, PC
		} else {
			return true, USER
		}
	}

	if b[1] != " " && b[1] == b[4] && b[4] == b[7] {
		if b[1] == "X" {
			return true, PC
		} else {
			return true, USER
		}
	}

	if b[2] != " " && b[2] == b[5] && b[5] == b[8] {
		if b[2] == "X" {
			return true, PC
		} else {
			return true, USER
		}
	}

	if b[0] != " " && b[0] == b[4] && b[4] == b[8] {
		if b[0] == "X" {
			return true, PC
		} else {
			return true, USER
		}
	}

	if b[2] != " " && b[2] == b[4] && b[4] == b[6] {
		if b[2] == "X" {
			return true, PC
		} else {
			return true, USER
		}
	}

	emptyCells := make([]int, 0)

	for i := 0; i < len(b); i++ {
		if b[i] == " " {
			emptyCells = append(emptyCells, i)
		}
	}

	if len(emptyCells) == 0 {
		return true, MARS
	}

	return false, NONE
}

func rnd(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func turn() string {
	turn := rnd(0, 1)

	if turn%2 == 0 {
		return PC
	}
	return USER
}

func pcMove(board []string) int {

	emptyCells := make([]int, 0)

	for i := 0; i < len(board); i++ {
		if board[i] == " " {
			emptyCells = append(emptyCells, i)
		}
	}

	index := rnd(0, len(emptyCells)-1)

	return emptyCells[index]
}

func displayBoard(b []string) {

	print("\033[H\033[2J")

	fmt.Println(`
+-----------------------+
|1      |2      |3      |
|   ` + b[0] + `   |   ` + b[1] + `   |   ` + b[2] + `   |
|       |       |       |
+-------+-------+-------+
|4      |5      |6      |
|   ` + b[3] + `   |   ` + b[4] + `   |   ` + b[5] + `   |
|       |       |       |
+-------+-------+-------+
|7      |8      |9      |
|   ` + b[6] + `   |   ` + b[7] + `   |   ` + b[8] + `   |
|       |       |       |
+-----------------------+
    `)
}
