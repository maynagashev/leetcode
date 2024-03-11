package trapping_rain_water

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	// two pointers
	left, right := 0, len(height)-1
	// leftMax and rightMax are the maximum height of the left and right side of the array
	leftMax, rightMax := height[left], height[right]
	res := 0
	for left < right {
		//
		if leftMax < rightMax {
			left++
			// if leftMax is greater than the current height, then we can trap water
			if leftMax > height[left] {
				res += leftMax - height[left]
			} else {
				// move the leftMax to the current height
				leftMax = height[left]
			}
		} else {
			right--
			// if rightMax is greater than the current height, then we can trap water
			if rightMax > height[right] {
				res += rightMax - height[right]
			} else {
				// move the rightMax to the current height
				rightMax = height[right]
			}
		}
	}
	return res
}
