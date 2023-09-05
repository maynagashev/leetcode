<?php

class Solution {
    // /**
    //  * @param String $s
    //  * @return String run 28%, mem 64%
    //  */
    // function removeStars($s) {

    //     while (isset($s[$i])) {
    //         if($s[$i]==='*') {
    //             $s = substr_replace($s, '', $i-1, 2);
    //             $i -= 2;
    //         }
    //         $i++;
    //     }
    //     return $s;
    // }
    /**
     * @param String $s
     * @return String
     * runtime 82.81% memory 34.38%
     * runtime 70% memory 59% (24.58MB)
     */
    function removeStars($s) {
        $res = [];
        for ($i =0; $i<strlen($s); $i++) {
            if($s[$i]=='*') {
                array_pop($res);
                continue;
            }
            $res[]=$s[$i];
        }
        return implode("", $res);
    }
    //     /**
    //  * @param String $s
    //  * @return String
    //  * timelimit exceeded
    //  */
    // function removeStars($s) {
    //     $s = str_split($s);
    //     $i = 0;
    //     while(isset($s[$i])){
    //         if($s[$i]==='*') {
    //             array_splice($s, $i-1, 2, []);
    //             $i -=2;
    //         }
    //         $i++;
    //     }
    //     return implode("", $s);
    // }
}