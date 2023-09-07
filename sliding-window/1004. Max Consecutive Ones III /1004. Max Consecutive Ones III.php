<?php

class Solution {

    function longestOnes($nums, $k) {
        $zeros = 0; // Счетчик нулей в текущем окне
        $res = 0;   // Максимальная длина последовательности 1
        $left = 0;  // Левая граница текущего окна

        // Проходим по массиву
        for ($right = 0; $right < count($nums); $right++) {
            // Если встретили 0, начинаем сжимать окно
            if ($nums[$right] === 0) {
                // Пока в окне больше или равно $k нулей, сдвигаем левую границу
                while ($zeros >= $k) {
                    if ($nums[$left] === 0) {
                        $zeros--;
                    }
                    $left++;
                }
                $zeros++; // Учитываем текущий 0, который можно заменить на 1
            }

            // Обновляем максимальную длину последовательности 1
            $res = max($right - $left + 1, $res);
        }

        return $res; // Возвращаем максимальную длину последовательности 1
    }

    /**
     * Не рабочий вариант, проходит только небольшую часть тестов
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer
     */
    function longestOnes2($nums, $k) {
        $availableZeros = $k;
        $max = 0;
        $left = 0;
        $right = 0;
        // в основном цикле ориентируемся на правый указатель, итерируем пока он не дойдет до последнего элемента
        while($right<count($nums)) {

            // если очередной элемент справа 1, проверяем максимум
            if ($nums[$right]===1) {
                $max = max($max, $right-$left+1);
            }

            // если очередной элемент 0
            if ($nums[$right]===0) {
                // и если есть доступные замены, меняем
                if ($availableZeros>0){
                    $availableZeros--;
                    $max = max($max, $right-$left+1);
                } else {
                    // текущий элемент $nums[$right]=0 и доступных нулей больше нет, серия закончилась,
                    // подтягиваем левую сторону до $right или пока не освободятся нули
                    // восстанавливая $availableZeros
                    for($left; $left<=$right; $left++) {
                        // если в текущем окне был ноль, значит мы его меняли, особождаем его, увеличиваем доступные нули в рамках $k
                        if ($nums[$left]===0 && $availabableZeros<$k) {
                            $availableZeros++;
                            // левый указатель сразу сдвигаем дальше и выходим пока
                            $left++;
                            echo "\t $right.  left++, left=$left, available++ = $availableZeros, right=$right, len=".($right-$left+1).", max=$max, avaialableZeros=$availableZeros break;\n";

                            break; // останавливаемся на текущей позиции, можем двинуть правую позицию
                        } else {
                            echo "\t $right. left=$left, right=$right, len=".($right-$left+1).", max=$max, avaialableZeros=$availableZeros\n";
                        }
                        echo "left++ in cycle\n";
                    }
                    echo "left++ in the end\n";
                    if ($left<$right) {
                        $left++;
                    }
                }
            }
            echo "$right. current=$nums[$right], left=$left, right=$right, len=".($right-$left+1).", max=$max, avaialableZeros=$availableZeros\n";
            // сдвигаем правый указатель дальше в любом случае
            $right++;
        }


        return $max;
    }
}