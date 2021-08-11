package meeting_rooms

import (
	"sort"
)

func CanAttendMeetings(intervals [][]int) bool {
	// sort by start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0]<intervals[j][0]
	})
	for i := 0; i < len(intervals)-1; i++ {
		// if first interval ends later than second stars, terminating
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}