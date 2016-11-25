package npuzzle

import "reflect"

func isGreedSolved(greed []int) bool {
	return reflect.DeepEqual(finalState, greed)
}

func getXYpos(i int) (int, int) {
	xPos := (i + boardSize) % boardSize
	yPos := i / boardSize
	return xPos, yPos
}

func findXYpos(greed []int, toFind int) (int, int) {
	len := boardSize * boardSize
	for i := 0; i < len; i++ {
		if greed[i] == toFind {
			return getXYpos(i)
		}
	}
	return -1, -1
}

func SearchInts(arr []int, toFind int) int {
	for i, value := range arr {
		if value == toFind {
			return i
		}
	}
	return -1
}

func GetEmptyValue(greed []int) int {
	for i, val := range greed {
		if val == 0 {
			return i
		}
	}
	return -1
}
