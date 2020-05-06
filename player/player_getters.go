package player

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (r *RadioooooPlayer) GetNextSongLink() (*Song, error) {
	payload, err := r.params.GetRequestPayload()
	if err != nil {
		return nil, fmt.Errorf("could not prepare request payload: %v", err)
	}

	response, responseBody, errs := r.requestAgent.
		Post(r.url).
		Send(payload).
		End()
	if errs != nil {
		return nil, fmt.Errorf("could not perform request with url \"%s\": %v", r.url, errs)
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not perform request with url \"%s\" (status %d): %v", r.url, response.StatusCode, errs)
	}

	var songInfo map[string]interface{}
	err = json.Unmarshal([]byte(responseBody), &songInfo)
	if err != nil {
		return nil, fmt.Errorf("could not parse json: %v", err)
	}
	if errorMsg, ok := songInfo["error"]; ok {
		return nil, &ErrorInJSON{errorMsg.(string)}
	}

	links, ok := songInfo["links"]
	if !ok {
		return nil, fmt.Errorf("no key \"links\" in song info")
	}

	linksStr := links.(map[string]interface{})
	linksLen := len(linksStr)
	if linksLen == 0 {
		return nil, fmt.Errorf("no song links in song info")
	}

	song := &Song{
		Links: make(map[string]string, linksLen),
	}
	for format, link := range linksStr {
		song.Links[format] = link.(string)
	}
	if artist, ok := songInfo["artist"]; ok {
		song.Artist = artist.(string)
	}
	if album, ok := songInfo["album"]; ok {
		song.Album = album.(string)
	}
	if title, ok := songInfo["title"]; ok {
		song.Title = title.(string)
	}
	if year, ok := songInfo["year"]; ok {
		song.Year = year.(string)
	}
	if country, ok := songInfo["country"]; ok {
		song.Country = country.(string)
	}
	if mood, ok := songInfo["mood"]; ok {
		song.Mood = mood.(string)
	}

	return song, nil
}
