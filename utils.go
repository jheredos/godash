package godash

// Filter applies a predicate to each member of a slice, returning a new slice of all items
// that evaluate to true. The slice may be any type T, and predicate must be a function that
// takes one item of type T and returns a bool
func Filter[A any, S ~[]A](slc S, predicate func(A) bool) S {
	res := make([]A, 0)
	for _, item := range slc {
		if predicate(item) {
			res = append(res, item)
		}
	}
	return res
}

// Map takes a slice and applies a function to each element, returning the results in a new
// slice of the same length
func Map[A any, B any, S ~[]A](slc S, f func(A) B) []B {
	res := make([]B, len(slc))
	for i, item := range slc {
		res[i] = f(item)
	}
	return res
}

// Reduce applies a function to every item of a slice using the return value of the previous
// iteration. It takes three arguments: 1) a slice of type A, 2) a function taking two 
// arguments of an accumulator (type B) and a slice element (type A) and returning type B, 
// and 3) a starting value of type B. The final return value will be of type B
func Reduce[A any, B any, S ~[]A](slc S, f func(B, A) B, start B) B {
	acc := start
	for _, item := range slc {
		acc = f(acc, item)
	}
	return acc
}

// FindIndex returns the index of the first item that matches the predicate. If none is found,
// it returns -1 instead
func FindIndex[A any, S ~[]A](slc S, predicate func(A) bool) int {
	for i, item := range slc {
		if predicate(item) {
			return i
		}
	}
	return -1
}

// Every returns true only if predicate returns true for all elements of slc
func Every[A any, S ~[]A](slc S, predicate func(A) bool) bool {
	for _, item := range slc {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Range returns a slice of int with values from start to end (exclusive), counting by 1
func Range(start, end int) []int {
	if start <= end {
		return make([]int, 0)
	}
	rng := make([]int, end - start)
	for i:=0; i < len(rng); i++ {
		rng[i] = i + start
	}
	return rng
}
