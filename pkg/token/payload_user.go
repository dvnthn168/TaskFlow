package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    int32  `json:"user_id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	FcmDevice string `json:"fcm_device"`
}

type UserPayload struct {
	ID        uuid.UUID `json:"id"`
	User      User      `json:"user"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expired_at"`
}

func NewPayload(user User, duration time.Duration) (*UserPayload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &UserPayload{
		ID:        tokenID,
		User:      user,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
	return payload, nil
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

func RePayloadBarber(id uuid.UUID, user User, duration time.Duration) (*UserPayload, error) {
	payload := &UserPayload{
		ID:        id,
		User:      user,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *UserPayload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}
