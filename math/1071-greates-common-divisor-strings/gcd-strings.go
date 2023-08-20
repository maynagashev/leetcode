package main

import "fmt"

func main() {
	str1 := "ABCDEF"
	str2 := "ABC"
	fmt.Print(gcdOfStrings(str1, str2))
}

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	a := len(str1)
	b := len(str2)
	var reminder int

	// находим НОД по алгоритму Евклида
	for b != 0 {
		reminder = a % b
		a = b
		b = reminder
	}
	// в "a" находится НОД длин строк

	// проверяем повторяется ли подстрока длиной "a" в "str1"
	for i := 0; i < len(str1); i += a {
		if str1[i:i+a] != str1[:a] || str1[i:i+a] != str2[:a] {
			return ""
		}
	}

	// проверяем повторяется ли подстрока длиной "a" в "str2"
	for i := 0; i < len(str2); i += a {
		if str2[i:i+a] != str1[:a] || str2[i:i+a] != str2[:a] {
			return ""
		}
	}

	// сравниваем подстроку длиной "a" с "str1" и "str2"
	if str1[:a] == str2[:a] {
		return str1[:a]
	}

	return ""
}
