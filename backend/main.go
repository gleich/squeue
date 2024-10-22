package main

import (
	"github.com/gleich/lumber/v3"
	"github.com/gleich/squeue/internal/secrets"
)

func main() {
	lumber.Debug("booted")
	secrets.Load()
	lumber.Debug(secrets.SECRETS.ClientID)
	lumber.Debug(secrets.SECRETS.ClientSecret)
}
