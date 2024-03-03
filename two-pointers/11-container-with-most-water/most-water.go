package _1_container_with_most_water

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	area := 0
	for left < right {
		if height[left] >= height[right] {
			// left bigger
			area = height[right] * (right - left)
			right--
		} else {
			// right bigger
			area = height[left] * (right - left)
			left++
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
