package main

func main() {

}

func reverse(x int) int {
	s := 0
	for x != 0 {
		t := x % 10
		x = x / 10
		s =  s * 10 + t
	}

	if s < -2147483648 || s > 2147483647 {
		return 0
	}
	return s
}