package npuzzle

import (
	"container/heap"
	"flag"
	"fmt"
	"reflect"
)

type Board struct {
	list     []int
	heur     int
	level    int
	sum      int
	ancestor *Board
}

type HeapQueue []Board

func (hq HeapQueue) Len() int      { return len(hq) }
func (hq HeapQueue) Swap(i, j int) { hq[i], hq[j] = hq[j], hq[i] }

func (hq HeapQueue) Less(i, j int) bool {
	if flag.Lookup("greedy").Value.(flag.Getter).Get().(bool) {
		return hq[i].heur < hq[j].heur
	}
	return hq[i].sum < hq[j].sum
}

func (hq *HeapQueue) Push(x interface{}) {
	*hq = append(*hq, x.(Board))
}

func (hq *HeapQueue) Pop() interface{} {
	old := *hq
	n := len(old)
	x := old[n-1]
	*hq = old[0 : n-1]
	return x
}

// Methods

func (hq *HeapQueue) Print() {
	for _, e := range *hq {
		fmt.Println(e.heur)
	}
}

func (hq *HeapQueue) Find(b Board) (int, *Board) {
	for i, e := range *hq {
		if reflect.DeepEqual(e.list, b.list) {
			return i, &e
		}
	}
	return -1, nil
}

func (hq *HeapQueue) Include(b Board) bool {
	_, ret := hq.Find(b)
	return ret != nil
}

func (hq *HeapQueue) Remove(b Board) {
	i, _ := hq.Find(b)
	if i != -1 {
		heap.Remove(hq, i)
	}
}
