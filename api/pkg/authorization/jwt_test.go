package authorization_test

import (
	"testing"

	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/authorization"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestServices_JWTParseToken(t *testing.T) {
	t.Parallel()

	testTokens, err := authorization.GenerateTokensFromPayload(1)
	if err != nil {
		logrus.Fatalf("error occurred while generating tokens: %s", err.Error())

		return
	}

	tokenWithInvalidMethod, err1 := jwt.New(jwt.SigningMethodHS384).SignedString([]byte("28asd471"))
	if err1 != nil {
		logrus.Fatalf("error occurred while signing jwt token: %s", err1.Error())

		return
	}

	testCases := []struct {
		name      string
		token     string
		tokenType int
		err       error
	}{
		{
			name:      "invalid token type",
			token:     "token:)",
			tokenType: -3,
			err:       authorization.ErrInvalidTokenType,
		},
		{
			name:      "invalid token",
			token:     "token:)",
			tokenType: authorization.RefreshToken,
			err:       authorization.ErrInvalidToken,
		},
		{
			name:      "invalid token signing method",
			token:     tokenWithInvalidMethod,
			tokenType: authorization.RefreshToken,
			err:       authorization.ErrInvalidToken,
		},
		{
			name:      "valid access token",
			token:     testTokens.AccessToken,
			tokenType: authorization.AccessToken,
			err:       nil,
		},
		{
			name:      "valid refresh token",
			token:     testTokens.RefreshToken,
			tokenType: authorization.RefreshToken,
			err:       nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := authorization.ParseToken(tc.token, tc.tokenType)
			assert.Equal(t, tc.err, err)
		})
	}
}
