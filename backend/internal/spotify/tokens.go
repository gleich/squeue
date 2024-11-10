package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gleich/lumber/v3"
	"github.com/gleich/squeue/internal/secrets"
)

type tokens struct {
	Access    string `json:"access_token"`
	Refresh   string `json:"refresh_token"`
	ExpiresAt int64  `json:"expires_at"`
}

func LoadTokens() tokens {
	return tokens{
		Access:    secrets.SECRETS.AccessToken,
		Refresh:   secrets.SECRETS.RefreshToken,
		ExpiresAt: 0,
	}
}

func (t *tokens) RefreshIfNeeded() {
	// add 60 to ensure that the token doesn't expire in the next 60 seconds
	if t.ExpiresAt-60 >= time.Now().Unix() {
		return
	}

	params := url.Values{
		"refresh_token": {t.Refresh},
		"grant_type":    {"refresh_token"},
		"client_id":     {secrets.SECRETS.ClientID},
	}
	req, err := http.NewRequest(
		"POST",
		"https://accounts.spotify.com/api/token?"+params.Encode(),
		nil,
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set(
		"Authorization",
		"Basic "+base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("%s:%s", secrets.SECRETS.ClientID, secrets.SECRETS.ClientSecret)),
		),
	)
	if err != nil {
		lumber.Error(err, "creating request for new token failed")
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		lumber.Error(err, "sending request for new token data failed")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		lumber.Error(err, "reading response body failed")
		return
	}
	if resp.StatusCode != http.StatusOK {
		lumber.ErrorMsg(resp.StatusCode, "when trying to get new token data:", string(body))
	}

	var tokens tokens
	err = json.Unmarshal(body, &tokens)
	if err != nil {
		lumber.Error(err, "failed to parse json")
		lumber.Debug("body:", string(body))
		return
	}

	*t = tokens
	lumber.Done(fmt.Sprintf("loaded new tokens [access_token = %s]", t.Access))
}
