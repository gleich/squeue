package main

import (
	"github.com/gleich/squeue/internal/secrets"
	"github.com/gleich/squeue/internal/spotify"
)

func main() {
	secrets.Load()
	tokens := spotify.LoadTokens()
	tokens.RefreshIfNeeded()
}
