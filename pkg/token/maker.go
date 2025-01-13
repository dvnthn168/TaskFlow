package token

import (
	"time"

	"github.com/google/uuid"
)

type Maker interface {
	CreateToken(user User, duration time.Duration) (string, *UserPayload, error)
	VerifyToken(token string) (*UserPayload, error)
	RefreshToken(id uuid.UUID, user User, duration time.Duration) (string, *UserPayload, error)
}
