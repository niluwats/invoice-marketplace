package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

type IssuerRepository interface {
	FindById(id int) (*domain.Issuer, *appErr.AppError)
	FindAll() ([]domain.Issuer, *appErr.AppError)
}

type IssuerRepositoryDb struct {
	db *pgxpool.Pool
}

func NewIssuerRepositoryDb(dbclient *pgxpool.Pool) IssuerRepositoryDb {
	return IssuerRepositoryDb{db: dbclient}
}

func (repo IssuerRepositoryDb) FindById(id int) (*domain.Issuer, *appErr.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT issuers.id,company_name,balance FROM issuers INNER JOIN investors ON issuers.investor_id=investors.id WHERE issuers.id=$1"

	var issuer domain.Issuer
	row := repo.db.QueryRow(ctx, query, id)

	err := row.Scan(&issuer.ID, &issuer.CompanyName, &issuer.Investor.Balance)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, appErr.NewNotFoundError("Issuer not found ")
		}
		return nil, appErr.NewUnexpectedError("Error querying issuer : " + err.Error())
	}

	return &issuer, nil
}

func (repo IssuerRepositoryDb) FindAll() ([]domain.Issuer, *appErr.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT issuers.id,company_name,balance FROM issuers INNER JOIN investors ON issuers.investor_id=investors.id"

	issuers := make([]domain.Issuer, 0)
	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, appErr.NewUnexpectedError("Error querying issuers : " + err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var issuer domain.Issuer
		err := rows.Scan(&issuer.ID, &issuer.CompanyName, &issuer.Investor.Balance)
		if err != nil {
			return nil, appErr.NewUnexpectedError("Error scanning issuers : " + err.Error())
		}

		issuers = append(issuers, issuer)
	}

	return issuers, nil
}
