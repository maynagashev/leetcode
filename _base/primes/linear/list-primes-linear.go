package main

func main() {

}

func getLeastPrimesLinear(n int) ([]int, []int) {
	lp := make([]int, n+1)
	primes := []int{}
	for i := 2; i <= n; i++ {
		if lp[i] == 0 {
			lp[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			x := p * i
			if p > lp[i] || x > n {
				break
			}
			lp[x] = p
		}
	}
	return primes, lp
}

func getLeastPrimesLinearWithComments(n int) ([]int, []int) {
	// Создаем массив lp, в котором будем хранить наименьший простой делитель для каждого числа от 2 до n.
	// Создаем массив primes, в котором будем хранить все простые числа от 2 до n.
	lp := make([]int, n+1)
	primes := []int{}

	// Начинаем перебирать все числа от 2 до n.
	for i := 2; i <= n; i++ {
		// Если lp[i] == 0, это означает, что число i является простым числом.
		// Было ли число i уже помечено как простое (то есть, имеет ли lp[i] значение 0 до этого момента, в предыдущих итерациях
		if lp[i] == 0 {
			// Записываем i в lp[i] и добавляем i в массив primes.
			lp[i] = i
			primes = append(primes, i)
		}

		// Для каждого простого числа p, начинаем умножать его на все числа от 2 до n.
		for _, p := range primes {
			// Вычисляем произведение p * i.
			x := p * i
			// Оптимизация лишних вычислений: Если p > lp[i] или произведение больше n, прерываем цикл.
			if p > lp[i] || x > n {
				break
			}
			// Записываем p в lp[x], если p является наименьшим простым делителем числа x.
			lp[x] = p
		}
	}

	// Возвращаем массив простых чисел primes и массив наименьших простых делителей lp.
	return primes, lp
}
