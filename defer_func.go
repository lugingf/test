package main

func checkNamed(item int) (result int) {
	defer func() {
		result = item * 2
	}()

	result = item + 10

	return result
}
