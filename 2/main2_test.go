package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}
	result := sliceExample(s)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	s := []int{1, 2, 3}
	result := addElements(s, 4)
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestCopySlice(t *testing.T) {
	s := []int{1, 2, 3}
	copied := copySlice(s)
	s[0] = 100
	if reflect.DeepEqual(s, copied) {
		t.Error("Копия не должна изменяться при изменении оригинала")
	}
}

func TestRemoveElement(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	result := removeElement(s, 2)
	expected := []int{1, 2, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}

	// Проверка на некорректный индекс
	result = removeElement(s, 10)
	if !reflect.DeepEqual(result, s) {
		t.Errorf("При неверном индексе слайс должен остаться прежним")
	}
}
