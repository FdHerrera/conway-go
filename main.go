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

func applyToEach(funcToApply func(grid [][]int, row int, col int), grid grid) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			funcToApply(grid, row, col)
		}
	}
}
