package plant

import "sort"

func SortByScore(data []Plant) []Plant {
	// Soring plants (worst to best)
	sort.Slice(data, func(i, j int) bool {
		return data[i].Points < data[j].Points
	})

	return data
}