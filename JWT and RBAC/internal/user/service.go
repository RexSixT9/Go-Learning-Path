package user

import (
	"context"
	"errors"
	"go-auth/internal/auth"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repo

	jwtSecret string
}

func NewService(repo *Repo, jwtSecret string) *Service {
	return &Service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResult struct {
	Token string     `json:"token"`
	User  PublicUser `json:"user"`
}

func (s *Service) createAuthResult(user User) (AuthResult, error) {
	token, err := auth.CreateToken(
		s.jwtSecret,
		user.ID.Hex(),
		user.Role,
	)
	if err != nil {
		return AuthResult{}, err
	}

	return AuthResult{
		Token: token,
		User:  ToPublicUser(user),
	}, nil
}

func (s *Service) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {

	email := strings.ToLower(strings.TrimSpace(input.Email))
	pass := strings.ToLower(strings.TrimSpace(input.Password))

	if email == "" || pass == "" {
		return AuthResult{}, errors.New("email and password are required")
	}

	if len(pass) < 8 {
		return AuthResult{}, errors.New("password must be at least 8 characters long")
	}

	_, err := s.repo.FindByEmail(ctx, email)
	if err == nil {
		return AuthResult{}, errors.New("user with this email already exists")
	}

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return AuthResult{}, errors.New("failed to check existing user: " + err.Error())
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return AuthResult{}, errors.New("failed to hash password")
	}

	now := time.Now().UTC()

	user := User{
		Email:        email,
		PasswordHash: string(hashBytes),
		Role:         "user",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return AuthResult{}, errors.New("failed to create user: " + err.Error())
	}

	return s.createAuthResult(createdUser)
}

func (s *Service) Login(ctx context.Context, input LoginInput) (AuthResult, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.ToLower(strings.TrimSpace(input.Password))

	if email == "" || password == "" {
		return AuthResult{}, errors.New("email and password are required")
	}

	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return AuthResult{}, errors.New("invalid email or password")
		}
		return AuthResult{}, errors.New("failed to find user: " + err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return AuthResult{}, errors.New("invalid email or password")
	}

	return s.createAuthResult(user)
}
