package params

import (
	"encoding/json"
	"fmt"
)

type Params struct {
	Mode      string
	Moods     []string
	Decades   []Decade
	Countries []string
}

func New(pMode string, pMoods []string, pDecades []Decade, pCountries []string) (*Params, error) {
	mode, err := canonicalizeMode(pMode)
	if err != nil {
		return nil, fmt.Errorf("could not canonicalize mode \"%s\": %v", pMode, err)
	}

	var moods []string
	if len(pMoods) == 0 {
		moods = []string{MOOD_SLOW, MOOD_FAST, MOOD_WEIRD}
	} else {
		moods, err = canonicalizeMoods(pMoods)
		if err != nil {
			return nil, fmt.Errorf("could not canonicalize moods %v: %v", pMoods, err)
		}
	}

	countries, err := canonicalizeCountries(pCountries)
	if err != nil {
		return nil, fmt.Errorf("could not canonicalize countries codes %v: %v", pCountries, err)
	}

	decades, err := canonicalizeDecades(pDecades)
	if err != nil {
		return nil, fmt.Errorf("could not canonicalize decades %v: %v", pDecades, err)
	}

	if (mode == MODE_TAXI || mode == MODE_EXPLORE) && (len(decades) == 0 || len(countries) == 0) {
		return nil, fmt.Errorf("decades and countries must be specified for mode \"%s\"", pMode)
	}

	return &Params{
		mode,
		moods,
		decades,
		countries,
	}, nil
}

func (p *Params) GetRequestPayload() (string, error) {
	payload := map[string]interface{}{
		"mode":  p.Mode,
		"moods": p.Moods,
	}
	if p.Mode == MODE_EXPLORE || p.Mode == MODE_TAXI {
		payload["decades"] = p.Decades
		payload["isocodes"] = p.Countries
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("could not create request payload json: %v", err)
	}
	return string(payloadJSON), nil
}
