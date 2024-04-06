package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Если ответ существует, верните список из двух элементов
// Если нет - то верните пустой список
// Нужно вывести два числа —– очки на двух фишках, в сумме дающие k.
// Если таких пар несколько, то можно вывести любую из них.
func twoSum(array []int, targetSum int) []int {
	m := make(map[int]int)
	var compliment int
	// проходим по массиву и сохраняем в хэш комплиентарные значения которые уже встречались, так же проверяем в мапе наличие комплимента для текущего числа
	for i := 0; i < len(array); i++ {
		compliment = targetSum - array[i] // 2 - (-1) = 3
		// Если ранее уже встречалось нужное число, отдаем текущее и найденное
		if j, ok := m[compliment]; ok {
			return []int{array[j], array[i]}
		}
		// Иначе просто сохраняем в хэш таблицу индекс текущего числа если кому то оно вдруг понадобится, мы сможем его вспомнить
		m[array[i]] = i
	}
	//fmt.Printf("map: %#v\n", m)
	return []int{}
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	array := readArray(scanner)
	targetSum := readInt(scanner)
	result := twoSum(array, targetSum)
	if len(result) == 0 {
		fmt.Println("None")
	} else {
		fmt.Print(result[0])
		fmt.Print(" ")
		fmt.Print(result[1])
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
