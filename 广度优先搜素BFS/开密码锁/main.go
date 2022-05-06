package main

import "fmt"

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}

func openLock(deadends []string, target string) int {
	deadM := make(map[string]bool, len(deadends))
	for _, v := range deadends {
		deadM[v] = true
	}

	q := make([]string, 0)

	q = append(q, "0000")

	step := make(map[string]int)

	for i := 0; i < len(q); i++ {
		if deadM[q[i]] {
			continue
		}

		if target == q[i] {
			return step[target]
		}

		for j := 0; j < 4; j++ {
			up := add(q[i], j)
			if !deadM[up] {
				q = append(q, up)
				step[up] = step[q[i]] + 1
			}

			down := sub(q[i], j)
			if !deadM[down] {
				q = append(q, down)
				step[down] = step[q[i]] + 1
			}
		}

		deadM[q[i]] = true

	}
	return -1
}

func add(str string, j int) string {
	bytes := []byte(str)
	if bytes[j] == '9' {
		bytes[j] = '0'
	} else {
		bytes[j] += 1
	}

	return string(bytes)
}

func sub(str string, j int) string {
	bytes := []byte(str)
	if bytes[j] == '0' {
		bytes[j] = '9'
	} else {
		bytes[j] -= 1
	}

	return string(bytes)
}
