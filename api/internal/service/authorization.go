package service

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/cache"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/config"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/pkg/authorization"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type AuthorizationImpl struct {
	userRepository    repository.User
	refreshTokenCache cache.RefreshToken
	blacklistCache    cache.Blacklist

	logger *logging.Logger
}

func NewAuthorizationImpl(
	userRepository repository.User,
	refreshTokenCache cache.RefreshToken,
	blacklistCache cache.Blacklist,
) *AuthorizationImpl {
	return &AuthorizationImpl{
		userRepository:    userRepository,
		refreshTokenCache: refreshTokenCache,
		blacklistCache:    blacklistCache,

		logger: logging.NewLogger("authorization service"),
	}
}

func (s *AuthorizationImpl) SignUpUser(userData dto.SigningUp) (int, *Error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, NewInternalServerError(err)
	}

	userData.Password = string(hashedPassword)
	id, err := s.userRepository.Create(rdto.UserCreating(userData))
	if err != nil {
		s.logger.Error(err.Error())

		return 0, UserAlreadyExistsError
	}

	return id, nil
}

func (s *AuthorizationImpl) GenerateTokens(userCredentials dto.UserCredentials) (*dto.AuthorizationTokens, *Error) {
	user, err := s.userRepository.GetUserByUsername(userCredentials.Username)
	if err != nil {
		return nil, InvalidUserNameOrPasswordError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredentials.Password)); err != nil {
		return nil, InvalidUserNameOrPasswordError
	}

	if err := s.blacklistCache.Delete(user.ID); err != nil {
		return nil, NewInternalServerError(err)
	}

	tokens, err := authorization.GenerateTokensFromPayload(user.ID)
	if err != nil {
		return nil, NewInternalServerError(err)
	}

	if err := s.refreshTokenCache.Save(user.ID, tokens.RefreshToken, config.Get().Auth.RefreshToken.TTL); err != nil {
		return nil, NewInternalServerError(err)
	}

	return tokens, nil
}

func (s *AuthorizationImpl) RefreshTokens(refreshToken string) (*dto.AuthorizationTokens, *Error) {
	tokenClaims, err := authorization.ParseToken(refreshToken, authorization.RefreshToken)
	if err != nil {
		return nil, InvalidTokenError
	}

	savedRefreshToken, err := s.refreshTokenCache.Get(tokenClaims.UserID)
	if err != nil {
		return nil, InvalidTokenError
	}

	if strings.Compare(savedRefreshToken, refreshToken) != 0 {
		return nil, InvalidTokenError
	}

	newlyGeneratedTokens, err := authorization.GenerateTokensFromPayload(tokenClaims.UserID)
	if err != nil {
		return nil, NewInternalServerError(err)
	}

	if err := s.refreshTokenCache.Save(
		tokenClaims.UserID, newlyGeneratedTokens.RefreshToken, config.Get().Auth.RefreshToken.TTL,
	); err != nil {
		return nil, NewInternalServerError(err)
	}

	return newlyGeneratedTokens, nil
}

func (s *AuthorizationImpl) LogoutUser(userID int, accessToken string) *Error {
	tokenClaims, err := authorization.ParseToken(accessToken, authorization.AccessToken)
	if err != nil {
		return InvalidTokenError
	}

	storageTimeInTheBlacklist := time.Until(time.Unix(tokenClaims.ExpiresAt, 0))
	if err := s.blacklistCache.Save(userID, accessToken, storageTimeInTheBlacklist); err != nil {
		return NewInternalServerError(err)
	}

	if err := s.refreshTokenCache.Delete(userID); err != nil {
		return NewInternalServerError(err)
	}

	return nil
}

func (s *AuthorizationImpl) IsUserLogout(userID int) bool {
	if err := s.blacklistCache.Get(userID); err != nil {
		return true
	}

	return false
}
