package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID       int32
	Username string
	Email    string
	Password string
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (int32, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *User) (int32, error) {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	var userID int32
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&userID)
	return userID, err
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	user := &User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (r *UserRepository) GenerateToken(ctx context.Context, userDto *usermd.UserDTO) (*usermd.SessionUserDTO, error) {
	user := token.User{
		UserID:    userDto.UserID,
		UserName:  userDto.UserName,
		Email:     userDto.Email,
		FcmDevice: "",
	}

	token, _, err := r.maker.CreateToken(user, r.config.AccessTokenDuration)
	if err != nil {
		return nil, err
	}

	refreshToken, payload, err := r.maker.CreateToken(user, r.config.RefreshTokenDuration)
	if err != nil {
		return nil, err
	}

	mtdt := metadata.ExtractMetadata(ctx)
	_, err = r.store.CreateSessionUser(ctx, usersqlc.CreateSessionUserParams{
		ID:           payload.ID,
		UserID:       user.UserID,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		FcmDevice:    user.FcmDevice,
		ExpiresAt:    payload.ExpiresAt,
	})

	if err != nil {
		return nil, err
	}

	return &usermd.SessionUserDTO{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil
}
