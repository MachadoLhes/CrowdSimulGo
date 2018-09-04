package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size int = 15

type point struct {
	x     int
	y     int
	isOut bool
	next  *point
}

func main() {

	board := [size][size]int{}

	board = createBorder(board)
	board = populate(board)

	printBoard(board)
}

func checkNeighborhood(board [size][size]int, point point) bool {
	if board[point.x][point.y] != 0 {
		return false
	}
	return true
}

func salaVazia(board [size][size]int) bool {
	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			if board[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func populate(board [size][size]int) [size][size]int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			rate := r1.Intn(100)
			if rate <= 35 {
				board[i][j] = r1.Intn(2)
			}
		}
	}

	return board
}

func createBorder(board [size][size]int) [size][size]int {
	for i := 0; i < size; i++ {
		board[i][0] = 2
	}
	for j := 0; j < size; j++ {
		board[0][j] = 2
	}
	for i := size - 1; i > 0; i-- {
		board[i][size-1] = 2
	}
	for j := size - 1; j > 0; j-- {
		board[size-1][j] = 2
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	borderSelect := r1.Intn(4)

	if borderSelect == 0 {
		board[r1.Intn(size)][0] = 3
	} else if borderSelect == 1 {
		board[0][r1.Intn(size)] = 3
	} else if borderSelect == 2 {
		board[r1.Intn(size)][size-1] = 3
	} else if borderSelect == 3 {
		board[size-1][r1.Intn(size)] = 3
	}

	return board
}

func printBoard(board [size][size]int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				fmt.Printf(" -")
			} else if board[i][j] == 1 {
				fmt.Printf(" ▲")
			} else if board[i][j] == 2 {
				fmt.Printf(" ■")
			} else if board[i][j] == 3 {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}
}
