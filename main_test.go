package main

import "testing"

func TestMatrixMedian(t *testing.T) {
	cases := []struct {
		in   [][]int
		want int
	}{
		{[][]int{{1, 3, 5}, {2, 6, 9}, {3, 6, 9}}, 5},
		{[][]int{{1, 3, 4}, {2, 5, 6}, {7, 8, 9}}, 5},
		{[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1},
		{[][]int{{5}}, 5},
		{[][]int{{1, 2, 3, 4, 5}}, 3},
	}
	for _, c := range cases {
		if got := MatrixMedian(c.in); got != c.want {
			t.Errorf("MatrixMedian(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
