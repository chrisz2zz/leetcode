package main

import "fmt"

func main() {
	fmt.Println(findNumberIn2DArray([][]int{{-5}}, -2))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	xl, xr := 0, len(matrix)-1
    if len(matrix) == 0 {
        return false
    }

	yl, yr := 0, len(matrix[0])-1

	for xl <= xr && yl <= yr {
		xm := xl + (xr-xl)/2
		ym := yl + (yr-yl)/2

		if matrix[xm][ym] == target {
			return true
		} else if matrix[xm][ym] > target {
			xr = xm - 1
			yr = ym - 1
		} else {
			xl = xm + 1
			yl = ym + 1
		}
	}
	return false
}