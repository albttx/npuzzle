package npuzzle

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"os"
	"sort"
)

func Resolve(greed []int) (*Board, int, int) {
	return aStarSearch(greed)
}

func aStarSearch(greed []int) (*Board, int, int) {
	var (
		openList  HeapQueue
		closeList HeapQueue
		curr      Board
		level     int
		iteration int
	)

	statesChan := make(chan Board)
	// 1. Get the first state, which is your start state.
	curr = Board{greed, heuristiques(greed), level, 0, nil}
	curr.sum = curr.heur + curr.level
	heap.Push(&openList, curr)
	for openList.Len() != 0 {
		// 2. Get the state with smallest heuristique
		sort.Sort(&openList)
		curr = heap.Pop(&openList).(Board)
		level = curr.level + 1

		if isGreedSolved(curr.list) {
			break
		}

		//statesChan := make(chan Board)
		// 3. Get all the possible states in which puzzle can be.
		go getStates(curr, level, statesChan)

		// 4. Add these new states in open state list
		for i := 1; i <= 4; i++ {
			state := <-statesChan
			if state.ancestor != nil && closeList.Include(state) == false {
				heap.Push(&openList, state)
			}
		}

		//// 5. Add processed sate in the closed list
		heap.Push(&closeList, curr)
		iteration++
	}
	return &curr, iteration, len(openList) + iteration
}

func backtrackSolution(st *Board) int {
	step := false
	if flag.Lookup("step").Value.(flag.Getter).Get().(bool) {
		step = true
	}
	i := 0
	if st.ancestor != nil {
		i += backtrackSolution(st.ancestor) + 1
	}
	if step {
		fmt.Print("Step ", i, ": ")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	Print(st.list)
	return i
}
