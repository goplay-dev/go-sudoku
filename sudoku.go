package go_sudoku

import (
	"math/rand"
)

type Game struct {
	SqrtDimension int32
	Board         [][]int32
}

func (g *Game) InitGame(maxInitNum int) {
	g.setInitBoards()
	g.setInitNum(maxInitNum)
}

func (g *Game) setInitBoards() {
	dimension := g.SqrtDimension * g.SqrtDimension
	g.Board = make([][]int32, dimension)

	for row := int32(0); row < dimension; row++ {
		g.Board[row] = make([]int32, dimension)

		for col := int32(0); col < dimension; col++ {
			g.Board[row][col] = 0
		}
	}
}

func (g *Game) setInitNum(maxNum int) {
	if maxNum <= 0 {
		return
	}

	num := g.randomNumber() + 1
	row := g.randomNumber()
	col := g.randomNumber()

	if g.Board[row][col] <= 0 {
		if g.validateRequest(num, row, col) {
			g.Board[row][col] = num
			g.setInitNum(maxNum - 1)
		} else {
			g.setInitNum(maxNum)
		}
	} else {
		g.setInitNum(maxNum)
	}
}

func (g *Game) RequestNum(num, row, col int32) bool {
	if g.Board[row][col] > 0 {
		return false
	}

	if g.validateRequest(num, row, col) {
		g.Board[row][col] = num
	}

	return true
}

func (g *Game) GenerateAnswer(row, col int32) bool {
	dimension := g.SqrtDimension * g.SqrtDimension

	if row == dimension-1 && col == dimension {
		return true
	}

	if col == dimension {
		row++
		col = 0
	}

	if g.Board[row][col] > 0 {
		return g.GenerateAnswer(row, col+1)
	}

	for num := int32(0); num < dimension+1; num++ {
		if g.validateRequest(num, row, col) {
			g.Board[row][col] = num

			if g.GenerateAnswer(row, col+1) {
				return true
			}
		}

		g.Board[row][col] = 0
	}

	return false
}

func (g *Game) validateColumn(num, row int32) bool {
	dimension := g.SqrtDimension * g.SqrtDimension
	for c := int32(0); c < dimension; c++ {
		if g.Board[row][c] == num {
			return false
		}
	}

	return true
}

func (g *Game) validateRow(num, col int32) bool {
	dimension := g.SqrtDimension * g.SqrtDimension
	for r := int32(0); r < dimension; r++ {
		if g.Board[r][col] == num {
			return false
		}
	}

	return true
}

func (g *Game) validateArea(num, row, col int32) bool {
	startRow := row - row%g.SqrtDimension
	startCol := col - col%g.SqrtDimension

	for r := int32(0); r < g.SqrtDimension; r++ {
		for c := int32(0); c < g.SqrtDimension; c++ {
			if g.Board[r+startRow][c+startCol] == num {
				return false
			}
		}
	}

	return true
}

func (g *Game) validateRequest(num int32, row int32, col int32) bool {
	return g.validateColumn(num, row) &&
		g.validateRow(num, col) &&
		g.validateArea(num, row, col)
}

func (g *Game) randomNumber() (num int32) {
	minNum := 0
	maxNum := int(g.SqrtDimension * g.SqrtDimension)

	num = int32(rand.Intn((maxNum)-minNum) + minNum)
	return
}
