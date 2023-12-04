package main

import (
	"log"
	"math/rand"
)

type grid [][]int

var (
	width  = 1980
	height = 1200
)

func main() {
	grid := initGrid()
	log.Print(grid)
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
