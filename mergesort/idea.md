# MergeSort

We apply a different strategy depending on the length of the slice the function is given.

## <= 512

We apply quicksort sequentially, as this will be quicker or comparable to starting a new goroutine to apply mergesort.

## > 512

We apply mergesort by starting two goroutines, which may defer to quicksort when size goes under again.
Using their results, we perform a parallel merge.
