package main

import (
	"reflect"
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

func Test_calcNextGen(t *testing.T) {
	cases := []struct {
		initial  [][]int
		expected [][]int
	}{
		{
			[][]int{
				{0, 1, 0},
				{1, 0, 0},
				{0, 1, 0},
			}, [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{0, 0, 0},
			},
		},
		{
			[][]int{
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
			[][]int{
				{0, 0, 0},
				{1, 1, 1},
				{0, 0, 0},
			},
		},
		{
			[][]int{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
			[][]int{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
		},
	}

	for i, e := range cases {
		nextGen := calcNextGen(e.initial)
		if !areEquals(nextGen, e.expected) {
			t.Errorf("Case [%d] failed, not equals\n %d \n %d", i, nextGen, e.expected)
		}
	}
}

func areEquals(actual, expected [][]int) bool {
	if len(actual) != len(expected) {
		return false
	}
	for i := 0; i < len(actual); i++ {
		if !reflect.DeepEqual(actual[i], expected[i]) {
			return false
		}
	}
	return true
}
