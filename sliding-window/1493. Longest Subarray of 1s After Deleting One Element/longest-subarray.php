<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function longestSubarray($nums) {
        $max = 0;
        $left = 0;
        $right = 0;
        $size = 0;
        $excluded = 0;
        while ($right < count($nums)) {
            if ($nums[$right]==1) {
                $size++;
                if ($size>$max) {
                    $max = $size;
                }
            }
            if ($nums[$right]==0) {
                // если excluded не в окне, один ноль можем пропустить
                if ($excluded<=$left) {
                    $excluded=$right;

                    // если excluded в окне, двигаем left до него
                } else {
                    while ($left < $excluded) {
                        $left++;
                        $size--;
                    }
                    $excluded = $left;
                    // $size = 0;
                    // $left=$right;

                }
            }
            echo "$right ({$nums[$right]}). $size.  [$left – $right] excluded $excluded\n";
            $right++;
        }
        if ($excluded==0) {
            $max--;
        }

        return $max;
    }
}