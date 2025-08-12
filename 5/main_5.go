package main

// intersection
// Возвращает true и срез пересечений, если есть хотя бы одно совпадение
func intersection(slice1, slice2 []int) (bool, []int) {
	// Создаем карту для хранения уникальных элементов первого слайса
	uniqueElements := make(map[int]bool)

	// Заполняем карту элементами первого слайса
	for _, num := range slice1 {
		uniqueElements[num] = true
	}

	// Создаем слайс для хранения пересечения
	var result []int

	// Проверяем элементы второго слайса на наличие в карте уникальных элементов
	var exists bool
	for _, num := range slice2 {
		if uniqueElements[num] {
			exists = true // Если элемент присутствует в карте, добавляем его в результат
			result = append(result, num)
			// Удаляем элемент из карты, чтобы избежать повторного добавления
			delete(uniqueElements, num)
		}
	}

	return exists, result
}
