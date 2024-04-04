package interfaces

import "golang.org/x/oauth2"

type IIntegration interface {
	GetAuthLink(redirectURI string, accountId int, returnURL string) string
	Authentificate(token oauth2.Token)
	GetToken(code string) oauth2.Token
}
