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
		if posL >= len(arrL) && posR < len(arrR) {
			arrSorted[pos] = arrR[posR]
			posR++
			continue
		}
		if posR >= len(arrR) && posL < len(arrL) {
			arrSorted[pos] = arrL[posL]
			posL++
			continue
		}
		if arrL[posL] < arrR[posR] {
			arrSorted[pos] = arrL[posL]
			posL++
		} else {
			arrSorted[pos] = arrR[posR]
			posR++
		}
	}
	for pos := range arr {
		arr[pos] = arrSorted[pos]
	}
}
