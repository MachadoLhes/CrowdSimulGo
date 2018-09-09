package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"
)

const size int = 15
const populationRate int = 5

func main() {
	board := genBoard()
	board = createDoor(board)
	board = populate(board)
	doorX := findDoor(board, "x")
	doorY := findDoor(board, "y")
	steps := 0
	printBoard(board)
	for emptyRoom(board) == false {
		clearScreen()
		steps++
		time.Sleep(1 * time.Second)
		board = moveBoard(board, doorX, doorY)
		printBoard(board)
	}
	fmt.Printf("Room successfuly evacuated in %d seconds\n", steps)
}

func emptyRoom(board [size][size]int) bool {
	isEmpty := true
	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			if board[i][j] != 0 {
				isEmpty = false
			}
		}
	}
	return isEmpty
}

func genBoard() [size][size]int {
	board := [size][size]int{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			board[i][j] = 0
		}
	}

	board = createBorder(board)

	return board
}

func populate(board [size][size]int) [size][size]int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			rate := r1.Intn(100)
			if rate <= populationRate {
				board[i][j] = r1.Intn(2)
			}
		}
	}
	return board
}

func createBorder(board [size][size]int) [size][size]int {
	auxMatrix := board
	for i := 0; i < size; i++ {
		auxMatrix[i][0] = 2
		auxMatrix[i][size-1] = 2
	}
	for j := 0; j < size; j++ {
		auxMatrix[0][j] = 2
		auxMatrix[size-1][j] = 2
	}
	return auxMatrix
}

func createDoor(board [size][size]int) [size][size]int {
	auxMatrix := board
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	borderSelect := r1.Intn(4)
	if borderSelect == 0 {
		auxMatrix[r1.Intn(size-1)+1][0] = 3
	} else if borderSelect == 1 {
		auxMatrix[0][r1.Intn(size-1)+1] = 3
	} else if borderSelect == 2 {
		auxMatrix[r1.Intn(size-1)+1][size-1] = 3
	} else if borderSelect == 3 {
		auxMatrix[size-1][r1.Intn(size-1)+1] = 3
	}
	return auxMatrix
}

func findDoor(board [size][size]int, coord string) int {
	coordPoint := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 3 {
				if coord == "x" {
					coordPoint = i
				} else if coord == "y" {
					coordPoint = j
				}
			}
		}
	}
	return coordPoint
}

func movePiece(board [size][size]int, pointX, pointY, doorX, doorY int) [size][size]int {
	auxMatrix := board

	if board[pointX][pointY] == 1 {
		dir := direction(board, pointX, pointY, doorX, doorY)
		if dir == "left" {
			if checkDoor(board, pointX, pointY, dir) {
				auxMatrix[pointX][pointY] = 0
			}
			if checkNeighborhood(board, pointX, pointY, dir) {
				auxMatrix[pointX][pointY-1] = 1
				auxMatrix[pointX][pointY] = 0
			}
		}
		if dir == "right" {
			if checkDoor(board, pointX, pointY, dir) {
				auxMatrix[pointX][pointY] = 0
			}
			if checkNeighborhood(board, pointX, pointY, dir) {
				auxMatrix[pointX][pointY+1] = 1
				auxMatrix[pointX][pointY] = 0
			}
		}
		if dir == "up" {
			if checkDoor(board, pointX, pointY, dir) {
				auxMatrix[pointX][pointY] = 0
			}
			if checkNeighborhood(board, pointX, pointY, dir) {
				auxMatrix[pointX-1][pointY] = 1
				auxMatrix[pointX][pointY] = 0
			}
		}
		if dir == "down" {
			if checkDoor(board, pointX, pointY, dir) {
				auxMatrix[pointX][pointY] = 0
			}
			if checkNeighborhood(board, pointX, pointY, dir) {
				auxMatrix[pointX+1][pointY] = 1
				auxMatrix[pointX][pointY] = 0
			}
		}
	}

	return auxMatrix
}

func moveBoard(board [size][size]int, doorX, doorY int) [size][size]int {
	moved := false
	if moved == false {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if board[i][j] == 1 {
					board = movePiece(board, i, j, doorX, doorY)
					moved = true
				}
			}
		}
	}
	return board
}

func direction(board [size][size]int, pointX, pointY, doorX, doorY int) string {
	direction := ""
	if pointX > doorX && board[doorX][pointY] != 2 {
		direction = "up"
	}
	if pointX < doorX && board[doorX][pointY] != 2 {
		direction = "down"
	}
	if pointY > doorY && board[pointX][doorY] != 2 {
		direction = "left"
	}
	if pointY < doorY && board[pointX][doorY] != 2 {
		direction = "right"
	}
	return direction
}

func checkNeighborhood(board [size][size]int, pointX, pointY int, direction string) bool {
	isAvailable := false
	if direction == "up" {
		if board[pointX-1][pointY] == 0 {
			isAvailable = true
		}
	} else if direction == "down" {
		if board[pointX+1][pointY] == 0 {
			isAvailable = true
		}
	}
	if direction == "left" {
		if board[pointX][pointY-1] == 0 {
			isAvailable = true
		}
	} else if direction == "right" {
		if board[pointX][pointY+1] == 0 {
			isAvailable = true
		}
	}
	return isAvailable
}

func checkDoor(board [size][size]int, pointX, pointY int, direction string) bool {
	isDoor := false
	if direction == "up" {
		if board[pointX-1][pointY] == 3 {
			isDoor = true
		}
	}
	if direction == "down" {
		if board[pointX+1][pointY] == 3 {
			isDoor = true
		}
	}
	if direction == "left" {
		if board[pointX][pointY-1] == 3 {
			isDoor = true
		}
	}
	if direction == "right" {
		if board[pointX][pointY+1] == 3 {
			isDoor = true
		}
	}
	return isDoor
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

func clearScreen() {
	// os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	exec.Command("clear")
}
