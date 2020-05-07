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
	if links == nil {
		return nil, fmt.Errorf("empty \"links\" in song info")
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
		if link != nil {
			song.Links[format] = link.(string)
		}
	}
	if artist, ok := info["artist"]; ok && artist != nil {
		song.Artist = artist.(string)
	}
	if album, ok := info["album"]; ok && album != nil {
		song.Album = album.(string)
	}
	if title, ok := info["title"]; ok && title != nil {
		song.Title = title.(string)
	}
	if year, ok := info["year"]; ok && year != nil {
		song.Year = year.(string)
	}
	if country, ok := info["country"]; ok && country != nil {
		song.Country = country.(string)
	}
	if mood, ok := info["mood"]; ok && mood != nil {
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
	if errorValue, ok := songInfo["error"]; ok {
		var msg string
		if errorValue != nil {
			msg = errorValue.(string)
		} else {
			msg = "unknow error"
		}
		return nil, &ErrorInJSON{msg}
	}

	return r.makeSong(songInfo)
}
