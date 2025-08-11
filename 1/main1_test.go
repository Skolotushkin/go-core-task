package main

import (
	"testing"
)

func TestGetType(t *testing.T) {
	num := 1
	str := "test"
	types := getType(num, str)

	if types[0] != "int" {
		t.Errorf("Ожидался int, получено %s", types[0])
	}
	if types[1] != "string" {
		t.Errorf("Ожидался string, получено %s", types[1])
	}
}

func TestConvertToString(t *testing.T) {
	result := convertToString(1, "test", true)
	expected := "1testtrue"
	if result != expected {
		t.Errorf("Ожидалось %s, получено %s", expected, result)
	}
}

func TestStringToRune(t *testing.T) {
	runes := stringToRune("go")
	expected := []rune{'g', 'o'}
	for i, r := range runes {
		if r != expected[i] {
			t.Errorf("Ожидалось %v, получено %v", expected[i], r)
		}
	}
}

func TestHashWithSalt(t *testing.T) {
	runes := []rune("golang")
	hash := hashWithSalt(runes, "go-2024")

	if hash != string("95dd2529f334a089217182009af3349771f4535027071e17a0a36475e845fa5d") {
		t.Errorf("Хэш не совпал %s", hash)
	}
}
