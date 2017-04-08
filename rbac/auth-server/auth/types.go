package auth

import (
	"encoding/json"
	"time"
)

type Service string

const (
	Product Service = "product"
	Order           = "order"
)

type AuthRequest struct {
	Username string
	Password string
	Service  Service
	Scope    *Scope
}

type Scope struct {
	ResourceType string
	ResourceName string
	Actions      []string
}

type TokenResponse struct {
	Token     string
	ExpiresIn int
	issuedAt  *time.Time
}

// Following structs are copied from github.com/docker/distribution/registry/auth/token/token.go.

// ResourceActions stores allowed actions on a named and typed resource.
type ResourceActions struct {
	Type    string   `json:"type"`
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}

// ClaimSet describes the main section of a JSON Web Token.
type ClaimSet struct {
	// Public claims
	Issuer     string `json:"iss"`
	Subject    string `json:"sub"`
	Audience   string `json:"aud"`
	Expiration int64  `json:"exp"`
	NotBefore  int64  `json:"nbf"`
	IssuedAt   int64  `json:"iat"`
	JWTID      string `json:"jti"`

	// Private claims
	Access []*ResourceActions `json:"access"`
}

// Header describes the header section of a JSON Web Token.
type Header struct {
	Type       string          `json:"typ"`
	SigningAlg string          `json:"alg"`
	KeyID      string          `json:"kid,omitempty"`
	X5c        []string        `json:"x5c,omitempty"`
	RawJWK     json.RawMessage `json:"jwk,omitempty"`
}

// Token describes a JSON Web Token.
type Token struct {
	Raw       string
	Header    *Header
	Claims    *ClaimSet
	Signature []byte
}
