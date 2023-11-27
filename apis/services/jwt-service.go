package services

import (
	"fmt"
	"os"

	"time"

	"github.com/aniket-skroman/skroman_sales_service.git/apis/helper"
	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userID, userType string) string
	GenerateTempToken(UserID, userType, dept string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID    string    `json:"user_id"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	jwt.StandardClaims
}

type jwtTempCustomClaim struct {
	UserID         string    `json:"user_id"`
	UserType       string    `json:"user_type"`
	UserDepartment string    `json:"dept"`
	CreatedAt      time.Time `json:"created_at"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "ydhnwb",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretkey := os.Getenv("JWT_SECRET")

	if secretkey != "" {
		secretkey = "ydhnwb"
	}
	return secretkey
}

func (j *jwtService) GenerateToken(UserID string, userType string) string {
	claims := &jwtCustomClaim{
		UserID,
		userType,
		time.Now().Add(10 * time.Minute),
		jwt.StandardClaims{
			ExpiresAt: int64(5 * time.Minute),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) GenerateTempToken(UserID, userType, dept string) string {
	UserID, _ = helper.EncryptData(UserID)
	userType, _ = helper.EncryptData(userType)
	dept, _ = helper.EncryptData(dept)

	claims := &jwtTempCustomClaim{
		UserID,
		userType,
		dept,
		time.Now().Add(10 * time.Minute),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
