package plant

import "sort"

func SortByScore(data []Plant) []Plant {
	// Soring blobs (worst to best)
	sort.Slice(data, func(i0, i1 int) bool {
		return data[i0].Points < data[i1].Points
	})
	return data
}
