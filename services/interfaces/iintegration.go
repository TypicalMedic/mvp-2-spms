package interfaces

type IIntegration interface {
	GetAuthLink(redirectURI string) string
	Authentificate(token string)
}
