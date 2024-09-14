package old_examples

import "fmt"

func SumNumbers[K comparable, V int64 | float64](data map[K]V) V {
	var result V
	for _, num := range data {
		result += num
	}
	return result
}

func SumNumbersInterface(data map[string]interface{}) interface{} {
	var resultInt int64
	var resultFloat float64
	for _, num := range data {
		switch item := num.(type) {
		case int64:
			resultInt += item
		case float64:
			resultFloat += item
		}
	}
	if float64(resultInt) > resultFloat {
		return resultInt
	}
	return resultFloat
}

func main_maos() {
	intMap := map[string]int64{
		"key1": 12,
		"key2": 24,
	}

	floatMap := map[string]float64{
		"key1": 21.24,
		"key2": 42.12,
	}

	fmt.Printf("Result ints: %d \n", SumNumbers(intMap))
	fmt.Printf("Result floats: %f \n", SumNumbers(floatMap))

	fmt.Printf("Result ints with instantiation: %d \n", SumNumbers[string, int64](intMap))
	fmt.Printf("Result floats with instantiation: %f \n", SumNumbers[string, float64](floatMap))

	//fmt.Printf("Result ints with interfaces: %d \n", SumNumbersInterface(intMap))
	//fmt.Printf("Result floats with interfaces: %d \n", SumNumbersInterface(floatMap))
}
