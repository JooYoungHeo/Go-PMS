package malygos

func SortingSlice(list []int) []int {
	len := len(list)
	i := 0

	for i < len-1 {
		flag := true
		target := list[i]
		j := i + 1

		for j < len {
			to := list[j]

			if target > to {
				list[i], list[j] = list[j], list[i]
				flag = false
				break
			}

			j++
		}

		if flag {
			i++
		}
	}

	return list
}
