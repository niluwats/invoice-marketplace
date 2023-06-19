package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

type InvestorRepository interface {
	FindById(id int) (*domain.Investor, *appErr.AppError)
	FindAll() ([]domain.Investor, *appErr.AppError)
	FindByEmail(email string) (*domain.Investor, *appErr.AppError)
	Save(investor domain.Investor) *appErr.AppError
}

type InvestorRepositoryDb struct {
	db *pgxpool.Pool
}

func NewInvestorRepositoryDb(dbclient *pgxpool.Pool) InvestorRepositoryDb {
	return InvestorRepositoryDb{db: dbclient}
}

func (repo InvestorRepositoryDb) FindById(id int) (*domain.Investor, *appErr.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT id,first_name,last_name,balance,is_issuer FROM investors WHERE ID=$1"

	var investor domain.Investor
	row := repo.db.QueryRow(ctx, query, id)

	err := row.Scan(&investor.ID, &investor.FirstName, &investor.LastName, &investor.Balance, &investor.IsIssuer)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, appErr.NewNotFoundError("Investor not found ")
		}
		return nil, appErr.NewUnexpectedError("Error querying investor : " + err.Error())
	}

	return &investor, nil
}

func (repo InvestorRepositoryDb) FindAll() ([]domain.Investor, *appErr.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT id,first_name,last_name,balance FROM investors"

	investors := make([]domain.Investor, 0)
	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, appErr.NewUnexpectedError("Error querying investors : " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var investor domain.Investor
		err := rows.Scan(&investor.ID, &investor.FirstName, &investor.LastName, &investor.Balance)
		if err != nil {
			return nil, appErr.NewUnexpectedError("Error scanning investors : " + err.Error())
		}

		investors = append(investors, investor)
	}

	return investors, nil
}

func (repo InvestorRepositoryDb) FindByEmail(email string) (*domain.Investor, *appErr.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT id,password,is_issuer FROM investors WHERE email=$1"

	var user domain.Investor
	row := repo.db.QueryRow(ctx, query, email)

	err := row.Scan(&user.ID, &user.Password, &user.IsIssuer)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, appErr.NewNotFoundError("Email not found ")
		}
		return nil, appErr.NewUnexpectedError("Error scanning investor by email : " + err.Error())
	}

	return &user, nil
}

func (repo InvestorRepositoryDb) Save(investor domain.Investor) *appErr.AppError {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	user, _ := repo.FindByEmail(investor.Email)

	if user != nil {
		return appErr.NewConflictError("Email already taken")
	}

	query := `INSERT INTO investors(first_name,last_name,balance,email,password,is_active,is_issuer) 
				VALUES($1,$2,$3,$4,$5,true,false)`

	_, err := repo.db.Exec(ctx, query, investor.FirstName, investor.LastName, investor.Balance, investor.Email, investor.Password)
	if err != nil {
		return appErr.NewUnexpectedError("Error inserting invoice : " + err.Error())
	}
	return nil
}
