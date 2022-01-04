package sort

func InsertSort(arr []int) {
	for currentIndex, current := range arr {
		position := currentIndex
		for position > 0 && arr[position-1] > current {
			arr[position] = arr[position-1]
			position--
		}
		arr[position] = current
	}
}
