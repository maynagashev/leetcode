package main

func main() {

}


func findFarmland(land [][]int) [][]int {
	cols := len(land)
	rows := len(land[0])

	seenFarms := [][]int{} // [x-left 0, y-top 1, x-right 2, y-bottom 3], [x, y, x, y], ...

	for r:=0; r<rows; r++ { // iterate rows
		startFarmCol := -1
		// endFarmCol := -1

		var currentFarm *[]int
		// fmt.Printf("%+v\n", currentFarm)
		for c:=0; c<cols; c++ { // iterate cols in every row
			// fmt.Println("Coordinate: ", r, c, "value", land[c][r])
			if land[c][r] == 1 {
				if startFarmCol < 0 {
					startFarmCol = c // register left X
					currentFarm = findFarm(c, r, &seenFarms)
					// fmt.Println("search ", r, c, "result", currentFarm)
					if currentFarm == nil {
						farm := []int{startFarmCol, r, startFarmCol, r}
						// fmt.Println("append farm", farm)
						seenFarms = append(seenFarms, farm)
						currentFarm = findFarm(c, r,  &seenFarms)
					}
				}
				updateCurrentFarm(r, c, currentFarm)
			}

			if land[c][r] == 0 {
				if currentFarm != nil {
					(*currentFarm)[2] = c-1 // register right X
					startFarmCol = -1  // reset left
					currentFarm = nil
				}
			}
		}
	}

	return seenFarms;
}

// [x-left 0, y-top 1, x-right 2, y-bottom 3],
func updateCurrentFarm(r int, c int, currentFarm *[]int) {
	// fmt.Println("update", r, c, currentFarm)
	f := *currentFarm
	if r > f[3] {
		f[3] = r
	}

	if c>f[2] {
		f[2] = c
	}

	// fmt.Println("result", currentFarm)

}

func findFarm(leftX int, y int, seenFarms *[][]int) *[]int {
	for _, f := range *seenFarms {
		if f[0] == leftX && y>=f[1]-1 && y<=f[3]+1 {
			return &f
		}
	}
	return nil
}