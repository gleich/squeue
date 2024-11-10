package spotify

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gleich/lumber/v3"
)

type CurrentSong struct {
	Timestamp  int64 `json:"timestamp"`
	ProgressMs int   `json:"progress_ms"`
	Item       struct {
		Album struct {
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			TotalTracks          int    `json:"total_tracks"`
			Type                 string `json:"type"`
		} `json:"album"`
		Artists []struct {
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"artists"`
		DurationMs int  `json:"duration_ms"`
		Explicit   bool `json:"explicit"`
	} `json:"item"`
}

func GetCurrentSong(t tokens) (CurrentSong, error) {
	req, err := http.NewRequest(
		"GET",
		"https://api.spotify.com/v1/me/player/currently-playing",
		nil,
	)
	if err != nil {
		lumber.Error(err, "failed to create request")
		return CurrentSong{}, err
	}
	req.Header.Add("Authorization", "Bearer "+t.Access)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		lumber.Error(err, "failed to send request")
		return CurrentSong{}, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		lumber.Error(err, "reading response body failed")
		return CurrentSong{}, nil
	}
	if resp.StatusCode != http.StatusOK {
		lumber.ErrorMsg(resp.StatusCode, "when trying to get new token data:", string(body))
		return CurrentSong{}, nil
	}

	var currentSong CurrentSong
	err = json.Unmarshal(body, &currentSong)
	if err != nil {
		lumber.ErrorMsg(resp.StatusCode, "when trying to get currently playing song", string(body))
		return CurrentSong{}, err
	}

	return currentSong, nil
}
