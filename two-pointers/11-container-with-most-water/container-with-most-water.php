<?php

class Solution {

    /**
     * @param Integer[] $height
     * @return Integer
     */
    function maxArea($height) {
        // идем с двух сторон выбирая большую высоту, вычисляем каждый раз объем и сравниваем с максимальным найденным
        $maxArea = 0;
        $leftPointer = 0;
        $rightPointer = count($height)-1;
        while ($leftPointer < $rightPointer) {
            $maxArea = max($maxArea, min($height[$leftPointer], $height[$rightPointer]) * ($rightPointer-$leftPointer));
            // Если левая сторона выше, двигаем правую навстречу левой
            if ($height[$leftPointer] > $height[$rightPointer]) {
                $rightPointer--;
            } else {
                $leftPointer++;
            }
        }

        return $maxArea;
    }
}