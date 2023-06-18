package dto

import (
	"fmt"
	"unicode"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
}

type NewUserRequest struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Balance   float64 `json:"balance"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
}

func (req *AuthRequest) IfInValidRequest() bool {
	return req.Email == "" || req.Password == ""
}

func (req *NewUserRequest) IfInValidRequest() bool {
	return (req.FirstName == "" || req.LastName == "" || req.Email == "" || req.Password == "" || fmt.Sprintf("%f", req.Balance) == "")
}

func GetAuthResponse(id, token string) *AuthResponse {
	return &AuthResponse{
		ID:          id,
		AccessToken: token,
	}
}

func (req NewUserRequest) IsValidPassword() bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(req.Password) >= 7 {
		hasMinLen = true
	}
	for _, char := range req.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
