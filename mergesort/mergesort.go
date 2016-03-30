package MergeSort

func MergeSort(dataToSort []int, threshold int) []int {

	len := len(dataToSort)

	if len <= 1 {
		return dataToSort
	}

	middle := len / 2

	var left []int;
	var right []int;

	if len <= threshold {
		left = MergeSort(dataToSort[:middle], threshold)
		right = MergeSort(dataToSort[middle:], threshold)
	} else{
		go func(data []int){left = MergeSort(dataToSort[:middle], threshold)}();
		go func(data []int){right = MergeSort(dataToSort[middle:], threshold)}();
	}

	return merge(left, right)

}

func merge(leftList, rightList []int) []int {

	size := len(leftList)+len(rightList)
	i, j := 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(leftList)-1 && j <= len(rightList)-1 {
			slice[k] = rightList[j]
			j++
		} else if j > len(rightList)-1 && i <= len(leftList)-1 {
			slice[k] = leftList[i]
			i++
		} else if leftList[i] < rightList[j] {
			slice[k] = leftList[i]
			i++
		} else {
			slice[k] = rightList[j]
			j++
		}
	}
	return slice
}
