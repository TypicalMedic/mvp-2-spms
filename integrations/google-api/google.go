package googleapi

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type Google struct {
	api GoogleAPI
}

func InintGoogle(googleAPI GoogleAPI) Google {
	return Google{
		api: googleAPI,
	}
}

func (g *Google) GetAuthLink(redirectURI string, state string) string {
	url := g.api.GetAuthLink(redirectURI, state)
	return url
}

func (g *Google) GetToken(code string) *oauth2.Token {
	token := g.api.GetToken(code)
	return token
}

func (g *Google) GetContext() context.Context {
	return g.api.Context
}
func (g *Google) GetClient() *http.Client {
	return g.api.Client
}

func (g *Google) Authentificate(token *oauth2.Token) {
	g.api.SetupClient(token)
}
