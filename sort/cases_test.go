package sort

var testCases = []struct {
	unsorted []int
	sorted   []int
	massage  string
}{
	{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
		"already sorted",
	},
	{
		[]int{5, 4, 3, 2, 1},
		[]int{1, 2, 3, 4, 5},
		"reversed ordered",
	},
	{
		[]int{1, 3, 3, 2, 5},
		[]int{1, 2, 3, 3, 5},
		"has repeat item",
	},
	{
		[]int{2, 1, 3, 0, 5},
		[]int{0, 1, 2, 3, 5},
		"random order",
	},
}
