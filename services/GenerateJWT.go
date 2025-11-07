package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenarateJwt(U_id string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Minute)
	
	claims := &UserClaims{
		UserID: U_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "todo-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := os.Getenv("JWT_SECRET")

	// fmt.Println(jwtSecret)

	if jwtSecret == "" {
		return "", fmt.Errorf("forgot to put the jwt key on env ")
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", fmt.Errorf("failed to stringify the tocken ")
	}

	return tokenString, nil
}
