package npuzzle

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/fatih/color"
)

func Print(greed []int) {
	len := len(greed)

	for i := 0; i < len; i++ {
		if i%boardSize == 0 {
			fmt.Printf("\n")
		}
		if greed[i] == 0 {
			fmt.Printf("%3s", " ")
		} else if greed[i] == finalState[i] {
			color.Set(color.FgGreen)
			fmt.Printf("%3d", greed[i])
		} else {
			color.Set(color.FgRed)
			fmt.Printf("%3d", greed[i])
		}
		color.Unset()
	}
	fmt.Printf("\n")
	color.Unset()
}

func isGreedSolvable(greed []int) bool {
	ret := false
	inversion := 0
	len := len(greed)

	for i, val := range greed {
		for j := i + 1; j < len; j++ {
			fs := SearchInts(finalState, greed[j])
			if val != 0 && finalState[fs] != 0 && val > finalState[fs] {
				inversion++
			}
		}
	}
	if boardSize%2 == 0 {
		_, y := getXYpos(GetEmptyValue(greed))
		if y%2 != inversion%2 {
			ret = true
		}
	} else if boardSize%2 == 1 {
		ret = inversion%2 == 1
	}
	if boardSize > 5 {
		ret = ret == false
	}
	return ret
}

func GenerateGreed(size string) string {
	r, err := exec.Command("python", "assets/npuzzle-gen.py", "--solvable", size).Output()
	if err != nil {
		log.Fatal("Error Generating puzzle\n", err)
	}
	return string(r)
}
