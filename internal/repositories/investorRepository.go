package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
)

type InvestorRepository interface {
	FindById(id int) (*domain.Investor, error)
	FindAll() ([]domain.Investor, error)
	FindInvestorBalance(id int) (float64, error)
}

type InvestorRepositoryDb struct {
	db *pgxpool.Pool
}

func NewInvestorRepositoryDb(dbclient *pgxpool.Pool) InvestorRepositoryDb {
	return InvestorRepositoryDb{db: dbclient}
}

func (repo InvestorRepositoryDb) FindById(id int) (*domain.Investor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT * FROM investors WHERE ID=$1"

	var investor domain.Investor
	row := repo.db.QueryRow(ctx, query, id)

	err := row.Scan(&investor.ID, &investor.FirstName, &investor.LastName, &investor.Balance)
	if err != nil {
		log.Println("Error scanning investor : ", err)
		return nil, fmt.Errorf("Error scanning investor : %s", err)
	}

	return &investor, nil
}

func (repo InvestorRepositoryDb) FindAll() ([]domain.Investor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT * FROM investors"

	investors := make([]domain.Investor, 0)
	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		log.Println("Error querying investors : ", err)
		return nil, fmt.Errorf("Error querying investors : %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var investor domain.Investor
		err := rows.Scan(&investor.ID, &investor.FirstName, &investor.LastName, &investor.Balance)
		if err != nil {
			log.Println("Error scanning investors : ", err)
			return nil, fmt.Errorf("Error scanning investors : %s", err)
		}

		investors = append(investors, investor)
	}

	return investors, nil
}

func (repo InvestorRepositoryDb) FindInvestorBalance(id int) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT balance FROM investors WHERE id=$1"
	row := repo.db.QueryRow(ctx, query, id)

	var balance float64
	err := row.Scan(&balance)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, fmt.Errorf("No investor found")
		}
		return 0, fmt.Errorf("Error scanning investor : %s", err)
	}
	return balance, nil
}
