package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

type InvestorRepository interface {
	FindById(ctx *context.Context, id int) (*domain.Investor, *appErr.AppError)
	FindAll(ctx *context.Context) ([]domain.Investor, *appErr.AppError)
	FindByEmail(ctx *context.Context, email string) (*domain.Investor, *appErr.AppError)
	Save(ctx *context.Context, investor domain.Investor) *appErr.AppError
}

type InvestorRepositoryDb struct {
	db *pgxpool.Pool
}

func NewInvestorRepositoryDb(dbclient *pgxpool.Pool) InvestorRepositoryDb {
	return InvestorRepositoryDb{db: dbclient}
}

func (repo InvestorRepositoryDb) FindById(ctx *context.Context, id int) (*domain.Investor, *appErr.AppError) {
	var investor domain.Investor
	err := repo.db.QueryRow(*ctx, "SELECT * FROM get_investor_by_ID($1)", id).Scan(&investor.ID, &investor.FirstName, &investor.LastName, &investor.Balance, &investor.IsIssuer)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, appErr.NewNotFoundError("Investor not found ")
		}
		return nil, appErr.NewUnexpectedError("Error querying investor : " + err.Error())
	}

	return &investor, nil
}

func (repo InvestorRepositoryDb) FindAll(ctx *context.Context) ([]domain.Investor, *appErr.AppError) {
	investors := make([]domain.Investor, 0)
	rows, err := repo.db.Query(*ctx, "SELECT * FROM GET_INVESTORS()")
	if err != nil {
		return nil, appErr.NewUnexpectedError("Error querying investors : " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var investor domain.Investor
		err := rows.Scan(&investor.ID, &investor.FirstName, &investor.LastName, &investor.Balance, &investor.Status)
		if err != nil {
			return nil, appErr.NewUnexpectedError("Error scanning investors : " + err.Error())
		}

		investors = append(investors, investor)
	}

	return investors, nil
}

func (repo InvestorRepositoryDb) FindByEmail(ctx *context.Context, email string) (*domain.Investor, *appErr.AppError) {
	var user domain.Investor
	row := repo.db.QueryRow(*ctx, "SELECT * FROM get_investor_by_email($1)", email)

	err := row.Scan(&user.ID, &user.Password, &user.IsIssuer)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, appErr.NewNotFoundError("Email not found ")
		}
		return nil, appErr.NewUnexpectedError("Error scanning investor by email : " + err.Error())
	}

	return &user, nil
}

func (repo InvestorRepositoryDb) Save(ctx *context.Context, investor domain.Investor) *appErr.AppError {
	user, _ := repo.FindByEmail(ctx, investor.Email)

	if user != nil {
		return appErr.NewConflictError("Email already taken")
	}

	_, err := repo.db.Exec(*ctx, "call SAVE_INVESTOR($1,$2,$3,$4,$5,true,false)", investor.FirstName, investor.LastName, investor.Balance, investor.Email, investor.Password)
	if err != nil {
		return appErr.NewUnexpectedError("Error inserting invoice : " + err.Error())
	}
	return nil
}
