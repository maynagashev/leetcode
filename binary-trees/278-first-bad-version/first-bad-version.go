package _78_first_bad_version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return true
}

// 1, 2, 3, 4*, 5
func firstBadVersion(n int) int {
	first := 1
	last := n
	mid := (first + last) / 2

	if isBadVersion(mid) {
		// середина плохая, исключаем из поиска правую часть, оставляем середину как первую известную плохую версию
		last = mid
	} else {
		// середина хорошая, откидываем левую часть включая середину
		first = mid + 1
	}

	for first < last {
		mid := (first + last) / 2
		if isBadVersion(mid) {
			last = mid
		} else {
			first = mid + 1
		}
	}

	return first
}
