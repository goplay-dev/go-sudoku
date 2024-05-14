package main

import (
	"fmt"
	sudoku "github.com/goplay-dev/go-sudoku"
)

func main() {
	game := &sudoku.Game{
		SqrtDimension: 3,
		Board:         nil,
	}

	game.InitGame(25)

	for _, row := range game.Board {
		fmt.Println(row)
	}
}
