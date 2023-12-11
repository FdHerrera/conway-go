package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

var (
	gridState grid
	height    int = 600
	width     int = 800
)

type grid [][]int

type Game struct{}

func (g Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	applyToEach(func(grid [][]int, row, col int) {
		if grid[row][col] == 1 {
			screen.Set(row, col, color.Black)
		}
	}, gridState)
	gridState = calcNextGen(gridState)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	gridState = initGrid()
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Conway's Game of Life by FdHerrera")
	game := Game{}
	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatalf("Fail while running game: [%d]", err)
	}
}

func initGrid() grid {
	grid := make([][]int, width)
	for i := 0; i < width; i++ {
		col := make([]int, height)
		grid[i] = col
	}

	applyToEach(
		func(grid [][]int, row, col int) {
			grid[row][col] = rand.Intn(2)
		},
		grid,
	)
	return grid
}

func calcNextGen(grid grid) grid {
	gridCopy := make([][]int, len(grid))
	for i, col := range grid {
		colCopy := make([]int, len(col))
		copy(colCopy, col)
		gridCopy[i] = colCopy
	}
	applyToEach(
		func(grid [][]int, row, col int) {
			neighboursCount := 0
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					if i == 0 && j == 0 {
						continue
					}
					if row+i < 0 || col+j < 0 {
						continue
					}
					if row+i == len(grid) || col+j == len(grid[0]) {
						continue
					}
					neighboursCount += grid[row+i][col+j]
				}
			}
			if neighboursCount < 2 {
				gridCopy[row][col] = 0
			}
			if neighboursCount > 3 {
				gridCopy[row][col] = 0
			}
			if grid[row][col] == 0 && neighboursCount == 3 {
				gridCopy[row][col] = 1
			}
		},
		grid,
	)
	return gridCopy
}

func applyToEach(funcToApply func(grid [][]int, row int, col int), grid grid) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			funcToApply(grid, row, col)
		}
	}
}
