<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     * 121ms (89,74%) 28.85MB (87,18%)
     */
    function longestSubarray($nums) {
        $max = 0;
        $left = 0;
        $right = 0;
        $excluded = -1;

        while ($right < count($nums)) {
            // $row = '';

            // очередная единичка
            if ($nums[$right]==1) {
                $max = max($max, $right - $left + intval($excluded<$left));
            }

            // очередной ноль
            if ($nums[$right]==0) {
                // если excluded не в окне, один ноль можем пропустить
                if ($excluded<$left) {
                    // исключаем текущий 0 и переходим на следующую итерацию
                    $excluded=$right;

                    // если excluded в окне, и мы встретили 0, то left = excluded+1
                } else {
                    // $row.=" '0' в окне (на поз. $excluded), left = excluded+1";
                    $left = $excluded + 1;
                    // исключаем текущий ноль заново
                    $excluded = $right;
                }
            }

            // echo "$right ({$nums[$right]}). $size. [$left – $right] excluded $excluded  $row\n";
            $right++;
        }
        // обработка кейса когда не было нулей
        if ($excluded<0) {
            $max--;
        }

        return $max;
    }
}