package services

import (
	"time"
	"context"
	"errors"

	"medods-test-task/internal/entities"
	"medods-test-task/internal/DAO"
	"medods-test-task/internal/repositories"
	"medods-test-task/internal/utils/hash"
	"medods-test-task/internal/utils/tokens"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersServiceInterface interface {
	SignUp(ctx context.Context, request DAO.SignUpRequest) error
	SignIn(ctx context.Context, request DAO.SignInRequest) (DAO.Tokens, error) 
	IdSignIn(ctx context.Context, id string) (DAO.Tokens, error) 
	RefreshTokens(ctx context.Context, refreshToken string) (DAO.Tokens, error) 
	CreateSession(ctx context.Context, userId primitive.ObjectID) (DAO.Tokens, error) 
}

type UsersService struct {
	repo         repositories.UsersRepo
	sessionRepo  repositories.SessionRepo
	hasher       hash.PasswordHasher
	tokenManager tokens.TokenManager

	accessTokenTTL         time.Duration
	refreshTokenTTL        time.Duration
}

func NewUsersService(repo repositories.UsersRepo, sessionRepo repositories.SessionRepo, hasher hash.PasswordHasher, tokenManager tokens.TokenManager,
	 accessTTL, refreshTTL time.Duration) *UsersService {
	return &UsersService{
		repo:                   repo,
		sessionRepo:            sessionRepo,
		hasher:                 hasher,
		tokenManager:           tokenManager,
		accessTokenTTL:         accessTTL,
		refreshTokenTTL:        refreshTTL,
	}
}

func (s *UsersService) SignUp(ctx context.Context, request DAO.SignUpRequest) error {
	passwordHash, err := s.hasher.Hash(request.Password)

	if err != nil {
		return err
	}

	if (s.repo.UserExistsByEmail(ctx, request.Email)) {
		return errors.New("you already have account")
	}

	user := entities.User{
		Password:     passwordHash,
		Email:        request.Email,
	}

	return s.repo.CreateUser(user)
}

func (s *UsersService) SignIn(ctx context.Context, request DAO.SignInRequest) (DAO.Tokens, error) {
	passwordHash, err := s.hasher.Hash(request.Password)

	if err != nil {
		return DAO.Tokens{}, err
	}

	user, err := s.repo.GetByCredentials(ctx, request.Email, passwordHash)

	if err != nil {
		return DAO.Tokens{}, err
	}

	return s.CreateSession(ctx, user.ID)
}

func (s *UsersService) IdSignIn(ctx context.Context, id string) (DAO.Tokens, error) {
	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return DAO.Tokens{}, err
	}

	user, err := s.repo.GetById(ctx, userId)

	if err != nil {
		return DAO.Tokens{}, err
	}

	return s.CreateSession(ctx, user.ID)
}

func (s *UsersService) RefreshTokens(ctx context.Context, refreshToken string) (DAO.Tokens, error) {
	refreshTokenHash, err := s.hasher.Hash(refreshToken)

	if err != nil {
		return DAO.Tokens{}, err
	}

	session, err := s.sessionRepo.GetByRefreshToken(ctx, refreshTokenHash)

	if err != nil {
		return DAO.Tokens{}, err
	}

	s.sessionRepo.DeleteSession(ctx, session)

	if time.Now() == session.ExpiresAt {
		return DAO.Tokens{}, errors.New("refresh token was expired")
	}

	return s.CreateSession(ctx, session.UserID)
}

func (s *UsersService) CreateSession(ctx context.Context, userId primitive.ObjectID) (DAO.Tokens, error) {
	var (
		res DAO.Tokens
		err error
	)

	res.AccessToken, err = s.tokenManager.NewJWT(userId.Hex(), s.accessTokenTTL)
	if err != nil {
		return  DAO.Tokens{}, err
	}

	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return  DAO.Tokens{}, err
	}

	refreshTokenHash, err := s.hasher.Hash(res.RefreshToken)
	if err != nil {
		return  DAO.Tokens{}, err
	}

	session := entities.Session{
		UserID: userId,
		RefreshToken: refreshTokenHash,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	}

	err = s.sessionRepo.CreateSession(ctx, session)
	if err != nil {
		return  DAO.Tokens{}, err
	}

	res.AccessTokenTTL = s.accessTokenTTL
	res.RefreshTokenTTL = s.refreshTokenTTL

	return res, err
}
