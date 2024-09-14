package main

type Item interface {
	*SomeStruct1 | *SomeStruct2
	MyFunc()
}

type SomeStruct1 struct {
	ID string
}

type SomeStruct2 struct {
	ID string
}

type SomeStruct3 struct {
	ID string
}

func (s *SomeStruct1) MyFunc() {
	println(s.ID)
}
func (s *SomeStruct2) MyFunc() {
	println(s.ID)
}

func (s *SomeStruct3) MyFunc() {
	println(s.ID)
}

func DoFunc[T Item](item T) {
	item.MyFunc()
}

func main() {
	item1 := SomeStruct1{ID: "1"}
	item2 := SomeStruct2{ID: "2"}
	//item3 := SomeStruct3{ID: "2"}

	DoFunc(&item1)
	DoFunc(&item2)
	//DoFunc(&item3)
}
