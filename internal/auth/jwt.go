package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("jwt-secret000")

type JWTClaim struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func ExtractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	return splitToken[1]
}

func GenerateJWT(email, id string, isIssuer bool) (string, error) {
	expiration := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	if isIssuer {
		claims.Role = "issuer"
	} else {
		claims.Role = "investor"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseClaims(signedToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (any, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractClaims(token *jwt.Token) (*JWTClaim, error) {
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err := errors.New("couldn't parse claims")
		return nil, err
	}
	return claims, nil
}

func ValidateTokenAndGetClaims(signedToken string) (id string, role string, err error) {
	token, err := ParseClaims(signedToken)
	if err != nil {
		return "", "", err
	}

	claims, err := ExtractClaims(token)
	if err != nil {
		return "", "", err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return "", "", err
	}
	return claims.Id, claims.Role, nil
}
