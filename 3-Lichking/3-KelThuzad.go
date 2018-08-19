package kelthuzad

func CheckIndex(list []string, target string) int {
	idx := -1

	for i, e := range list {
		if target == e {
			idx = i
			break
		}
	}

	return idx
}
