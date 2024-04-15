package main

func main() {

}

func eratosthenesEffective(n int) []bool {
	numbers := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		numbers[i] = true
	}
	for num := 2; num < n; num++ {
		if numbers[num] {
			for j := num * num; j <= n; j += num {
				numbers[j] = false
			}
		}
	}
	numbers[0] = false
	numbers[1] = false
	return numbers
}

func eratosthenesEffectiveWithComments(n int) []bool {
	// Создаем срез numbers, в котором будем хранить информацию о том, является ли число простым.
	numbers := make([]bool, n+1)

	// Инициализируем все элементы среза как true, кроме нулевого и первого элементов.
	for i := 2; i <= n; i++ {
		numbers[i] = true
	}

	// Применяем алгоритм Решета Эратосфена.
	for num := 2; num < n; num++ {
		// Если число является простым, то все его кратные числа будут составными.
		if numbers[num] {
			// Обходим все кратные числа num и помечаем их как составные.
			// Начинаем с num^2, потому что все меньшие числа уже были помечены ранее.
			for j := num * num; j <= n; j += num {
				numbers[j] = false
			}
		}
	}

	// Помечаем нулевое и первое числа как составные, так как они не являются простыми.
	numbers[0] = false
	numbers[1] = false

	// Возвращаем срез numbers, содержащий информацию о простых числах до n.
	return numbers
}
