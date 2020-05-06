package player

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (r *RadioooooPlayer) requestSongInfo() (string, error) {
	payload, err := r.params.GetRequestPayload()
	if err != nil {
		return "", fmt.Errorf("could not prepare request payload: %v", err)
	}
	response, responseBody, errs := r.requestAgent.
		Post(r.url).
		Send(payload).
		End()
	if errs != nil {
		return "", fmt.Errorf("could not perform request with url \"%s\": %v", r.url, errs)
	}
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("could not perform request with url \"%s\" (status %d): %v", r.url, response.StatusCode, errs)
	}
	return responseBody, nil
}

func (r *RadioooooPlayer) makeSong(info map[string]interface{}) (*Song, error) {
	links, ok := info["links"]
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
	if artist, ok := info["artist"]; ok {
		song.Artist = artist.(string)
	}
	if album, ok := info["album"]; ok {
		song.Album = album.(string)
	}
	if title, ok := info["title"]; ok {
		song.Title = title.(string)
	}
	if year, ok := info["year"]; ok {
		song.Year = year.(string)
	}
	if country, ok := info["country"]; ok {
		song.Country = country.(string)
	}
	if mood, ok := info["mood"]; ok {
		song.Mood = mood.(string)
	}
	return song, nil
}

func (r *RadioooooPlayer) GetNextSong() (*Song, error) {
	songResponse, err := r.requestSongInfo()
	if err != nil {
		return nil, fmt.Errorf("could not request song info: %v", err)
	}

	var songInfo map[string]interface{}
	err = json.Unmarshal([]byte(songResponse), &songInfo)
	if err != nil {
		return nil, fmt.Errorf("could not parse json with song info: %v", err)
	}
	if errorMsg, ok := songInfo["error"]; ok {
		return nil, &ErrorInJSON{errorMsg.(string)}
	}

	return r.makeSong(songInfo)
}
