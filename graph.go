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

func DijkstraList(list [][]WeightedEdge, source, sink int) []int {
	weights := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		if i == source {
			weights[i] = 0
		} else {
			weights[i] = 10000
		}
	}

	seen := make([]bool, len(list))
	prev := make([]int, len(list))

	for hasUnvisited(seen, weights) {
		current := getLowestUnvisited(seen, weights)
		seen[current] = true
		edges := list[current]
		for _, edge := range edges {
			if seen[edge.to] {
				continue
			}

			newWeight := weights[current] + edge.weight
			if weights[edge.to] > newWeight {
				weights[edge.to] = newWeight
				prev[edge.to] = current
			}

		}

	}

	out := []int{sink}
	for i := 0; i < len(prev); i++ {
		vertex := out[i]
		out = append(out, prev[vertex])
		if prev[vertex] == source {
			break
		}
	}

	slices.Reverse(out)
	return out

}

func hasUnvisited(seen []bool, weights []int) bool {

	found := false

	for i, isSeen := range seen {
		if !isSeen {
			if weights[i] < 10000 {
				found = true
				break
			}
		}
	}

	return found

}

func getLowestUnvisited(seen []bool, weights []int) int {

	index := -1
	minWeight := 10000

	for i, isSeen := range seen {
		if !isSeen {
			if weights[i] < minWeight {
				minWeight = weights[i]
				index = i
			}
		}
	}

	return index

}
