/*
[[10,15,12,14,18,5],[5,12,18,13,15,8],[4,7,4,2,10,18],[20,9,9,19,20,5],[10,15,10,15,16,20],[9,6,11,10,12,11],[7,10,6,12,20,8],[3,4,4,18,10,2]] => 41
[[8,16,12,18,9],[19,18,8,2,8],[8,5,5,13,4],[15,9,3,19,2],[8,7,1,8,17],[8,2,8,15,5],[8,17,1,15,3],[8,8,5,5,16],[2,2,18,2,9]] => 28
FAILS: [[3,20,7,7,16,8,7,12,11,19,1],[10,14,3,3,9,13,4,12,14,13,1],[10,1,14,11,1,16,2,7,16,7,19],[13,20,17,15,3,13,8,10,7,8,9],[4,14,18,15,11,9,19,3,15,12,15],[14,12,16,19,2,12,13,3,11,10,9],[18,12,10,16,19,9,18,4,14,2,4]] => 15
*/
package paint_house_2

import "fmt"
func minCostII(costs [][]int) int {
	minCost := -1
	excludeOnTop := -1
	for i:=0; i<len(costs); i++ {
		for startColorIndex, _ := range costs[0] {
			lastPickedColor := startColorIndex -1
			iterationMinCost := 0
			fmt.Println("Iteration:", lastPickedColor)
			for house, colors := range costs {
				if house == 0 {
					lastPickedColor = excludeOnTop
				}
				lastPickedColor = pickMinCostColor(colors, lastPickedColor)
				iterationMinCost += colors[lastPickedColor]
				if house == len(costs)-2 {
					excludeOnTop = lastPickedColor
				}
			}
			if minCost<0 || iterationMinCost<minCost {
				minCost = iterationMinCost
			}
			fmt.Println("iterationMinCost:", iterationMinCost)
		}

		// rotate
		costs = append(costs[1:], costs[0])
		fmt.Println("Rotated:", costs)
	}
	return minCost
}

func pickMinCostColor(colors []int, excludeColor int) int {
	fmt.Println("pickMinCostColor from:", colors, "exclude:", excludeColor)
	minCostColor := 0
	if excludeColor == minCostColor {
		minCostColor++
	}
	for color, cost := range colors {
		if color == excludeColor {
			continue
		}

		if cost < colors[minCostColor] {
			minCostColor = color
		}
	}
	fmt.Println("picked color:", minCostColor, "with cost:", colors[minCostColor])
	return minCostColor
}