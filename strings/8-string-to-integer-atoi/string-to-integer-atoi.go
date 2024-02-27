package main

import (
	"math"
)

func myAtoi(s string) int {
	var res int
	var sign int

	for i := 0; i < len(s); i++ {
		// если пробел и знак еще не определен пропускаем его
		if s[i] == ' ' && sign == 0 {
			continue
		}

		// обработка знака, если еще не установлен
		if s[i] == '+' && sign == 0 {
			sign = 1
			continue
		}
		if s[i] == '-' && sign == 0 {
			sign = -1
			continue
		}

		// если цифра то пока все ок, заодно устанавливаем знак если не установлен
		if s[i] >= '0' && s[i] <= '9' {
			if sign == 0 {
				sign = 1
			}

			// проверяем лимит, например 2147483646 => 214748364 (6) > 214748364 (7)
			if res+int(s[i]-'0') > math.MaxInt32/10+7 {
				if sign > 0 {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}

			res = res*10 + int(s[i]-'0')
			continue
		}

		// если не пробел и не первый знак и не число, выходим из цикла и возвращаем текущий результат
		break
	}

	return res * sign
}
