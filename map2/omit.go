package map2

// Omit returns a partial copy of an map omitting the keys specified.
func Omit(names []string, obj interface{}) (interface{}, error) {
	return pick(names, obj, false)
}
