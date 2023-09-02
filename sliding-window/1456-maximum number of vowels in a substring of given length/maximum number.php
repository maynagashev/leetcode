<?php
class Solution {

    /**
     * @param String $s
     * @param Integer $k
     * @return Integer
     */
    function maxVowels($s, $k) {
        $vowels = ['a', 'e', 'i', 'o', 'u'];
        $maxVowels = 0;
        $currentVowels = 0;
        $letters = str_split($s);
        for($i=0; $i<count($letters); $i++) {
            // наполняем окно
            if ($i<$k) {
                if (in_array($letters[$i], $vowels)) {
                    $maxVowels = ++$currentVowels;
                }
                continue;
            }
            // двигаем окно (если выбыла гласная отнимаем, если добавилась прибавляем)
            if (in_array($letters[$i-$k], $vowels)) {
                $currentVowels--;
            }
            if (in_array($letters[$i], $vowels)) {
                $currentVowels++;
            }
            if ($currentVowels>$maxVowels) {
                $maxVowels = $currentVowels;
            }
        }
        return $maxVowels;
    }
}