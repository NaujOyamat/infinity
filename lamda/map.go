package lamda

// MapPredicate function
type MapPredicate func(interface{}) interface{}

// MapStringPredicate function
type MapStringPredicate func(string) string

// MapBytePredicate function
type MapBytePredicate func(byte) byte

// MapIntPredicate function
type MapIntPredicate func(int) int

// MapInt16Predicate function
type MapInt16Predicate func(int16) int16

// MapInt32Predicate function
type MapInt32Predicate func(int32) int32

// MapInt64Predicate function
type MapInt64Predicate func(int64) int64

// MapFloat32Predicate function
type MapFloat32Predicate func(float32) float32

// MapFloat64Predicate function
type MapFloat64Predicate func(float64) float64

// MapBoolPredicate function
type MapBoolPredicate func(bool) bool

// Map creates a new array populated with the results of calling
// a provided function on every element in the calling array
func Map(array []interface{}, pred MapPredicate) []interface{} {
	result := make([]interface{}, 0)
	for _, item := range array {
		result = append(result, pred(item))
	}
	return result
}

// MapString creates a new array populated with the results of calling
// a provided function on every element in the calling array
func MapString(array []string, pred MapStringPredicate) []string {
	result := make([]string, 0)
	for _, item := range array {
		result = append(result, pred(item))
	}
	return result
}

// MapByte creates a new array populated with the results of calling
// a provided function on every element in the calling array
func MapByte(array []byte, pred MapBytePredicate) []byte {
	result := make([]byte, 0)
	for _, item := range array {
		result = append(result, pred(item))
	}
	return result
}

// MapInt creates a new array populated with the results of calling
// a provided function on every element in the calling array
func MapInt(array []int, pred MapIntPredicate) []int {
	result := make([]int, 0)
	for _, item := range array {
		result = append(result, pred(item))
	}
	return result
}

// MapInt16 creates a new array populated with the results of calling
// a provided function on every element in the calling array
func MapInt16(array []int16, pred MapInt16Predicate) []int16 {
	result := make([]int16, 0)
	for _, item := range array {
		result = append(result, pred(item))
	}
	return result
}

// MapInt32 creates a new array populated with the results of calling
// a provided function on every element in the calling array
func MapInt32(array []int32, pred MapInt32Predicate) []int32 {
	result := make([]int32, 0)
	for _, item := range array {
		result = append(result, pred(item))
	}
	return result
}

// MapInt64 creates a new array populated with the results of calling
// a provided function on every element in the calling array
func MapInt64(array []int64, pred MapInt64Predicate) []int64 {
	result := make([]int64, 0)
	for _, item := range array {
		result = append(result, pred(item))
	}
	return result
}

// MapFloat32 creates a new array populated with the results of calling
// a provided function on every element in the calling array
func MapFloat32(array []float32, pred MapFloat32Predicate) []float32 {
	result := make([]float32, 0)
	for _, item := range array {
		result = append(result, pred(item))
	}
	return result
}

// MapFloat64 creates a new array populated with the results of calling
// a provided function on every element in the calling array
func MapFloat64(array []float64, pred MapFloat64Predicate) []float64 {
	result := make([]float64, 0)
	for _, item := range array {
		result = append(result, pred(item))
	}
	return result
}

// MapBool creates a new array populated with the results of calling
// a provided function on every element in the calling array
func MapBool(array []bool, pred MapBoolPredicate) []bool {
	result := make([]bool, 0)
	for _, item := range array {
		result = append(result, pred(item))
	}
	return result
}
