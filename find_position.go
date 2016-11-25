package npuzzle

func getMin(args ...int) int {
	var ret int = 1000000

	for _, v := range args {
		if v < ret {
			ret = v
		}
	}
	return ret
}

// n: map size
// v: value search
// http://stackoverflow.com/questions/40494145/algorithm-find-number-position-in-snail-2d-array/
// Return the position on the board where the value is
func findGoalPosition(n, v int) [2]int {
	var r, span int

	if v == 0 {
		v = n * n
	}
	span = n
	for v > span {
		v -= span
		r += 1
		span -= r % 2
	}
	d := r / 4
	m := r % 4
	c := n - 1 - d
	y := []int{d, d + v, c, c - v}[m]
	x := []int{d + v - 1, c, c - v, d}[m]
	return [2]int{y, x}
}

// n: map size
// y, x: position on the board
// return the value supposed to be at Y,X
func getValueInBoard(n, y, x int) int {
	var val int

	if y == 0 {
		return x + 1
	}
	k := getMin(y-1, n-y-1, x, n-x-1)
	before := (4 * k * n) - (2 * k) - (4 * k * k)
	side := n - 2*(k+1)
	if x+k == n-1 {
		val = y - k
	} else if y+k == n-1 {
		val = (n - x - k - 1) + side + 1
	} else if x == k {
		val = (n - y - k - 1) + 2*side + 2
	} else {
		val = (x - k) + (3*side + 2)
	}
	if val+n+before == n*n {
		return 0
	}
	return val + n + before
}
