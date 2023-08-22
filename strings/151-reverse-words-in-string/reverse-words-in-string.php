<?php

class Solution {

    /**
     * @param String $s
     * @return String
     */
    function reverseWords($s) {
        $words = explode(" ", $s);
        $reversedArray = [];
        for($i=count($words); $i>=0; $i--) {
            $w = trim($words[$i]);
            if ($w=='') {
                continue;
            }
            $reversedArray[] = $w;
        }
        return implode(" ", $reversedArray);
    }
}