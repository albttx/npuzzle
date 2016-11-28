package npuzzle

import (
	"io/ioutil"
	"log"
	"sort"
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
	for _, e := range raw {
		val, err := strconv.Atoi(e)
		if err != nil {
			log.Fatal("Atoi error converting line")
		}
		*greed = append(*greed, val)
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

	sortGreed := make([]int, boardSize*boardSize)
	copy(sortGreed, greed)
	sort.Ints(sortGreed)
	for i, val := range sortGreed {
		if val != i {
			log.Fatal("Duplicate value")
		}
	}
	return greed
}
