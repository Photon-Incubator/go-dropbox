package dropbox

import (
	"context"
	"net/http"

	"google.golang.org/appengine/urlfetch"
)

// Config for the Dropbox clients.
type Config struct {
	HTTPClient  *http.Client
	AccessToken string
}

// NewConfig with the given access token.
func NewConfig(ctx context.Context, accessToken string) *Config {
	return &Config{
		HTTPClient:  urlfetch.Client(ctx),
		AccessToken: accessToken,
	}
}
