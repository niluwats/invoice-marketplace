package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

type IssuerRepository interface {
	FindById(ctx *context.Context, id int) (*domain.Issuer, *appErr.AppError)
	FindAll(ctx *context.Context) ([]domain.Issuer, *appErr.AppError)
}

type IssuerRepositoryDb struct {
	db *pgxpool.Pool
}

func NewIssuerRepositoryDb(dbclient *pgxpool.Pool) IssuerRepositoryDb {
	return IssuerRepositoryDb{db: dbclient}
}

func (repo IssuerRepositoryDb) FindById(ctx *context.Context, id int) (*domain.Issuer, *appErr.AppError) {
	var issuer domain.Issuer
	row := repo.db.QueryRow(*ctx, "SELECT * FROM GET_ISSUER_BY_ID($1)", id)

	err := row.Scan(&issuer.ID, &issuer.CompanyName, &issuer.Investor.Balance)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, appErr.NewNotFoundError("Issuer not found ")
		}
		return nil, appErr.NewUnexpectedError("Error querying issuer : " + err.Error())
	}

	return &issuer, nil
}

func (repo IssuerRepositoryDb) FindAll(ctx *context.Context) ([]domain.Issuer, *appErr.AppError) {
	issuers := make([]domain.Issuer, 0)
	rows, err := repo.db.Query(*ctx, "SELECT * FROM GET_ISSUERS()")
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
