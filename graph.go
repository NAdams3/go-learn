package main

import (
	"slices"
)

func GraphMatrixBreadthFirstSearch(matrix [][]int, source, needle int) []int {
	walks := []int{source}
	paths := [][]int{}
	counts := []int{}

	for n := 0; n < len(walks); n++ {
		vertex := walks[n]
		for i, conn := range matrix[vertex] {
			if conn == 0 {
				continue
			}
			if !slices.Contains(walks, i) {
				walks = append(walks, i)
				if n == 0 {
					paths = append(paths, []int{source, i})
					counts = append(counts, conn)
				} else {

					for c, path := range paths {
						if path[len(path)-1] == vertex {
							paths[c] = append(paths[c], i)
							counts[c] += conn
						}

					}
				}
			}
		}
	}

	shortestPath := []int{}
	pathCost := 0

	for i, path := range paths {
		if slices.Contains(path, needle) {
			if len(shortestPath) == 0 {
				shortestPath = path
				pathCost = counts[i]
			} else {
				if counts[i] < pathCost {
					shortestPath = path
					pathCost = counts[i]
				}
			}

		}
	}

	return shortestPath
}
