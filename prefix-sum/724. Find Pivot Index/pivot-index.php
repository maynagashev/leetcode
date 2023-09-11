<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function pivotIndex($nums) {
        $sum = array_reduce($nums, fn($carry, $v) => $carry += $v, 0);
        $leftSum = 0;
        foreach($nums as $i => $v) {
            if ($leftSum == $sum - $v - $leftSum) {
                return $i;
            }
            $leftSum += $v;
        }
        return -1;
    }
}