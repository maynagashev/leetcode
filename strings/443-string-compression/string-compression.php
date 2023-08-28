<?php

class Solution {

    /**
     * @param String[] $chars
     * @return Integer
     * [ a a a a b b b c d d] => a4b3cd2
     * [ a a b c ] => a2bc
     */
    function compress(&$chars) {
        $insertPosition = 0;
        $currentCount = 0;
        $currentChar = null;
        foreach($chars as $i => $v) {
            if ($v !== $currentChar) {
                // store result
                if (!is_null($currentChar)) {
                    $chars[$insertPosition++] = $currentChar;
                    if ($currentCount>1) {
                        foreach(str_split($currentCount) as $digit) {
                            $chars[$insertPosition++] = $digit;
                        }
                    }
                }

                // reset char
                $currentChar = $v;
                $currentCount = 1;
                continue;
            }
            $currentCount++;
        }
        // store last result
        $chars[$insertPosition++] = $currentChar;
        if ($currentCount>1) {
            foreach(str_split($currentCount) as $digit) {
                $chars[$insertPosition++] = $digit;
            }
        }

        // echo "char=$currentChar, count=$currentCount, insertPosition=$insertPosition\n";
        // print_r($chars);
        return $insertPosition;
    }
}