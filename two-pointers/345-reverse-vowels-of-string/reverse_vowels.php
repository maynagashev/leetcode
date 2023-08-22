<?php

class Solution
{

    /**
     * @param String $s
     * @return String
     * asdasdsssaaaaaaaa
     */
    function reverseVowels($s)
    {
        $i = 0;
        $last = strlen($s) - 1;
        while ($i < $last) {
            // если гласная, находим последнюю гласную и меняем местами
            if ($this->isVowel($s[$i])) {
                while ($last > $i) {
                    if ($this->isVowel($s[$last])) {
                        $tmp = $s[$i];
                        $s[$i] = $s[$last];
                        $s[$last] = $tmp;
                        $last--;
                        break;
                    }
                    $last--;
                }
            }
            $i++;
        }
        return $s;
    }

    function isVowel($letter)
    {
        return in_array(strtolower($letter), ['a', 'e', 'i', 'o', 'u']);
    }
}