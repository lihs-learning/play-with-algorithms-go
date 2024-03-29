package sort

func MergeSort(arr []int) {
	var arrL, arrR []int
	if len(arr) < 2 {
		return
	}

	mid := len(arr) / 2

	arrL = arr[:mid]
	arrR = arr[mid:]

	MergeSort(arrL)
	MergeSort(arrR)

	if arrL[len(arrL)-1] <= arrR[0] {
		return
	}

	arrSorted := make([]int, len(arr))
	for pos, posL, posR := 0, 0, 0; pos < len(arrSorted); pos++ {
		switch {
		case posL >= len(arrL) && posR < len(arrR):
			arrSorted[pos] = arrR[posR]
			posR++
		case posR >= len(arrR) && posL < len(arrL):
			arrSorted[pos] = arrL[posL]
			posL++
		case arrL[posL] < arrR[posR]:
			arrSorted[pos] = arrL[posL]
			posL++
		case arrL[posL] >= arrR[posR]:
			arrSorted[pos] = arrR[posR]
			posR++
		}
	}
	for pos := range arr {
		arr[pos] = arrSorted[pos]
	}
}
