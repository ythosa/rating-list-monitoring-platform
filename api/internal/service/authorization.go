package service

import (
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/cache"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/logging"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/internal/repository/rdto"
	"golang.org/x/crypto/bcrypt"
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

	tokens, err := auth
}

func (s *AuthorizationImpl) RefreshTokens(refreshToken string) (*dto.AuthorizationTokens, *Error) {
	panic("implement me")
}

func (s *AuthorizationImpl) LogoutUser(userID int, accessToken string) *Error {
	panic("implement me")
}

func (s *AuthorizationImpl) IsUserLogout(userID int) bool {
	panic("implement me")
}
