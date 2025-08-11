package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
)

func getType(vars ...interface{}) []string {
	types := make([]string, len(vars))
	for i, v := range vars {
		types[i] = reflect.TypeOf(v).String()
	}
	return types
}

func convertToString(vars ...interface{}) string {
	result := ""
	for _, v := range vars {
		result += fmt.Sprintf("%v", v)
	}
	return result
}

func stringToRune(s string) []rune {
	return []rune(s)
}

func hashWithSalt(runes []rune, salt string) string {
	// Вставляем соль в середину среза рун
	mid := len(runes) / 2
	runesWithSalt := append([]rune{}, runes[:mid]...)
	runesWithSalt = append(runesWithSalt, []rune(salt)...)
	runesWithSalt = append(runesWithSalt, runes[mid:]...)

	// Хэшируем SHA256
	hash := sha256.New()
	hash.Write([]byte(string(runesWithSalt)))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	// 1. Создает несколько переменных различных типов данных:
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i // Тип complex64

	// 2. Определяет тип каждой переменной и выводит его на экран.
	types := getType(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("Типы переменных:", types)

	// 3. Преобразует все переменные в строковый тип и объединяет их в одну строку.
	combined := convertToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("Конвертированная в строку:", combined)

	// 4. Преобразовать эту строку в срез рун.
	runes := stringToRune(combined)
	fmt.Println("Конвертированная в руну:", runes)

	// 5. Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024"
	hash := hashWithSalt(runes, "go-2024")
	fmt.Println("Захэшированв в SHA256 с солью:", hash)
}
