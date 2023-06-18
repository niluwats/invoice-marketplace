package service

import (
	"errors"

	"github.com/niluwats/invoice-marketplace/internal/auth"
	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyUser(req dto.AuthRequest) (*dto.AuthResponse, error)
	Register(req dto.NewUserRequest) error
}

type DefaultAuthService struct {
	repo repositories.InvestorRepository
}

func NewAuthService(repo repositories.InvestorRepository) DefaultAuthService {
	return DefaultAuthService{repo}
}

func (s DefaultAuthService) VerifyUser(req dto.AuthRequest) (*dto.AuthResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = user.CheckPassword(req.Password)
	if err != nil {
		return nil, errors.New("Incorrect password")
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.ID, user.IsIssuer)
	if err != nil {
		return nil, err
	}

	return dto.GetAuthResponse(user.ID, tokenString), nil
}

func (s DefaultAuthService) Register(req dto.NewUserRequest) error {
	if !req.IsValidPassword() {
		return errors.New("Invalid password")
	}
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	investor := domain.Investor{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Balance:   req.Balance,
		Email:     req.Email,
		Password:  string(hashedPw),
	}
	err = s.repo.Save(investor)
	if err != nil {
		return err
	}
	return nil
}
