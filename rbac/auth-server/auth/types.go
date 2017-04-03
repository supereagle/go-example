package auth

type Service string

const (
	Product Service = "product"
	Order           = "order"
)

type AuthRequest struct {
	Username     string
	Password     string
	ResourceType string
	ResourceName string
	Actions      []string
}
