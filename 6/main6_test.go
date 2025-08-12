package main

import "testing"

func TestRandomNumberGenerator(t *testing.T) {
	ch := randomNumberGenerator()

	n1 := <-ch
	n2 := <-ch

	// Проверка диапозона
	if n1 < 0 || n1 >= 100 {
		t.Errorf("Число %d вне диапазона 0..99", n1)
	}
	if n2 < 0 || n2 >= 100 {
		t.Errorf("Число %d вне диапазона 0..99", n2)
	}

	// Проверка уникальности
	if n1 == n2 {
		t.Logf("Выпали одинаковые числа %d и %d — это возможно, но маловероятно", n1, n2)
	}
}
