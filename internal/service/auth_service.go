package service

import (
	"context"

	"github.com/niluwats/invoice-marketplace/internal/auth"
	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyUser(ctx context.Context, req dto.AuthRequest) (*dto.AuthResponse, *appErr.AppError)
	Register(ctx context.Context, req dto.NewUserRequest) *appErr.AppError
}

type DefaultAuthService struct {
	repo repositories.InvestorRepository
}

func NewAuthService(repo repositories.InvestorRepository) DefaultAuthService {
	return DefaultAuthService{repo}
}

func (s DefaultAuthService) VerifyUser(ctx context.Context, req dto.AuthRequest) (*dto.AuthResponse, *appErr.AppError) {
	if req.IfInValidRequest() {
		return nil, appErr.NewValidationError("All fields required")
	}

	user, err_ := s.repo.FindByEmail(&ctx, req.Email)
	if err_ != nil {
		return nil, err_
	}

	ok := user.CheckPassword(req.Password)
	if !ok {
		return nil, appErr.NewAuthenticationError("Incorrect password")
	}

	tokenString, errJwt := auth.GenerateJWT(user.Email, user.ID, user.IsIssuer)
	if errJwt != nil {
		return nil, appErr.NewUnexpectedError("Error generating JWT : " + errJwt.Error())
	}

	return dto.GetAuthResponse(user.ID, tokenString), nil
}

func (s DefaultAuthService) Register(ctx context.Context, req dto.NewUserRequest) *appErr.AppError {
	if req.IfInValidRequest() {
		return appErr.NewValidationError("All fields required")
	}

	if !req.IsValidPassword() {
		return appErr.NewValidationError("Invalid password")
	}
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return appErr.NewUnexpectedError("Error hashing password : " + err.Error())
	}

	investor := domain.Investor{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Balance:   req.Balance,
		Email:     req.Email,
		Password:  string(hashedPw),
	}
	err_ := s.repo.Save(&ctx, investor)
	if err_ != nil {
		return err_
	}
	return nil
}
