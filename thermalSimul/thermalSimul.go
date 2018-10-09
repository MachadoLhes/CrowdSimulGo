package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const width int = 162
const height int = 122
const screenScale int = 4

func main() {
	board := genBoard()
	renderBoard(board)
}

func genBoard() [height][width]float64 {
	board := [height][width]float64{}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			board[i][j] = 0
		}
	}

	board = createBorder(board)
	board = setInitialState(board)

	return board
}

func createBorder(board [height][width]float64) [height][width]float64 {
	auxMatrix := board
	for i := 0; i < height; i++ {
		auxMatrix[i][0] = 35
		auxMatrix[i][width-1] = 35
	}
	for i := 0; i < width; i++ {
		auxMatrix[0][i] = 35
		auxMatrix[height-1][i] = 35
	}
	return auxMatrix
}

func setInitialState(board [height][width]float64) [height][width]float64 {
	auxMatrix := board
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			auxMatrix[i][j] = 35
		}
	}
	for i := 1; i < height-1; i++ {
		auxMatrix[i][1] = 0
		auxMatrix[i][width-2] = 100
	}
	return auxMatrix
}

func thermalRule(board [height][width]float64, cellX, cellY int) float64 {
	newTemp := ((4 * (board[cellX][cellY-1] + board[cellX][cellY+1] + board[cellX-1][cellY] + board[cellX+1][cellY])) + board[cellX+1][cellY-1] + board[cellX-1][cellY+1] + board[cellX-1][cellY-1] + board[cellX+1][cellY+1]) / 20

	return newTemp
}

func updateBoard(board [height][width]float64) [height][width]float64 {
	auxMatrix := board
	for i := 1; i < height-1; i++ {
		for j := 2; j < width-2; j++ {
			auxMatrix[i][j] = thermalRule(board, i, j)
		}
	}
	return auxMatrix
}

func compareMatrix(board [height][width]float64, auxMatrix [height][width]float64) bool {
	areEqual := true
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			if board[i][j] != auxMatrix[i][j] {
				areEqual = false
			}
		}
	}
	return areEqual
}

func thermalEquilibrium(board [height][width]float64) {
	iterations := 0
	for compareMatrix(updateBoard(board), board) != true {
		iterations++
		board = updateBoard(board)
	}
	fmt.Printf("Achieved thermal equilibrium in %d iterations\n", iterations)
}

func printBoard(board [height][width]float64) {
	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			fmt.Printf("%.2f ", board[i][j])
		}
		fmt.Println()
	}
}

func changeColor(renderer *sdl.Renderer, color, x, y int) {
	if color == 0 {
		renderer.SetDrawColor(82, 228, 239, 255)
	} else if color > 0 && color < 11 {
		renderer.SetDrawColor(73, 238, 201, 255)
	} else if color > 10 && color < 21 {
		renderer.SetDrawColor(65, 237, 148, 255)
	} else if color > 20 && color < 31 {
		renderer.SetDrawColor(56, 236, 90, 255)
	} else if color > 31 && color < 40 {
		renderer.SetDrawColor(255, 255, 255, 255)
	} else if color > 40 && color < 51 {
		renderer.SetDrawColor(67, 235, 48, 255)
	} else if color > 50 && color < 61 {
		renderer.SetDrawColor(117, 234, 40, 255)
	} else if color > 60 && color < 71 {
		renderer.SetDrawColor(170, 233, 31, 255)
	} else if color > 70 && color < 81 {
		renderer.SetDrawColor(228, 232, 23, 255)
	} else if color > 80 && color < 91 {
		renderer.SetDrawColor(231, 171, 15, 255)
	} else if color > 90 && color < 100 {
		renderer.SetDrawColor(230, 103, 7, 255)
	} else if color == 100 {
		renderer.SetDrawColor(229, 30, 0, 255)
	}
	renderer.DrawPoint(int32(y), int32(x))
}

func renderBoard(board [height][width]float64) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Thermal Simulation", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(width), int32(height), sdl.WINDOW_FULLSCREEN_DESKTOP)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	renderer.SetLogicalSize(int32(width), int32(height))

	renderer.Clear()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
		for compareMatrix(updateBoard(board), board) != true {
			for i := 1; i < height-1; i++ {
				for j := 1; j < width-1; j++ {
					color := int(board[i][j])
					changeColor(renderer, color, i, j)
				}
			}
			renderer.Present()
			board = updateBoard(board)
		}
	}
}
