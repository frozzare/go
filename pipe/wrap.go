package pipe

// Wrap wraps existing function and returns a pipeable function.
//
// Example:
//   loop := Wrap(slices.Map, func(item interface{}) interface{} {
//      return item.(int) + 1
//   })
//
//   p := Pipe(loop)
//   v, err := p([]int{1,2,3})
func Wrap(fn interface{}, args ...interface{}) interface{} {
	return func(first interface{}) interface{} {
		res, err := callFunc(fn, append([]interface{}{first}, args...)...)

		if err != nil {
			return err
		}

		if len(res) > 0 {
			return res[0]
		}

		return nil
	}
}
