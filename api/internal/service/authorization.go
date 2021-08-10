package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/config"
	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/logging"

	"golang.org/x/crypto/bcrypt"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/cache"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/repository/rdto"
	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/authorization"
)

type AuthorizationImpl struct {
	userRepository    repository.User
	refreshTokenCache cache.RefreshToken
	blacklistCache    cache.Blacklist
	logger            *logging.Logger
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
		logger:            logging.NewLogger("authorization service"),
	}
}

func (s *AuthorizationImpl) SignUpUser(userData dto.SigningUp) (uint, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("error while crypting password: %w", err)
	}

	userData.Password = string(hashedPassword)

	id, err := s.userRepository.Create(rdto.UserCreating(userData))
	if err != nil {
		s.logger.Error(err)

		return 0, UserAlreadyExistsError
	}

	return id, nil
}

func (s *AuthorizationImpl) GenerateTokens(userCredentials dto.UserCredentials) (*dto.AuthorizationTokens, error) {
	user, err := s.userRepository.GetUserByUsername(userCredentials.Username)
	if err != nil {
		return nil, InvalidUsernameOrPasswordError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredentials.Password)); err != nil {
		return nil, InvalidUsernameOrPasswordError
	}

	if err := s.blacklistCache.Delete(user.ID); err != nil {
		return nil, fmt.Errorf("error while deleting user from cache blacklist: %w", err)
	}

	latestRefreshToken, errGettingLatestRefreshToken := s.refreshTokenCache.Get(user.ID)
	_, errParsingLatestRefreshToken := authorization.ParseToken(
		latestRefreshToken, config.Get().AuthTokens.RefreshToken,
	)

	if errGettingLatestRefreshToken != nil || errParsingLatestRefreshToken != nil {
		return s.generateTokensAndSaveRefreshToken(user.ID)
	}

	accessToken, err := authorization.GenerateTokenFromPayload(user.ID, config.Get().AuthTokens.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("error while generating token: %w", err)
	}

	return &dto.AuthorizationTokens{
		RefreshToken: latestRefreshToken,
		AccessToken:  accessToken,
	}, nil
}

func (s *AuthorizationImpl) RefreshTokens(refreshToken string) (*dto.AuthorizationTokens, error) {
	tokenClaims, err := authorization.ParseToken(refreshToken, config.Get().AuthTokens.RefreshToken)
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

	return s.generateTokensAndSaveRefreshToken(tokenClaims.UserID)
}

func (s *AuthorizationImpl) generateTokensAndSaveRefreshToken(userID uint) (*dto.AuthorizationTokens, error) {
	tokens, err := authorization.GenerateTokensFromPayload(userID, config.Get().AuthTokens)
	if err != nil {
		return nil, fmt.Errorf("error while generating tokens: %w", err)
	}

	if err := s.refreshTokenCache.Save(
		userID, tokens.RefreshToken, config.Get().AuthTokens.RefreshToken.TTL,
	); err != nil {
		return nil, fmt.Errorf("error while saving refresh token in cache: %w", err)
	}

	return (*dto.AuthorizationTokens)(tokens), nil
}

func (s *AuthorizationImpl) LogoutUser(userID uint, accessToken string) error {
	tokenClaims, err := authorization.ParseToken(accessToken, config.Get().AuthTokens.AccessToken)
	if err != nil {
		return InvalidTokenError
	}

	storageTimeInTheBlacklist := time.Until(time.Unix(tokenClaims.ExpiresAt, 0))
	if err := s.blacklistCache.Save(userID, accessToken, storageTimeInTheBlacklist); err != nil {
		return fmt.Errorf("error while adding user to cahce blacklist: %w", err)
	}

	if err := s.refreshTokenCache.Delete(userID); err != nil {
		return fmt.Errorf("error while deleting user refresh token from cache: %w", err)
	}

	return nil
}

func (s *AuthorizationImpl) IsUserLogout(userID uint) bool {
	return s.blacklistCache.Get(userID) != nil
}
