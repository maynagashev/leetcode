<?php

class Solution
{

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Float
     */
    function findMaxAverage($nums, $k)
    {
        $maxSum = 0;
        $sum = 0;
        for ($i = 0; $i < count($nums); $i++) {
            // наполняем окно
            if ($i < $k) {
                $sum += $nums[$i];
                $maxSum = $sum;
                continue;
            }
            // дальше сдвигаем окно, добавляем элемент, удаляем начало
            $sum -= $nums[$i - $k];
            $sum += $nums[$i];
            if ($sum > $maxSum) {
                $maxSum = $sum;
            }

        }
        return $maxSum / $k;
    }
}