package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type RefreshTokenData struct {
	Token     string
	ExpiresAt time.Time
}

func GenerateRefreshToken() (*RefreshTokenData, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	token := hex.EncodeToString(bytes)

	expiry := time.Now().Add(10 * time.Hour) // 7 days

	return &RefreshTokenData{
		Token:     token,
		ExpiresAt: expiry,
	}, nil

}

func GenerateHash(tocken string) string {
	hash := sha256.Sum256([]byte(tocken))
	return hex.EncodeToString(hash[:])
}
