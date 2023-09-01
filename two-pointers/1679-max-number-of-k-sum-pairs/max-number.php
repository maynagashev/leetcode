<?php
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer
     */
    function maxOperations($nums, $k) {
        $max = 0;
        while($this->hasPairs($nums, $k)){
            $max++;
        }
        return $max;
    }

    function hasPairs(&$nums, &$k) {
        foreach($nums as $i => $v1){
            foreach($nums as $j => $v2){
                if ($i == $j) {
                    continue;
                }
                // если сумма найдена, удаляем элементы
                if ($v1+$v2 == $k) {
                    unset($nums[$i]);
                    unset($nums[$j]);
                    return true;
                }
            }
        }
        return false;
    }
}