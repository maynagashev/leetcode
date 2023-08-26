<?php

class Solution {

    function increasingTriplet($nums) {
        $smallest = PHP_INT_MAX;
        $secondSmallest = PHP_INT_MAX;

        foreach($nums as $num) {
            if ($num <= $smallest) {
                $smallest = $num;
            } elseif ($num <= $secondSmallest) {
                $secondSmallest = $num;
            } else {
                return true;
            }
        }
        return false;

    }
    /**
     * @param Integer[] $nums
     * @return Boolean
     */
    function increasingTriplet2($nums) {
        $i = $j = $k = null;
        foreach($nums as $v) {
            if (is_null($i) || $v < $i) {
                $i = $v;
            }
            if ((is_null($j) || $v < $j) && $v > $i) {
                $j = $v;
            }
            if ((is_null($k) || $v < $k) && !is_null($j) && $v > $j) {
                $k = $v;
            }
            if (!is_null($k) && $i<$j && $j<$k) {
                // echo "$i $j $k";
                return true;
            }

        }
        // echo "$i $j $k";
        return false;
    }
}