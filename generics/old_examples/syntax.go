package old_examples

import "fmt"

// Мы можем задать интерфейс с перечислением возможных типов
type SomeInterface interface {
}

// Можем указать возможные типы прям в сигнатуре
func FuncGen[T int | int8 | int32 | SomeInterface](argName T) {
	fmt.Println(argName)
}

func main_c() {
	ch := make(chan SomeInterface)

	// Вызов функции с аргументами
	FuncGen(1)
	FuncGen([]string{"abc", "def"})
	// Можем подсказать компилятору аргументы какого типа мы передаем
	// Providing the type argument to GMin, in this case int, is called instantiation.
	// Instantiation happens in two steps. First, the compiler substitutes all type arguments
	// for their respective type parameters throughout the generic function or type. Second,
	// the compiler verifies that each type argument satisfies the respective constraint.
	// We’ll get to what that means shortly, but if that second step fails, instantiation fails and the program is invalid.
	FuncGen[string]("some string")
}
