<?php
class Solution {

    /**
     * @param Integer[] $arr
     * @return Boolean
     */
    function uniqueOccurrences($arr) {
        $occurences = [];
        $counts = [];
        foreach($arr as $v) {
            $occurences[$v] += 1;
        }
        foreach($occurences as $v => $count) {
            if (isset($counts[$count])) {
                return false;
            }
            $counts[$count] = true;
        }
        return true;
    }

    /**
     * @param Integer[] $arr
     * @return Boolean
     */
    function uniqueOccurrences2($arr) {
        $freq = array_count_values($arr);
        return count(array_unique($freq)) == count($freq);
    }
}