package pipe

// Prop returns a function that when supplied an map returns the indicated key of that map, if it exists.
func Prop(prop string) func(map[string]interface{}) interface{} {
	return func(obj map[string]interface{}) interface{} {
		return obj[prop]
	}
}
