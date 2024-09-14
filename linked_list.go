//На месте программиста Бори вполне могли оказаться и вы. Поэтому вам тоже нужна практика. Самостоятельно
//
//напишите для типа *Queue метод Pop(), который будет убирать первый элемент из очереди и возвращать:
//
//строку, хранимую в этом элементе;
//логическое значение, равное true при успешной операции, и false, если метод вызван для пустой очереди.
//
//При правильной реализации метода программа должна напечатать "На золотом крыльце сидели: царь, царевич, король, королевич.".
//Не забудьте присвоить полю last значение nil, когда будет удалён последний элемент из очереди.
//

// {"На", "золотом", "крыльце", "сидели:", "царь,", "царевич,", "король,", "королевич."}

//type QueueItem struct {
//value string     // "На"
//next  *QueueItem // "золотом"
//}

//type QueueItem struct {
//value string     // "золотом"
//next  *QueueItem // крыльце
//}

//type QueueItem struct {
//value string     // "крыльце"
//next  *QueueItem // сидели
//}

//type QueueItem struct {
//value string     // "сидели"
//next  *QueueItem // nil
//}

package main

import (
	"fmt"
	"strings"
)

// Queue описывает очередь.
type Queue struct {
	first *QueueItem // указатель на первый элемент очереди
	last  *QueueItem // указатель на последний элемент очереди
}

// QueueItem описывает элемент очереди.
type QueueItem struct {
	value string     // данные
	next  *QueueItem // указатель на следующий элемент
}

// Pop удаляет первый элемент из очереди и возвращает хранимую там строку.
// вставьте здесь метод Pop() для типа *Queue
// ...

// Push добавляет в конец очереди элемент с указанной строкой.
func (queue *Queue) Push(value string) {
	item := &QueueItem{value: value}
	if queue.first == nil { // нет элементов
		// очередь пустая, поэтому добавляемый элемент
		// станет и первым и последним
		queue.first = item
		queue.last = item
		return
	}
	queue.last.next = item // текущий последний элемент должен указывать
	// на добавленный элемент
	queue.last = item // item становится последним элементом
}

func (queue *Queue) Pop() (string, bool) {
	if queue.first == nil {
		return "", false
	}

	first := *queue.first
	if queue.first.next == nil {
		queue.last = nil
		queue.first = nil
		return first.value, true
	}
	queue.first = queue.first.next

	return first.value, true
}

func linkedList() {
	list := []string{"На", "золотом", "крыльце", "сидели:", "царь,", "царевич,",
		"король,", "королевич."}
	queue := &Queue{}

	for _, v := range list {
		queue.Push(v)
	}
	list = list[:0]
	for {
		v, ok := queue.Pop()
		if !ok {
			break
		}
		list = append(list, v)
	}
	fmt.Print(strings.Join(list, " "))
}
