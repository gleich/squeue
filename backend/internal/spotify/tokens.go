package spotify

import (
	"time"

	"github.com/gleich/squeue/internal/secrets"
)

type tokens struct {
	Access    string `json:"access_token"`
	Refresh   string `json:"refresh_token"`
	ExpiresAt int64  `json:"expires_at"` // expected to not be in the refresh data (aka nil)
}

func loadTokens() tokens {
	return tokens{
		Access:    secrets.SECRETS.AccessToken,
		Refresh:   secrets.SECRETS.RefreshToken,
		ExpiresAt: 0,
	}
}

func (t *tokens) refreshIfNeeded() {
	// subtract 30 to ensure that token gets refreshed a few seconds before expiring
	if t.ExpiresAt-5 >= time.Now().Unix() {
		return
	}
}
