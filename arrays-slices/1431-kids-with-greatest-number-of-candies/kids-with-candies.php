<?php


/**
 * @param Integer[] $candies
 * @param Integer $extraCandies
 * @return Boolean[]
 */
function kidsWithCandies($candies, $extraCandies)
{
    $max = max($candies);
    foreach ($candies as $i => $count) {
        $candies[$i] = $count + $extraCandies >= $max;
    }
    return $candies;
}
