<?php

class RecentCounter {

    private SplQueue $requests;

    /**
     */
    function __construct()
    {
        $this->requests = new SplQueue();
    }

    /**
     * @param Integer $t
     * @return Integer
     */
    function ping($t)
    {
        $this->requests->enqueue($t);

        while ($this->requests->bottom() < $t - 3000) {
            $this->requests->dequeue();
        }
        return $this->requests->count();
    }
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * $obj = RecentCounter();
 * $ret_1 = $obj->ping($t);
 */