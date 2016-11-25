package npuzzle

import (
	"flag"
	"fmt"
)

// Structure definition

var (
	finalState []int
	boardSize  int
)

func NewGame(txtmap string) {
	var greed []int

	greed = fillBoard(txtmap)
	setFinalState()
	if isGreedSolvable(greed) == false && !flag.Lookup("linear").Value.(flag.Getter).Get().(bool) {
		fmt.Println("The greed is unsolvable")
		return
	}
	board, timeComplexity, sizeComplexity := Resolve(greed)
	nbState := backtrackSolution(board)

	fmt.Println("Solved in", nbState)
	fmt.Println("Complexity in time", timeComplexity)
	fmt.Println("Complexity in size", sizeComplexity)
}
