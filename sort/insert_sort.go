package sort

func InsertSort(arr []int) {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		current := arr[currentIndex]
		position := currentIndex
		for position > 0 && arr[position-1] > current {
			arr[position] = arr[position-1]
			position--
		}
		arr[position] = current
	}
}
