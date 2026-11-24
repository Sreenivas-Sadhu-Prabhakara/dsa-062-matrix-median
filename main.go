package main

import (
	"fmt"
	"sort"
)

// MatrixMedian returns the median of a matrix whose rows are individually
// sorted (total element count odd). Binary search on the value range; each
// probe counts elements ≤ mid with a per-row binary search.
func MatrixMedian(m [][]int) int {
	lo, hi := m[0][0], m[0][len(m[0])-1]
	total := 0
	for _, row := range m {
		if row[0] < lo {
			lo = row[0]
		}
		if row[len(row)-1] > hi {
			hi = row[len(row)-1]
		}
		total += len(row)
	}
	need := total/2 + 1 // the median is the need-th smallest
	for lo < hi {
		mid := lo + (hi-lo)/2
		count := 0
		for _, row := range m {
			count += sort.SearchInts(row, mid+1) // elements ≤ mid
		}
		if count < need {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	fmt.Println(MatrixMedian([][]int{{1, 3, 5}, {2, 6, 9}, {3, 6, 9}}))
}
