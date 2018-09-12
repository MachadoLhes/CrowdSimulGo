package main

import (
	"fmt"
	"math/rand"
	"time"
)

const width int = 100
const height int = 25
const spawnChance int = 20
const populationRate int = 20

func main() {
	board := genBoard()
	board = populate(board)
	population := populationCount(board)
	left := leftCount(board, 0)
	tp := leftCount(board, 0)
	steps := 0
	printBoard(board)
	time.Sleep(1 * time.Second)
	for true {
		board = moveBoard(board)
		steps++
		left = leftCount(board, left)
		tp = leftCount(board, tp)
		printBoard(board)
		time.Sleep(800 * time.Millisecond)
		fmt.Printf("People: %d\n", population)
		fmt.Printf("People that left: %d\n", left)
		fmt.Printf("People that left in the last minute: %d\n", tp)
		if steps%60 == 0 {
			tp = 0
		}
	}
}

func genBoard() [height][width]int {
	board := [height][width]int{}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			board[i][j] = 0
		}
	}

	board = createBorder(board)
	board = createDoor(board)

	return board
}

func createBorder(board [height][width]int) [height][width]int {
	auxMatrix := board
	for i := 0; i < width; i++ {
		auxMatrix[0][i] = 2
		auxMatrix[height-1][i] = 2
	}
	return auxMatrix
}

func spawn(board [height][width]int) [height][width]int {
	auxMatrix := board
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	doesItSpawn := r1.Intn(100)
	if doesItSpawn > spawnChance {
		spawnPoint := r1.Intn(height-2) + 1
		auxMatrix[spawnPoint][width-1] = 1
	}
	return auxMatrix
}

func populationCount(board [height][width]int) int {
	population := 0
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			if board[i][j] == 1 {
				population++
			}
		}
	}
	return population
}

func leftCount(board [height][width]int, left int) int {
	for i := 1; i < height-1; i++ {
		if board[i][1] == 1 {
			left++
		}
	}
	return left
}

func populate(board [height][width]int) [height][width]int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			rate := r1.Intn(100)
			if rate <= populationRate {
				board[i][j] = r1.Intn(2)
			}
		}
	}
	return board
}

func createDoor(board [height][width]int) [height][width]int {
	auxMatrix := board
	for i := 1; i < height-1; i++ {
		auxMatrix[i][0] = 3
	}
	return auxMatrix
}

func movePiece(board [height][width]int, pointX, pointY int) [height][width]int {
	auxMatrix := board

	if board[pointX][pointY] == 1 {

		if checkDoor(board, pointX, pointY) {
			auxMatrix[pointX][pointY] = 0
			auxMatrix[pointX][width-1] = 1
		}
		if checkNeighborhood(board, pointX, pointY) {
			auxMatrix[pointX][pointY-1] = 1
			auxMatrix[pointX][pointY] = 0
		}

	}

	return auxMatrix
}

func moveBoard(board [height][width]int) [height][width]int {
	auxMatrix := board
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if auxMatrix[i][j] == 1 {
				moveChance := r1.Intn(100)
				if moveChance > 25 {
					board = movePiece(board, i, j)
				}
			}
		}
	}
	return board
}

func checkNeighborhood(board [height][width]int, pointX, pointY int) bool {
	isAvailable := false

	if board[pointX][pointY-1] == 0 {
		isAvailable = true
	}

	return isAvailable
}

func checkDoor(board [height][width]int, pointX, pointY int) bool {
	isDoor := false

	if board[pointX][pointY-1] == 3 {
		isDoor = true
	}

	return isDoor
}

func printBoard(board [height][width]int) {
	for i := 0; i < height; i++ {
		for j := 1; j < width; j++ {
			if board[i][j] == 0 {
				fmt.Printf(" -")
			} else if board[i][j] == 1 {
				fmt.Printf(" ▲")
			} else if board[i][j] == 2 {
				fmt.Printf(" ■")
			} else if board[i][j] == 3 {
				fmt.Printf(" -")
			}
		}
		fmt.Println()
	}
}
