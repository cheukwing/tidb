package main

import (
	"sync"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	mergesort(src, 0, len(src))
}

func mergesort(src []int64, start, end int) {
	if end-start <= 512 {
		quicksort(src, start, end)
		return
	}

	middle := ((end - start) / 2) + start
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		mergesort(src, start, middle)
	}()
	go func() {
		defer wg.Done()
		mergesort(src, middle, end)
	}()

	wg.Wait()

	merge(src, start, end)
}

func merge(src []int64, start, end int) {
	middle := ((end - start) / 2) + start
	merged := make([]int64, end-start)

	// TODO: parallelise
	idx := 0
	l, r := start, middle
	for l < middle && r < end {
		if src[l] < src[r] {
			merged[idx] = src[l]
			l++
		} else {
			merged[idx] = src[r]
			r++
		}
		idx++
	}

	if l == middle {
		for ; r < end; r++ {
			merged[idx] = src[r]
			idx++
		}
	} else {
		for ; l < middle; l++ {
			merged[idx] = src[l]
			idx++
		}
	}

	for i, n := range merged {
		src[start+i] = n
	}
}

func quicksort(src []int64, start, end int) {
	if end-start <= 1 {
		return
	}

	pivotIdx := partition(src, start, end)
	quicksort(src, start, pivotIdx)
	quicksort(src, pivotIdx+1, end)
}

func partition(src []int64, start, end int) int {
	pivot := src[end-1]
	smalls := start - 1
	for i := start; i < end; i++ {
		if src[i] < pivot {
			smalls++
			src[smalls], src[i] = src[i], src[smalls]
		}
	}
	src[smalls+1], src[end-1] = src[end-1], src[smalls+1]
	return smalls + 1
}
