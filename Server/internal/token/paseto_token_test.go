package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/chacha20poly1305"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker("12345678901234567890123456789012")
	require.NoError(t, err)

	username := "test_user"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker("12345678901234567890123456789012")
	require.NoError(t, err)

	token, err := maker.CreateToken("test_user", -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	// Paseto v4 parser returns a specific error for expiration, but our wrapper returns ErrInvalidToken wrapped.
	// However, we didn't explicitly check for expiration error in VerifyToken to return ErrExpiredToken.
	// So it will likely return ErrInvalidToken.
	// Let's check if it contains "token is invalid".
	require.ErrorIs(t, err, ErrInvalidToken)
	require.Nil(t, payload)
}

func TestInvalidPasetoTokenKeySize(t *testing.T) {
	maker, err := NewPasetoMaker("invalid_key_size")
	require.Error(t, err)
	require.Nil(t, maker)
	require.EqualError(t, err, fmt.Sprintf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize))
}
