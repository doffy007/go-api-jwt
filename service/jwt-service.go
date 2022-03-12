package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Contract of JWT can do
type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer string
}

//NewJWTService is a method create new instace of service JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer: "doffy",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "doffy"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 45).Unix(),
			Issuer: j.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j * jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_*jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected sign method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}