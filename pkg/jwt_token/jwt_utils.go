package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type Claims struct {
	Email  string `json:"username"`
	UserId uint   `json:"userid"`
	jwt.StandardClaims
}

func GetJWTToken(email string, userId uint, hour time.Duration) (string, error) {
	claims := Claims{
		Email:  email,
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * hour).Unix(),
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("failed to create token")
	}
	return signedToken, nil
}

func VerifyToken(jwtToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrInvalidToken
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, ErrInvalidToken
			} else {
				return nil, err
			}
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrInvalidToken
}
