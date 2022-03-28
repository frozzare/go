package pipe

import "github.com/frozzare/go/slices"

func All(predicate interface{}) interface{} {
	return Wrap(slices.All, predicate)
}
func Any(predicate interface{}) interface{} {
	return Wrap(slices.Any, predicate)
}

func Filter(predicate interface{}) interface{} {
	return Wrap(slices.Filter, predicate)
}

func Fold(predicate interface{}) interface{} {
	return Wrap(slices.Fold, predicate)
}

func Map(predicate interface{}) interface{} {
	return Wrap(slices.Map, predicate)
}

func Reduce(predicate interface{}) interface{} {
	return Wrap(slices.Map, predicate)
}

func Reject(predicate interface{}) interface{} {
	return Wrap(slices.Reject, predicate)
}

func Uniq(predicate interface{}) interface{} {
	return Wrap(slices.Uniq, predicate)
}
