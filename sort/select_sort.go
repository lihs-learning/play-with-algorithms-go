package sort

func SelectSort(arr []int) {
	for orderIndex := range arr {
		targetIndex := orderIndex
		for finderIndex, finder := range arr[orderIndex:] {
			if finder < arr[targetIndex] {
				targetIndex = orderIndex + finderIndex
			}
		}
		arr[orderIndex], arr[targetIndex] = arr[targetIndex], arr[orderIndex]
	}
}
