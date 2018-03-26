package sliceutil

// Reject returns a new slice containing all values
// in the slice that don't satisfy the predicate function.
func Reject(input, predicate interface{}) interface{} {
	return filter(input, predicate, false)
}
