<?php
class Solution {

    /**
     * @param String $str1
     * @param String $str2
     * @return String
     */
    function gcdOfStrings($str1, $str2) {
        $currentDivisor = '';
        $maxDivisor = '';
        for ($i =0; $i<strlen($str1); $i++) {
            $currentDivisor .= $str1[$i];
            if ($this->divides($currentDivisor, $str1) && $this->divides($currentDivisor, $str2)) {
                $maxDivisor = $currentDivisor;
            }
        }
        return $maxDivisor;
    }

    function divides($divisor, $str) {
        return preg_match("/^($divisor)+$/", $str);
    }
}