package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	transformed := map[string]int{}

	for point, set := range in {
		for _, v := range set {
			transformed[strings.ToLower(v)] = point
		}
	}

	return transformed
}
