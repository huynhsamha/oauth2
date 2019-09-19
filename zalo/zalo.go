// Package zalo provides constants for using OAuth2 to access Zalo Social API.
package zalo // import "github.com/huynhsamha/oauth2/zalo"

import (
	"github.com/huynhsamha/oauth2"
)

// Endpoint is Zalo's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://oauth.zaloapp.com/v3/auth",
	TokenURL: "https://oauth.zaloapp.com/v3/access_token",
}
