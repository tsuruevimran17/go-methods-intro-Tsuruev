package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (u Student) Info() string {
	return fmt.Sprintf("Студент: %s Возраст: %d ", u.Name, u.Age)
}

// TODO: добавь метод Info() со значением-получателем (Student),
// который вернёт строку вида: "Студент: Имя (Возраст: N)"
// Подсказка: fmt.Sprintf поможет собрать строку.

func main() {
	s := Student{Name: "Амина", Age: 18}
	// Ожидается печать строки с именем и возрастом
	fmt.Println(s.Info())
}
