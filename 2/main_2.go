package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Возвращает новый слайс с четными числами
func sliceExample(slice []int) []int {
	result := []int{}
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

// Добавляет число в конец слайса
func addElements(slice []int, num int) []int {
	return append(slice, num)
}

// Возвращает копию слайса
func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// Удаляет элемент по индексу
func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) { // проверка границ
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// 1. Генерация случайных значений
	rand.New(rand.NewSource(time.Now().UnixNano()))
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100) // случайные числа от 0 до 99
	}
	fmt.Println("1. Оригинальный слайс:", originalSlice)

	// 2. Четные числа
	evens := sliceExample(originalSlice)
	fmt.Println("2. Четные числа:", evens)

	// 3. Добавление числа
	newSlice := addElements(originalSlice, 777)
	fmt.Println("3. После добавления 777:", newSlice)

	// 4. Копия слайса
	copied := copySlice(originalSlice)
	fmt.Println("4. Копия слайса:", copied)

	// Проверка, что изменения в оригинале не влияют на копию
	originalSlice[0] = -1
	if originalSlice[0] == copied[0] {
		fmt.Println("4.1. Изменения на оригинале влияют на копию!")
	}
	fmt.Println("4.1. Проверка на независимость пройдена")

	// 5. Удаление элемента по индексу
	removed := removeElement(originalSlice, 0)
	fmt.Println("5. После удаления элемента с индексом 0:", removed)
}
