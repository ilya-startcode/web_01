package main

import (
	list "example/list/storage"
	"fmt"
)

func main() {
	l := list.NewList()

	l.Add(3)
	l.Add(4)
	l.Add(45)
	l.Add(3)
	l.Add(8)
	l.Add(19)
	l.Add(110)
	l.Add(1)
	l.Add(87)
	l.Add(54)

	fmt.Println("Длина листа:", l.Len())

	values, ok := l.GetAll()
	if ok {
		fmt.Println("Элементы листа:", values)
	} else {
		fmt.Println("Пусто")
	}

	l.RemoveByIndex(2)
	fmt.Println("Длина листа после удаления элемента по индексу:", l.Len())

	values, ok = l.GetAll()
	if ok {
		fmt.Println("Элементы в листе после удаления:", values)
	} else {
		fmt.Println("Пусто")
	}
	l.RemoveAllByValue(11)
	l.RemoveByValue(5)
	values, ok = l.GetAll()
	if ok {
		fmt.Println("Элементы в листе после удаления по значению:", values)
	} else {
		fmt.Println("Пусто")
	}

	fmt.Println(l.GetByIndex(3))
	fmt.Println(l.GetByValue(27))
	fmt.Println(l.GetAllByValue(2))

	fmt.Println("Вывод элементов листа:")
	l.Print()

	l.Clear()
	fmt.Println("Длина листа после очистки:", l.Len())
}
