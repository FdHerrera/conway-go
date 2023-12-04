package main

import (
	"testing"
)

func Test_initGrid(t *testing.T) {
	grid := initGrid()

	if len(grid) != 1980 {
		t.Error("Should have 1980 default rows")
	}

	for _, col := range grid {
		if len(col) != 1200 {
			t.Error("Each col should default 1200 cols")
		}
	}

	width = 10
	height = 10

	grid = initGrid()

	if len(grid) != 10 {
		t.Error("Should have 10 cols")
	}

	colsWithTenRows := false

	for _, col := range grid {
		if len(col) != 10 {
			colsWithTenRows = true
		}
	}

	if colsWithTenRows {
		t.Error("Each col should have 10 rows")
	}
}
