package auth

import (
	"time"

	"distrodakwah_backend/config"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type CredUser struct {
	Email string
}

type LoginCredetials struct {
	Email    string
	Password string
}

// Hash make a password hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword verify the hashed password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateJWT(user *CredUser) (string, error) {
	claim := Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "distrodakwah.id",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(config.JWTSECRET)
}
