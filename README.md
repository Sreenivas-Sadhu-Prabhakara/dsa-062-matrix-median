# Median of a Row-Sorted Matrix

> Striver SDE Sheet 062 · Day 11 Binary Search — binary search on the value, not a position; each probe counts how many elements don't exceed it. Go, stdlib only, tested.

## Problem

A matrix has each row sorted (columns unordered) and an odd number of elements. Find the median without flattening and sorting — the O(RC log RC) baseline this problem wants beaten.

## Approach

The median is the smallest value v such that at least half-plus-one of all elements are ≤ v. That predicate is monotone in v, so binary search the *value range* (global min to global max): each probe asks every row "how many of yours are ≤ mid?", a per-row upper-bound binary search (`sort.SearchInts(row, mid+1)`). Searching the value space rather than positions is the mental shift this problem teaches — the matrix is never rearranged, and the answer converges to an element that actually exists because the count only jumps at real values.

**Complexity:** O(R · log C · log(maxv−minv)) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
