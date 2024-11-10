package main

import (
	"encoding/json"

	"github.com/gleich/lumber/v3"
	"github.com/gleich/squeue/internal/secrets"
	"github.com/gleich/squeue/internal/spotify"
)

func main() {
	secrets.Load()
	tokens := spotify.LoadTokens()
	tokens.RefreshIfNeeded()

	currentSong, err := spotify.GetQueue(tokens)
	if err != nil {
		lumber.Fatal(err, "failed to load current song")
	}

	debugBytes, _ := json.Marshal(currentSong)
	lumber.Debug(string(debugBytes))
}
