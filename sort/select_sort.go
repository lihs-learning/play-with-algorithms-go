package sort

func SelectSort(arr []int) {
	for orderIndex, order := range arr {
		targetIndex := orderIndex
		for finderIndex, finder := range arr[orderIndex:] {
			if finder < order {
				targetIndex = orderIndex + finderIndex
			}
		}
		arr[orderIndex], arr[targetIndex] = arr[targetIndex], arr[orderIndex]
	}
}
