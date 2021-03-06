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
	Player    []string
}

func New(pMode string, pMoods []string, pDecades []Decade, pCountries []string) (*Params, error) {
	// mode
	mode, err := canonicalizeMode(pMode)
	if err != nil {
		return nil, fmt.Errorf("could not canonicalize mode \"%s\": %v", pMode, err)
	}

	// moods
	var moods []string
	if len(pMoods) == 0 {
		moods = []string{MOOD_SLOW, MOOD_FAST, MOOD_WEIRD}
	} else {
		moods, err = canonicalizeMoods(pMoods)
		if err != nil {
			return nil, fmt.Errorf("could not canonicalize moods %v: %v", pMoods, err)
		}
	}

	// countries
	if mode == MODE_EXPLORE && len(pCountries) > 1 {
		return nil, fmt.Errorf("more than one country specified for mode \"%s\" (use \"%s\")", mode, MODE_TAXI)
	}
	countries, err := canonicalizeCountries(pCountries)
	if err != nil {
		return nil, fmt.Errorf("could not canonicalize countries codes %v: %v", pCountries, err)
	}

	// decades
	if mode == MODE_EXPLORE && len(pDecades) > 1 {
		return nil, fmt.Errorf("more than one decade specified for mode \"%s\" (use \"%s\")", mode, MODE_TAXI)
	}
	decades, err := canonicalizeDecades(pDecades)
	if err != nil {
		return nil, fmt.Errorf("could not canonicalize decades %v: %v", pDecades, err)
	}

	countriesLen := len(countries)
	if (mode == MODE_TAXI || mode == MODE_EXPLORE) && (len(decades) == 0 || countriesLen == 0) {
		errMsg := fmt.Sprintf("decades and countries must be specified for mode \"%s\"", mode)
		if countriesLen == 0 {
			return nil, &NoCountryError{errMsg}
		}
		return nil, fmt.Errorf(errMsg)
	}

	return &Params{
		Mode:      mode,
		Moods:     moods,
		Decades:   decades,
		Countries: countries,
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
