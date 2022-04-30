/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start

var res [][]string
var builder = strings.Builder{}

func solveNQueens(n int) [][]string {
	record := make([]int, 0, n)
	res = make([][]string, 0)
	reserve(n, record)
	return res
}

func reserve(n int, record []int) {
	if len(record) == n {
		tmp := make([]string, 0, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if j == record[i] {
					builder.WriteString("Q")
				} else {
					builder.WriteString(".")
				}
			}
			tmp = append(tmp, builder.String())
			builder.Reset()
		}
		res = append(res, tmp)
		return
	}

	for i := 0; i < n; i++ {
		if !in(i, record) {
			continue
		}

		record = append(record, i)
		reserve(n, record)
		record = record[:len(record)-1]
	}
}

func in(n int, record []int) bool {
	for i := 0; i < len(record); i++ {
		if n == record[i] || n == record[i]+(len(record)-i) || n == record[i]-(len(record)-i) {
			return false
		}
	}

	return true
}

// @lc code=end

