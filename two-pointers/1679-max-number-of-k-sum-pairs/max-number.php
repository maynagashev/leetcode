<?php
class Solution {
    /**
     *
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer
     * 3, 1, 3, 4, 3
     * 1, 2, 3, 4
     */
    function maxOperations($nums, $k) {
        $numFreq = [];
        $operations = 0;

        foreach ($nums as $num) {
            $complement = $k - $num;

            // Если комплимент уже встречался ранее и доступен, засчитываем операцию, берем след элемент
            if (isset($numFreq[$complement]) && $numFreq[$complement] > 0) {
                $operations++;
                $numFreq[$complement]--;
                // если комплимент ранее не встречался, добавляем текущий элемент в копилку и идем дальше
            } else {
                if (isset($numFreq[$num])) {
                    $numFreq[$num]++;
                } else {
                    $numFreq[$num] = 1;
                }
            }
        }

        return $operations;
    }
}