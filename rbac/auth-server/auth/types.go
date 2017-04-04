package auth

import "time"

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
