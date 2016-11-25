package npuzzle

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Public function

func ParseMap(maps string) string {
	f, err := ioutil.ReadFile(maps)
	if err != nil {
		log.Fatal("Error opening the file", err)
	}
	return string(f)
}

// Private function

func fillLineInBoard(line string, greed *[]int) {
	raw := strings.Fields(line)
	if len(raw) != boardSize {
		log.Fatal("MAP ERROR: Different size")
	}
	for i, e := range raw {
		val, err := strconv.Atoi(e)
		if err != nil {
			log.Fatal("Atoi error converting line")
		}
		if i >= boardSize {
			log.Fatal("Line too big for boardSize")
		}
		*greed = append(*greed, val)
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

func fillBoard(file string) []int {
	var greed []int
	var err error

	y := 0
	f := strings.Split(file, "\n")
	for _, line := range f {
		if strings.HasPrefix(line, "#") {
		} else if boardSize == 0 {
			boardSize, err = strconv.Atoi(line)
			if err != nil {
				log.Fatal("Atoi error", err)
			} else if boardSize <= 1 {
				log.Fatal("Map too small")
			}
		} else {
			if y < boardSize {
				fillLineInBoard(line, &greed)
			}
			y++
		}
	}
	return greed
}
