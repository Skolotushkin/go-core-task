package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	expected := []string{"apple", "cherry", "43", "lead", "gno1"}
	result := difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestDifferenceAll(t *testing.T) {
	slice1 := []string{"a", "b"}
	slice2 := []string{"x", "y"}

	expected := []string{"a", "b"}
	result := difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestNoDifference(t *testing.T) {
	slice1 := []string{"a", "b"}
	slice2 := []string{"a", "b", "c"}

	expected := []string{}
	result := difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось пустой слайс, получено %v", result)
	}
}

func TestDifferenceEmptySlice1(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{"a", "b"}

	expected := []string{}
	result := difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось пустой слайс, получено %v", result)
	}
}

func TestDifferenceEmptySlice2(t *testing.T) {
	slice1 := []string{"a", "b"}
	slice2 := []string{}

	expected := []string{"a", "b"}
	result := difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}
