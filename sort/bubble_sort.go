package sort

func BubbleSort(arr []int) {
	lastUnsortedIndex := len(arr)
	for swapped := true; swapped; lastUnsortedIndex-- {
		swapped = false
		for i := 1; i < lastUnsortedIndex; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
	}
}
