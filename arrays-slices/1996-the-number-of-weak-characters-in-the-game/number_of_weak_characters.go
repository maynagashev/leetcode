/*
https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/
*/
package main

import "sort"

func main() {

}

/*
Time: O(NlogN) - sort + iterate array
Space: O(1) - sort in place + constant number of variables
*/
func numberOfWeakCharacters(properties [][]int) int {

	// Sort by attack DESC (strongest at the beginning) & defense ASC
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1] // defence asc
		}
		return properties[i][0] > properties[j][0] // attack desc
	})

	weakCounter := 0
	maxDefenseBefore := 0 // till now

	for i := 0; i < len(properties); i++ {
		// first character will be skipped because we will only fill maxDefense in the first iteration
		// starting from second character - all chars will be possibly weaker by attack
		// and we checking only defense of previous characters
		if properties[i][1] < maxDefenseBefore {
			// here we know that defense is lower than before (some char before has defense higher),
			// and that all characters before have better attack (because of sort)
			weakCounter++
		} else if properties[i][1] > maxDefenseBefore {
			maxDefenseBefore = properties[i][1]
		}
		// else: current defense even with maxDefense - skipping
	}

	return weakCounter
}
