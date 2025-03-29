package jwt

import (
	"Go-Starter-Template/internal/utils"
	"Go-Starter-Template/pkg/entities/domain"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
)

type (
	JWTService interface {
		GenerateTokenUser(userId string, role string) string
		ValidateTokenUser(token string) (*jwt.Token, error)
		GetUserIDByToken(token string) (string, string, error)
	}
	jwtEmailClaim struct {
		Email string `json:"email"`
		jwt.RegisteredClaims
	}

	jwtUserClaim struct {
		UserID string `json:"user_id"`
		Role   string `json:"role"`
		jwt.RegisteredClaims
	}

	jwtService struct {
		secretKey string
		issuer    string
	}
)

func getSecretKey() string {
	utils.LoadEnv()
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if os.Getenv("JWT_SECRET") != "" {
		secretKey = "pujodarmawan"
	}
	return secretKey
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "FP SWE KELOMPOK 3",
	}
}

func (j *jwtService) GenerateTokenUser(userId string, role string) string {
	claims := jwtUserClaim{
		userId,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 120)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tx, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}
	return tx
}

func (j *jwtService) parseToken(t_ *jwt.Token) (any, error) {
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
	}
	return []byte(j.secretKey), nil
}

func (j *jwtService) ValidateTokenUser(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &jwtUserClaim{}, j.parseToken)
}

func (j *jwtService) GetUserIDByToken(token string) (string, string, error) {
	t_Token, err := j.ValidateTokenUser(token)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", "", domain.ErrTokenExpired
		}
		return "", "", domain.ErrTokenInvalid
	}
	if !t_Token.Valid {
		return "", "", domain.ErrTokenInvalid
	}

	claims := t_Token.Claims.(*jwtUserClaim)

	id := fmt.Sprintf("%v", claims.UserID)
	role := fmt.Sprintf("%v", claims.Role)
	return id, role, nil
}
