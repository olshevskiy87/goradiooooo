package params

import (
	"fmt"
	"strings"
)

const (
	MOOD_SLOW  = "SLOW"
	MOOD_FAST  = "FAST"
	MOOD_WEIRD = "WEIRD"
)

func canonicalizeMoods(pMoods []string) ([]string, error) {
	var (
		moodsMap = make(map[string]bool)
		moods    []string
	)
	for _, mood := range pMoods {
		var moodUpper = strings.ToUpper(mood)
		if moodUpper != MOOD_SLOW && moodUpper != MOOD_FAST && moodUpper != MOOD_WEIRD {
			return nil, fmt.Errorf("unknown mood \"%s\"", mood)
		}
		if _, ok := moodsMap[moodUpper]; !ok {
			moodsMap[moodUpper] = true
			moods = append(moods, moodUpper)
		}
	}
	return moods, nil
}
