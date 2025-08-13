package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	expectedBool := true
	expectedSlice := []int{64, 3} // порядок по первому слайсу

	ok, result := intersection(a, b)

	if ok != expectedBool {
		t.Errorf("Ожидалось %v, получено %v", expectedBool, ok)
	}
	if !reflect.DeepEqual(result, expectedSlice) {
		t.Errorf("Ожидалось %v, получено %v", expectedSlice, result)
	}
}

func TestIntersectionNoCommon(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	ok, result := intersection(a, b)

	if ok {
		t.Error("Не должно быть пересечений, но возвращено true")
	}
	if len(result) != 0 {
		t.Errorf("Ожидался пустой срез, получено %v", result)
	}
}

func TestIntersectionDuplicates(t *testing.T) {
	a := []int{1, 2, 2, 3}
	b := []int{2, 3, 3}

	expectedBool := true
	expectedSlice := []int{2, 3} // без повторов

	ok, result := intersection(a, b)

	if ok != expectedBool {
		t.Errorf("Ожидалось %v, получено %v", expectedBool, ok)
	}
	if !reflect.DeepEqual(result, expectedSlice) {
		t.Errorf("Ожидалось %v, получено %v", expectedSlice, result)
	}
}

func TestIntersectionEmptySlices(t *testing.T) {
	a := []int{}
	b := []int{1, 2}

	ok, result := intersection(a, b)
	if ok {
		t.Error("Ожидалось false для пустого первого слайса")
	}
	if len(result) != 0 {
		t.Errorf("Ожидался пустой срез, получено %v", result)
	}

	ok, result = intersection([]int{1, 2}, []int{})
	if ok {
		t.Error("Ожидалось false для пустого второго слайса")
	}
	if len(result) != 0 {
		t.Errorf("Ожидался пустой срез, получено %v", result)
	}
}
