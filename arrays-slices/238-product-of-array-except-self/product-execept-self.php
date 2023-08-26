<?php

class Solution
{

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function productExceptSelf($nums)
    {

// 1 2 3 4
        $left = [1];
        for ($i = 0; $i < array_key_last($nums); $i++) {
            $left[$i + 1] = $left[$i] * $nums[$i];
        }

// 1 1 2 6
// print_r($left);
        $right = end($nums); // 4
        for ($i = array_key_last($nums) - 1; $i >= 0; $i--) {
            $left[$i] *= $right;
            $right *= $nums[$i];
        }
// print_r($left);
        return $left;
    }
}