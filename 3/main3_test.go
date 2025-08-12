package main

import (
	"reflect"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	value, ok := m.Get("one")
	if !ok {
		t.Error("Ключ 'one' должен существовать")
	}
	if value != 1 {
		t.Errorf("Ожидалось 1, получено %d", value)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key", 42)
	m.Remove("key")
	_, ok := m.Get("key")
	if ok {
		t.Error("Ключ 'key' не должен существовать после удаления")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	copied := m.Copy()
	if !reflect.DeepEqual(m.data, copied) {
		t.Errorf("Копия должна совпадать с оригиналом")
	}

	// Проверка на независимость копии
	copied["a"] = 100
	if m.data["a"] == 100 {
		t.Errorf("Failure: копия не должна зависить от оригинала")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("key", 123)
	if !m.Exists("key") {
		t.Error("Ключ 'key' должен существовать")
	}
	if m.Exists("absent") {
		t.Error("Ключ 'absent' не должен существовать")
	}
}

func TestGetNonExistent(t *testing.T) {
	m := NewStringIntMap()
	_, ok := m.Get("missing")
	if ok {
		t.Error("Ключ 'missing' не должен существовать")
	}
}
