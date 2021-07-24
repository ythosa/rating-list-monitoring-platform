package service

import "github.com/ythosa/rating-list-monitoring-platfrom-api/internal/dto"

type Authorization interface {
	SignUpUser(userData dto.SigningUp) (int, *Error)
	GenerateTokens(userCredentials dto.UserCredentials) (*dto.AuthorizationTokens, *Error)
	RefreshTokens(refreshToken string) (*dto.AuthorizationTokens, *Error)
	LogoutUser(userID int, accessToken string) *Error
	IsUserLogout(userID int) bool
}
