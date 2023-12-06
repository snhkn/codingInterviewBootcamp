package main

func maxChar(str string) (string, int) {
	charMap := make(map[string]int)
	for _, char := range str {
		charMap[string(char)]++
	}
	var max_key string
	var max_value int

	for key, value := range charMap {
		if value > max_value {
			max_key = key
			max_value = value
		}
	}
	return max_key, max_value
}
