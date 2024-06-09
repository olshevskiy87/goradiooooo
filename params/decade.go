package params

import (
	"fmt"
	"slices"
)

type Decade uint16

const (
	DECADE_MIN = 1900
	DECADE_MAX = 2020
)

var extra_decades = []Decade{
	2070, // Original compositions & visions of the Future by artists of Today.
}

func canonicalizeDecades(pDecades []Decade) ([]Decade, error) {
	var (
		decadesMap = make(map[Decade]bool, len(pDecades))
		decades    []Decade
	)
	for _, decade := range pDecades {
		if decade < DECADE_MIN {
			return nil, fmt.Errorf("decade can't be less than %d", DECADE_MIN)
		} else if decade > DECADE_MAX && !slices.Contains(extra_decades, decade) {
			return nil, fmt.Errorf("decade can't be more than %d", DECADE_MAX)
		}
		if _, ok := decadesMap[decade]; !ok {
			decadesMap[decade] = true
			decades = append(decades, decade)
		}
	}
	return decades, nil
}
