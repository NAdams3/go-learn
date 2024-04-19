package main

import (
	"fmt"
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
						if n >= len(path) {
							continue
						}
						if path[n] == vertex {
							path = append(path, i)
							counts[c] += conn
						}

					}
				}
			}
		}
	}

	fmt.Printf("paths: %v, counts: %v \n", paths, counts)

	return []int{}
}
