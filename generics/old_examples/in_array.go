package old_examples

import "fmt"

func InArrayFloat64(search float64, values []float64) bool {
	for _, v := range values {
		if search == v {
			return true
		}
	}
	return false
}

func InArrayInt64(search int64, values []int64) bool {
	for _, v := range values {
		if search == v {
			return true
		}
	}
	return false
}

func InArrayString(search string, values []string) bool {
	for _, v := range values {
		if search == v {
			return true
		}
	}
	return false
}

func InArray[T comparable](search T, values []T) bool {
	for _, v := range values {
		if search == v {
			return true
		}
	}
	return false
}

func main_arr() {
	valsInt := []int64{1, 2, 3, 4, 5}
	valsFloat := []float64{1, 2, 3, 4, 5}
	valsString := []string{"A", "B", "C", "D"}

	fmt.Println("Floats:", InArrayFloat64(3, valsFloat))
	fmt.Println("Int64s:", InArrayInt64(3, valsInt))
	fmt.Println("Strings:", InArrayString("C", valsString))

	fmt.Println("Floats:", InArray(3, valsInt))
	fmt.Println("Int64s:", InArray(3, valsFloat))
	fmt.Println("Strings:", InArray("C", valsString))

	fmt.Println("Bytes:", InArray([]byte("C")[0], []byte("ABCD")))
	fmt.Println("Runes:", InArray([]rune("C")[0], []rune("ABCD")))
}
