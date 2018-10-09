package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const width int = 202
const height int = 152
const alpha float64 = 1
const screenScale int = 5

func main() {
	board := genBoard()
	renderBoard(board)
	// thermalEquilibrium(board)
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
	newTemp := alpha * (((4 * (board[cellX][cellY-1] + board[cellX][cellY+1] + board[cellX-1][cellY] + board[cellX+1][cellY])) + board[cellX+1][cellY-1] + board[cellX-1][cellY+1] + board[cellX-1][cellY-1] + board[cellX+1][cellY+1]) / 20)

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
	for areEqual == true {
		for i := 1; i < height-1; i++ {
			for j := 1; j < width-1; j++ {
				if board[i][j] != auxMatrix[i][j] {
					areEqual = false
				}
			}
		}
	}
	return areEqual
}

func thermalEquilibrium(board [height][width]float64) {
	iterations := 0
	for compareMatrix(updateBoard(board), board) != true {
		fmt.Println(iterations)
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
	// --- COLOR SCALE ---
	if color == 0 {
		renderer.SetDrawColor(124, 221, 244, 255)
	} else if color > 0 && color < 11 {
		renderer.SetDrawColor(114, 243, 232, 255)
	} else if color > 10 && color < 21 {
		renderer.SetDrawColor(104, 242, 194, 255)
	} else if color > 20 && color < 31 {
		renderer.SetDrawColor(94, 241, 150, 255)
	} else if color > 31 && color < 40 {
		// renderer.SetDrawColor(85, 240, 101, 255)
		// renderer.SetDrawColor(200, 230, 191, 255)
		renderer.SetDrawColor(255, 255, 255, 255)
	} else if color > 40 && color < 51 {
		renderer.SetDrawColor(102, 239, 75, 255)
	} else if color > 50 && color < 61 {
		renderer.SetDrawColor(141, 239, 66, 255)
	} else if color > 60 && color < 71 {
		renderer.SetDrawColor(184, 238, 56, 255)
	} else if color > 70 && color < 81 {
		renderer.SetDrawColor(233, 237, 47, 255)
	} else if color > 80 && color < 91 {
		renderer.SetDrawColor(236, 187, 38, 255)
	} else if color > 90 && color < 100 {
		renderer.SetDrawColor(235, 128, 29, 255)
	} else if color == 100 {
		renderer.SetDrawColor(235, 65, 20, 255)
	}

	// --- COLOR GUESS ---
	// red := uint8(color)
	// blue := 100 - red
	// renderer.SetDrawColor(red, 20, blue, 255)

	renderer.DrawPoint(int32(y), int32(x))
}

func renderBoard(board [height][width]float64) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Thermal Simulation", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(width*screenScale), int32(height*screenScale), 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	renderer.SetScale(float32(screenScale), float32(screenScale))

	renderer.Clear()

	iterations := 0

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
		for compareMatrix(board, updateBoard(board)) != true {
			for i := 1; i < height-1; i++ {
				for j := 1; j < width-1; j++ {
					color := int(board[i][j])
					changeColor(renderer, color, i, j)
				}
			}
			renderer.Present()
			board = updateBoard(board)
			fmt.Println(iterations)
			iterations++
		}
	}
}
