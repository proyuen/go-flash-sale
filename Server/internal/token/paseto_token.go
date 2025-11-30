package token

import (
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	symmetricKey paseto.V4SymmetricKey
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	key, err := paseto.V4SymmetricKeyFromBytes([]byte(symmetricKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create symmetric key: %w", err)
	}

	maker := &PasetoMaker{
		symmetricKey: key,
	}
	return maker, nil
}

// CreateToken 生成token
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	token := paseto.NewToken()
	token.Set("id", payload.ID.String())
	token.Set("username", payload.Username)
	token.SetIssuedAt(payload.IssuedAt)
	token.SetExpiration(payload.ExpiredAt)

	return token.V4Encrypt(maker.symmetricKey, nil), nil
}

func (maker *PasetoMaker) VerifyToken(tokenString string) (*Payload, error) {
	parser := paseto.NewParser()
	token, err := parser.ParseV4Local(maker.symmetricKey, tokenString, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	idStr, err := token.GetString("id")
	if err != nil {
		return nil, ErrInvalidToken
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, ErrInvalidToken
	}

	username, err := token.GetString("username")
	if err != nil {
		return nil, ErrInvalidToken
	}

	issuedAt, err := token.GetIssuedAt()
	if err != nil {
		return nil, ErrInvalidToken
	}

	expiredAt, err := token.GetExpiration()
	if err != nil {
		return nil, ErrInvalidToken
	}

	payload := &Payload{
		ID:        id,
		Username:  username,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
