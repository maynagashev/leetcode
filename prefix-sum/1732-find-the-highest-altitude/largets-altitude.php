<?php

class Solution {

    /**
     * @param Integer[] $gain
     * @return Integer
     */
    function largestAltitude($gain) {
        $max = 0;
        $current = 0;
        foreach($gain as $g) {
            $current += $g;
            if ($current>$max) {
                $max = $current;
            }
        }
        return $max;
    }
}