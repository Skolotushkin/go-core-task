package main

// difference
// Возвращает элементы, которые есть в slice1, но отсутствуют в slice2
func difference(slice1, slice2 []string) []string {
	set := make(map[string]struct{})
	for _, v := range slice2 {
		set[v] = struct{}{}
	}

	result := []string{}
	for _, v := range slice1 {
		if _, exists := set[v]; !exists {
			result = append(result, v)
		}
	}
	return result
}
