package token

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/docker/libtrust"
	"github.com/supereagle/go-example/rbac/auth-server/auth"
	"github.com/supereagle/go-example/rbac/auth-server/config"
)

const (
	issuer         = "Auth Server Demo"
	tokenSeparator = "."
)

func GenToken(ar *auth.AuthRequest, config *config.Config) (string, error) {
	pk, err := libtrust.LoadKeyFile(config.Token.Key)
	if err != nil {
		return "", fmt.Errorf("fail to load key file %s as %s", config.Token.Key, err.Error())
	}

	header := auth.Header{
		Type:       "JWT",
		SigningAlg: config.Token.Algorithm,
		KeyID:      pk.KeyID(),
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("unable to json marshal header as %s", err.Error())
	}

	jwtID, err := randString(16)
	if err != nil {
		return "", fmt.Errorf("fail to generate jwt id as %s", err.Error())
	}

	now := time.Now().Unix()
	claimSet := auth.ClaimSet{
		Issuer:     issuer,
		Subject:    ar.Username,
		Audience:   string(ar.Service),
		NotBefore:  now - 1,
		IssuedAt:   now,
		Expiration: now + config.Token.Expiration,
		JWTID:      jwtID,
		Access:     []*auth.ResourceActions{},
	}

	if len(ar.Scope.Actions) > 0 {
		claimSet.Access = []*auth.ResourceActions{
			&auth.ResourceActions{Type: ar.Scope.ResourceType, Name: ar.Scope.ResourceName, Actions: ar.Scope.Actions},
		}
	}

	claimsJSON, err := json.Marshal(claimSet)
	if err != nil {
		return "", fmt.Errorf("unable to json marshal claimSet as %s", err.Error())
	}

	payload := fmt.Sprintf("%s%s%s", joseBase64UrlEncode(headerJSON), tokenSeparator, joseBase64UrlEncode(claimsJSON))

	sig, _, err := pk.Sign(strings.NewReader(payload), 0)
	if err != nil {
		return "", fmt.Errorf("unable to sign jwt payload as %s", err.Error())
	}

	return fmt.Sprintf("%s%s%s", payload, tokenSeparator, joseBase64UrlEncode(sig)), nil
}

// Copy from github.com/docker/libtrust/util.go.
func joseBase64UrlEncode(b []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}

// Copy from github.com/vmware/harbor/src/ui/service/token/authutils.go
func randString(length int) (string, error) {
	const alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rb := make([]byte, length)
	_, err := rand.Read(rb)
	if err != nil {
		return "", err
	}
	for i, b := range rb {
		rb[i] = alphanum[int(b)%len(alphanum)]
	}
	return string(rb), nil
}
