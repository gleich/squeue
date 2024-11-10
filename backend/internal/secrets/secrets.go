package secrets

import (
	"github.com/caarlos0/env/v11"
	"github.com/gleich/lumber/v3"
	"github.com/joho/godotenv"
)

var SECRETS Secrets

type Secrets struct {
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`

	AccessToken  string `env:"ACCESS_TOKEN"`
	RefreshToken string `env:"REFRESH_TOKEN"`
}

func Load() {
	err := godotenv.Load()
	if err != nil {
		lumber.Fatal(err, "loading .env file failed")
	}
	loadedSecrets, err := env.ParseAs[Secrets]()
	if err != nil {
		lumber.Fatal(err, "parsing required env vars failed")
	}
	SECRETS = loadedSecrets
	lumber.Done("loaded secrets")
}
