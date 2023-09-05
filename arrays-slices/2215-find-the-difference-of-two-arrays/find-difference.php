<?php

class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Integer[][]
     */
    function findDifference($nums1, $nums2) {
        return [array_unique(array_diff($nums1, $nums2)), array_unique(array_diff($nums2, $nums1))];
    }
}