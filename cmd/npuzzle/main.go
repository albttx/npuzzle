package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ale-batt/npuzzle"
)

var (
	file   string
	size   string
	heur   int
	greedy bool
	linear bool
	step   bool
)

func init() {
	flag.StringVar(&file, "f", "", "Map file")
	flag.StringVar(&size, "size", "3", "size to create map")
	flag.BoolVar(&greedy, "greedy", false, "Use greedy search algorithm")
	flag.BoolVar(&linear, "linear", false, "Goal state linear unstead of snail")
	flag.BoolVar(&step, "step", false, "Print solution step by step")
	flag.IntVar(&heur, "heur", 0, `Heuristiques:
	0 - Manhattan (default value)
	1 - TileOut
	2 - Misplaced Tile (hamming)`)
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	var txtmaps string

	if len(file) == 0 {
		fmt.Println("No map selected, i'll generate one !")
		txtmaps = npuzzle.GenerateGreed(size)
		fmt.Println(txtmaps, "\n--- Start Solver ---")
	} else {
		txtmaps = npuzzle.ParseMap(file)
	}
	npuzzle.NewGame(txtmaps)
}
