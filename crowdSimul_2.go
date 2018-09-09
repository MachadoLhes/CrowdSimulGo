package main

import "fmt"

const size int = 15

func main() {
	board := genBoard()
	board[5][10] = 1
	board[2][13] = 1
	board[9][2] = 1
	board[0][3] = 3
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
	board = moveBoard(board, 0, 3)
	printBoard(board)
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

func createBorder(board [size][size]int) [size][size]int {
	for i := 0; i < size; i++ {
		board[i][0] = 2
		board[i][size-1] = 2
	}
	for j := 0; j < size; j++ {
		board[0][j] = 2
		board[size-1][j] = 2
	}
	return board
}

func movePiece(board [size][size]int, pointX, pointY, doorX, doorY int) [size][size]int {
	auxMatrix := board

	if board[pointX][pointY] == 1 {
		dir := direction(pointX, pointY, doorX, doorY)
		if dir == "left" {
			if checkDoor(board, pointX, pointY, dir) {
				auxMatrix[pointX][pointY] = 0
			}
			if checkNeighborhood(board, pointX, pointY, dir) {
				auxMatrix[pointX][pointY-1] = 1
				auxMatrix[pointX][pointY] = 0
			}
		} else if dir == "right" {
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
		} else if dir == "down" {
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
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 1 {
				board = movePiece(board, i, j, doorX, doorY)
			}
		}
	}
	return board
}

func direction(pointX, pointY, doorX, doorY int) string {
	direction := ""
	if pointX > doorX {
		direction = "up"
	}
	if pointX < doorX {
		direction = "down"
	}
	if pointY > doorY {
		direction = "left"
	}
	if pointY < doorY {
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
	} else if direction == "down" {
		if board[pointX+1][pointY] == 3 {
			isDoor = true
		}
	}
	if direction == "left" {
		if board[pointX][pointY-1] == 3 {
			isDoor = true
		}
	} else if direction == "right" {
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
