package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	numbers := make([]int, N)
	indexes := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&numbers[i])
		indexes[i] = i
	}

	facts := genFactorials(numbers)

	for pi, max := 0, fact(N); pi < max; pi++ {
		if pi == 0 {
			printAt(numbers, indexes)
			continue
		}

		if pi%2 == 1 {
			swap(indexes, N-2, N-1)
			printAt(numbers, indexes)
			continue
		}

		for i, v := range facts {
			if pi%v != 0 {
				continue
			}

			if i == N-3 {
				// Since the order doesn't matter for the last three elements, this
				// avoids some extra computation.
				swap(indexes, i, N-1)
				break
			}

			j := i + 1 + findNextIndex(indexes[i+1:], indexes[i])
			swap(indexes, i, j)

			// Sort until the last three elements
			for x := range indexes[i+1 : N-3] {
				x += i + 1
				y := x + 1 + findLowestIndex(indexes[x+1:])
				swap(indexes, x, y)
			}

			break
		}

		printAt(numbers, indexes)
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}

	r := 1

	for ; n > 1; n-- {
		r *= n
	}

	return r
}

func findLowestIndex(s []int) int {
	var lowest, index int

	for i, v := range s {
		if i == 0 || v < lowest {
			lowest = v
			index = i
		}
	}

	return index
}

func findNextIndex(s []int, c int) int {
	var lowest, index int = 0, -1

	for i, v := range s {
		if v > c && (index == -1 || v < lowest) {
			lowest = v
			index = i
		}
	}

	return index
}

func genFactorials(s []int) []int {
	l := len(s)

	if len(s) < 3 {
		return nil
	}

	facts := make([]int, 0, l-2)

	for i := l - 1; i > 1; i-- {
		facts = append(facts, fact(i))
	}

	return facts
}

func printAt(s []int, indexes []int) {
	for i, v := range indexes {
		if i > 0 {
			fmt.Print(" ")
		}

		fmt.Print(s[v])
	}

	fmt.Println()
}

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}
