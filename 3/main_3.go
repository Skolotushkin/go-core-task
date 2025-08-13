package main

type StringIntMap struct {
	data map[string]int
}

// NewStringIntMap — конструктор
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

// Add
// Добавляет новую пару "ключ-значение"
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// Remove
// Удаляет элемент по ключу
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Copy
// Возвращает копию карты
func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int, len(m.data))
	for k, v := range m.data {
		newMap[k] = v
	}
	return newMap
}

// Exists
// Проверяет наличие ключа
func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

// Get
// Возвращает значение и флаг успеха
func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}
