package core

type Credential struct {
	OpenID    string
	AppID     string
	SecretKey string
}

func NewCredential(openID string, appID string, secretKey string) *Credential {
	return &Credential{
		OpenID:    openID,
		AppID:     appID,
		SecretKey: secretKey,
	}
}
