package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/muwanyou/nonba-sdk/util"
)

type Signer interface {
	Type() string
	Sign() string
}

type HMACSigner struct {
	credential *Credential
}

func NewHMACSigner(credential *Credential) *HMACSigner {
	return &HMACSigner{credential: credential}
}

func (s *HMACSigner) Type() string {
	return "HMAC-SHA256"
}

func (s *HMACSigner) Sign() string {
	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)
	ciphertext := util.GenerateHMAC(sha256.New, []byte(s.credential.SecretKey), []byte(s.credential.OpenID+s.credential.SecretKey+timestamp))
	signature := fmt.Sprintf("open_id=%s,app_id=%s,timestamp=%s,signature=%s", s.credential.OpenID, s.credential.AppID, timestamp, hex.EncodeToString(ciphertext))
	return string(signature)
}
