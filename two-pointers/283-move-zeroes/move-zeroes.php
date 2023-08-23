<?php

class Solution {

    /**
     * https://leetcode.com/problems/move-zeroes/
     * @param Integer[] $nums
     * @return NULL
     * 1 0 0 2 0 3 0 4
     * 0 0 0 0 0 0 0 1
     * 1 0 0 0 0 0 0 0
     * Два указателя:
     *  - первый идет просто по массиву подряд быстро ($i)
     *  - во втором (медленном lastNonZeroIndex) фиксируем куда вставить ненулевые элементы в начале массива, по время следующей итерации
     * space: O(1)
     * time: O(n)
     */
    function moveZeroes(&$nums) {
        $lastNonZeroIndex = 0;
        foreach($nums as $v) {
            if ($v != 0) {
                $nums[$lastNonZeroIndex++] = $v;
            }
        }
        for($i=$lastNonZeroIndex; $i<count($nums); $i++) {
            $nums[$i] = 0;
        }
    }
}