package ring

func occurences[T comparable](s []T) map[T]int {
	res := make(map[T]int)

	for _, value := range s {
		if _, exists := res[value]; exists {
			res[value]++
			continue
		}

		res[value] = 1
	}

	return res
}

func deleteAtIndex[T any](s *[]T, ix int) {
	if len(*s) == 0 {
		return
	}

	copy((*s)[ix:], (*s)[ix+1:])
	*s = (*s)[:len(*s)-1]
}
