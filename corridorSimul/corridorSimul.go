package main

import (
	"fmt"
	"math/rand"
	"time"
)

const width int = 30
const height int = 10
const spawnChance int = 20

func main() {
	board := genBoard()
	for true {
		board = spawn(board)
		printBoard(board)
		board = moveBoard(board)
		time.Sleep(500 * time.Millisecond)
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
		for j := 0; j < width; j++ {
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
