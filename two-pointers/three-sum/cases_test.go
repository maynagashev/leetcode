package three_sum

type testCase struct {
	description string
	input       []int
	target      int
	expected    bool
}

var testCases = []testCase{
	{
		description: "thirty one",
		input:       []int{1, 2, 4, 6, 8, 20},
		target:      31,
		expected:    false,
	},
	{
		description: "six",
		input:       []int{1, 2, 3},
		target:      6,
		expected:    true,
	},
}
