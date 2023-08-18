<?php

class Solution {

    /**
     * @param String $word1
     * @param String $word2
     * @return String
     */
    function mergeAlternately($word1, $word2) {
        $w1 = str_split($word1);

        $res = '';
        $len1 = strlen($word1);
        $len2 = strlen($word2);
        $maxLen = max($len1, $len2);
        for($i = 0; $i<$maxLen; $i++) {
            // Если индекс пока в рамках длины первого слова, добавляем букву из него
            if ($i<$len1) {
                $res .= $word1[$i];
            }
            if ($i<$len2) {
                $res .= $word2[$i];
            }
        }
        return $res;
    }
}
