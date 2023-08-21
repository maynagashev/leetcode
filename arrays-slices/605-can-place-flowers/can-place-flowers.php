<?php


/**
 * @param Integer[] $flowerbed
 * @param Integer $n
 * @return Boolean
 * 1 0 0 0 0 0 1
 */
function canPlaceFlowers($flowerbed, $n)
{
    if ($n == 0) {
        return true;
    }
    $availableCount = 0;
    foreach ($flowerbed as $i => $v) {
        $prev = $flowerbed[$i - 1] ?? 0;
        $next = $flowerbed[$i + 1] ?? 0;
        if ($v == 0 && $prev == 0 && $next == 0) {
            $availableCount++;
            $flowerbed[$i] = 1; // помечаем клетку как занятую (т.к. мы ее уже посчитали)
            if ($availableCount >= $n) {
                return true;
            }
        }
    }
    return false;
}