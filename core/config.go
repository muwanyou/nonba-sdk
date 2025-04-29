package core

import (
	"net/http"
)

type Config struct {
	HTTPClient *http.Client
	Credential *Credential
	Signer     Signer
}
