package _5_sort_colors

func sortColors(colors []int) []int {

	red := 0
	white := 0
	blue := len(colors) - 1

	for white <= blue {
		if colors[white] == 0 {
			if colors[red] != 0 {
				colors[red], colors[white] = colors[white], colors[red]
			}
			white++
			red++
		} else if colors[white] == 1 {
			white++
		} else {
			if colors[blue] != 2 {
				colors[white], colors[blue] = colors[blue], colors[white]
			}
			blue--
		}
	}

	return colors
}
