package plant

import "sort"

func SortByScore(data []Plant) []Plant {
	// Soring plants (worst to best)
	sort.Slice(data, func(i, j int) bool {
		return len(data[i].Segments) < len(data[j].Segments)
	})

	return data
}