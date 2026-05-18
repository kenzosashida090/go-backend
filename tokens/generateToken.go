package tokens

import (
	"errors"
	"fmt"
	"os"
	"time"

	"db-go.com/api/models"
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	Email   string `json:"email"`
	User_ID int64  `json:"user_id"`
	// Token string `json:"token"`
	jwt.RegisteredClaims
}

var signRefreshKey = []byte(os.Getenv("JWT_REFRESH"))
var signAccessKey = []byte(os.Getenv("JWT_KEY"))

func GenerateRefreshToken(email string, userId int64) (string, time.Time, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(24 * time.Hour))

	claims := TokenClaims{
		email,
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(signRefreshKey)
	if err != nil {
		return "", time.Time{}, errors.New("Something bad with toikens")
	}
	return ss, expiresAt.Time, nil
}

func GenerateAccesToken(user *models.User) (string, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(30 * time.Minute))
	claims := TokenClaims{
		user.Email,
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signAccessKey)
	if err != nil {
		return "", errors.New("Something went wrong signing token")
	}
	return ss, nil
}

func SaveRefreshTokenDB(user *models.User) (string, error) {
	ss, expiresAt, err := GenerateRefreshToken(user.Email, user.ID)
	err = models.Save(user.ID, expiresAt, ss)
	if err != nil {
		return "", errors.New("Something bad with toikens")
	}
	return ss, nil
}

func ValidateToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (any, error) {
		return signAccessKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Something went wrong validating singning method")
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, errors.New("Something went wrong obtening the claims.")
	}
	return claims, nil
}
