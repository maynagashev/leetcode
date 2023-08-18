package main

import "strings"

func main() {

}

// func mergeAlternately1(word1 string, word2 string) string {
//     res := ""
//     max := len(word1)
//     if (len(word2) > max) {
//         max = len(word2)
//     }
//     for i := 0; i<max; i++ {
//         if (i < len(word1)) {
//             res += string(word1[i])
//         }
//         if (i < len(word2)) {
//             res += string(word2[i])
//         }
//     }
//     return res
// }

func mergeAlternately(word1 string, word2 string) string {
	sb := strings.Builder{}
	max := len(word1)
	if len(word2) > max {
		max = len(word2)
	}
	for i := 0; i < max; i++ {
		if i < len(word1) {
			sb.WriteByte(word1[i])
		}
		if i < len(word2) {
			sb.WriteByte(word2[i])
		}
	}
	return sb.String()
}
