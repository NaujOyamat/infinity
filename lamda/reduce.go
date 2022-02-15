package lamda

// ReducePredicate function
type ReducePredicate func(previousValue, currentValue interface{}, idx int) interface{}

// Reduce method executes a user-supplied "reducer" callback function on each
// element of the array, in order, passing in the return value from the calculation
// on the preceding element. The final result of running the reducer across all
// elements of the array is a single value.
func Reduce(array []interface{}, pred ReducePredicate, initial interface{}) interface{} {
	result := initial
	for i, item := range array {
		result = pred(result, item, i)
	}
	return result
}
