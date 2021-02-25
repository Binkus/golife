package linq

// Copy2DimSliceByValue :
func Copy2DimSliceByValue(matrix [][]int) [][]int {
	duplicate := make([][]int, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]int, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

// All :
func All(matrix [][]int, f func(int) bool) bool {
	for _, slice := range matrix {
		for _, v := range slice {
			if !f(v) {
				return false
			}
		}
	}
	return matrix != nil && len(matrix) > 0
}

// Any :
func Any(matrix [][]int, f func(int) bool) bool {
	for _, slice := range matrix {
		for _, v := range slice {
			if f(v) {
				return true
			}
		}
	}
	return false
}

// EqualsByValue :
func EqualsByValue(firstArrays [][]int, secondaryArrays [][]int) bool {
	if firstArrays == nil || secondaryArrays == nil || len(firstArrays) != len(secondaryArrays) || len(firstArrays[0]) != len(secondaryArrays[0]) {
		return false
	}
	for i := 0; i < len(firstArrays); i++ {
		for j := 0; j < len(firstArrays[i]); j++ {
			if len(secondaryArrays[i]) != len(firstArrays[i]) || firstArrays[i][j] != secondaryArrays[i][j] {
				return false
			}
		}
	}
	return true
}
