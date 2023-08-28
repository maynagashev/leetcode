<?php

class Solution {

    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     * abc   ahbgdc
     */
    function isSubsequence($s, $t) {
        if ($s == '') {
            return true;
        }
        $sArr = str_split($s);
        $currentChar = current($sArr);
        for ($i = 0; $i<strlen($t); $i++){
            if ($t[$i] == $currentChar) {
                $currentChar = next($sArr);
                if (!$currentChar) {
                    return true;
                }
            }
        }
        // echo "currentChar=$currentChar";
        return false;
    }

}