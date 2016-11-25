package npuzzle

import (
	"flag"
	"math"
)

type heurFct func(int, int, int, int) int

func heuristiques(greed []int) int {
	var fct heurFct

	switch flag.Lookup("heur").Value.(flag.Getter).Get().(int) {
	case 0:
		fct = heurManhattanCase
	case 1:
		fct = heurTileOutCase
	case 2:
		fct = heurMisplacedTileCase
	default:
		fct = heurManhattanCase
	}
	return heuristiquesIterate(greed, fct)
}

func heuristiquesIterate(greed []int, f heurFct) int {
	var heur int
	var yPos, xPos, yFPos, xFPos int

	len := boardSize * boardSize
	for i := 0; i < len; i++ {
		if greed[i] != 0 {
			xPos, yPos = getXYpos(i)
			xFPos, yFPos = findXYpos(finalState, greed[i])
			heur += f(xPos, yPos, xFPos, yFPos)
		}
	}
	return heur
}

func heurManhattanCase(xPos, yPos, xFPos, yFPos int) int {
	var dy, dx float64

	dy = math.Abs(float64(yPos - yFPos))
	dx = math.Abs(float64(xPos - xFPos))
	return int(dx+dy) + heurTileOutCase(xPos, yPos, xFPos, yFPos)
}

func heurTileOutCase(xPos, yPos, xFPos, yFPos int) int {
	var heur int

	if xPos != xFPos {
		heur += 1
	}
	if yPos != yFPos {
		heur += 1
	}
	return heur
}

func heurMisplacedTileCase(xPos, yPos, xFPos, yFPos int) int {
	if xPos == xFPos && yPos == yFPos {
		return 0
	}
	return 1
}
