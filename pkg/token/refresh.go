package token

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type RefreshTokenManager interface {
	GenerateSecret() (string, error)
	Format(id uuid.UUID, secret string) string
	Parse(tokenStr string) (id uuid.UUID, secret string, err error)
}

type opaqueManager struct {
	length int
}

func NewRefreshTokenManager(length int) RefreshTokenManager {
	return &opaqueManager{
		length: length,
	}
}

func (m *opaqueManager) GenerateSecret() (string, error) {
	b := make([]byte, m.length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	secret := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)

	return secret, nil
}

func (m *opaqueManager) Format(id uuid.UUID, secret string) string {
	return fmt.Sprintf("rt_%s_%s", id.String(), secret)
}

func (m *opaqueManager) Parse(tokenStr string) (id uuid.UUID, secret string, err error) {
	parts := strings.Split(tokenStr, "_")
	if len(parts) != 3 || parts[0] != "rt" {
		return uuid.Nil, "", fmt.Errorf("formato de token invalido")
	}

	parsedID, err := uuid.Parse(parts[1])
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("id do token corrompido: %w", err)
	}

	return parsedID, parts[2], nil
}
