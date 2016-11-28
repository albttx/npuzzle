package npuzzle

import "flag"

//      2
//  1-> | <-3
//      4

func swapNode(greed []int, i1, i2 int) bool {
	tmp := greed[i1]
	greed[i1] = greed[i2]
	greed[i2] = tmp
	return true
}

func getStates(greed Board, level int, statesChan chan<- Board) {
	for i := 1; i <= 4; i++ {
		go ChangeState(greed, i, level, statesChan)
	}
}

func ChangeState(greed Board, nextPos int, level int, statesChan chan<- Board) {

	cpy := make([]int, boardSize*boardSize)
	copy(cpy, greed.list)

	i := GetEmptyValue(cpy)
	xPos, yPos := getXYpos(i)

	ret := false
	switch {
	case nextPos == 1 && xPos-1 >= 0:
		ret = swapNode(cpy, i, i-1)
	case nextPos == 3 && xPos+1 < boardSize:
		ret = swapNode(cpy, i, i+1)
	case nextPos == 2 && yPos+1 < boardSize:
		ret = swapNode(cpy, i, i+boardSize)
	case nextPos == 4 && yPos-1 >= 0:
		ret = swapNode(cpy, i, i-boardSize)
	}
	if ret {
		heur := heuristiques(cpy)
		statesChan <- Board{cpy, heur, level, heur + level, &greed}
	} else {
		statesChan <- Board{}
	}
}

func setFinalState() {
	finalState = make([]int, boardSize*boardSize)
	i := 0
	if flag.Lookup("linear").Value.(flag.Getter).Get().(bool) {
		for i = 0; i < boardSize*boardSize; i++ {
			finalState[i] = i + 1
		}
		finalState[i-1] = 0
		return
	}
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			finalState[i] = getValueInBoard(boardSize, y, x)
			i++
		}
	}
}
