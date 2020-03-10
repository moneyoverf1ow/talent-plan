package main

import "sync"

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	mergeSort(src, 0, len(src))
}

func mergeSort(a []int64, start int, end int) {
	if end-start < 2 {
		return
	}
	m := (start + end) / 2
	if end-start >= 1<<11 {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			mergeSort(a, start, m)
		}()
		mergeSort(a, m, end)
		wg.Wait()
	} else {
		// don't need use goroutine when slice is very small
		mergeSort(a, start, m)
		mergeSort(a, m, end)
	}
	merge(a, start, end)
}

func merge(src []int64, start int, end int) {
	middle := (start + end) / 2
	l, r := src[start:middle], src[middle:end]
	res := make([]int64, end-start)
	for lLen, rLen, i, j, k := len(l), len(r), 0, 0, 0; k < lLen+rLen; k += 1 {
		if i < lLen && j < rLen {
			if l[i] < r[j] {
				res[k] = l[i]
				i += 1
			} else {
				res[k] = r[j]
				j += 1
			}
		} else if i == lLen && j < rLen {
			res[k] = r[j]
			j += 1
		} else if i < lLen && j == rLen {
			res[k] = l[i]
			i += 1
		}
	}
	copy(src[start:end], res)
}
