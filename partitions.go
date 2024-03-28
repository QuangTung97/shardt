package shardt

func allocatePartitions(count int, addrs []string, previous []string) []string {
	min := count / len(addrs)
	max := min + 1

	maxCount := count - min*len(addrs)

	previousAlloc := map[string][]int

	var result []string
	for i, addr := range addrs {
		if i < maxCount {
			for k := 0; k < max; k++ {
				result = append(result, addr)
			}
		} else {
			for k := 0; k < min; k++ {
				result = append(result, addr)
			}
		}
	}

	return result
}
