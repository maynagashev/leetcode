package main

import (
	"fmt"
	meeting_rooms "leetcode/slices/252-meeting-rooms"
)

func main() {
	intervals := [][]int{{0,30},{5,10},{15,20}}
	fmt.Println(meeting_rooms.CanAttendMeetings(intervals))

}