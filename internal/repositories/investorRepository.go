package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
)

type InvestorRepository interface {
	FindById(id int) (*domain.Investor, error)
	FindAll() ([]domain.Investor, error)
	UpdateBalance(id int, amount float64) error
}

type InvestorRepositoryDb struct {
	db *pgxpool.Pool
}

func NewInvestorRepositoryDb(dbclient *pgxpool.Pool) InvestorRepositoryDb {
	return InvestorRepositoryDb{db: dbclient}
}

func (repo InvestorRepositoryDb) FindById(id int) (*domain.Investor, error) {
	query := "SELECT * FROM INVESTORS WHERE ID=$1"

	var investor domain.Investor
	row := repo.db.QueryRow(context.Background(), query, id)

	err := row.Scan(&investor)
	if err != nil {
		log.Println("Error scanning investor : ", err)
		return nil, fmt.Errorf("Error scanning investor : %s", err)
	}

	return &investor, nil
}

func (repo InvestorRepositoryDb) FindAll() ([]domain.Investor, error) {
	query := "SELECT * FROM INVESTORS"

	investors := make([]domain.Investor, 0)
	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		log.Println("Error querying investors : ", err)
		return nil, fmt.Errorf("Error querying investors : %s", err)
	}

	for rows.Next() {
		var investor domain.Investor
		err := rows.Scan(&investor)
		if err != nil {
			log.Println("Error scanning investors : ", err)
			return nil, fmt.Errorf("Error scanning investors : %s", err)
		}

		investors = append(investors, investor)
	}

	return investors, nil
}

func (repo InvestorRepositoryDb) UpdateBalance(id int, amount float64) error {
	query := "UPDATE INVESTORS SET balance=balance-$1 WHERE ID=$2"

	_, err := repo.db.Exec(context.Background(), query, amount, id)
	if err != nil {
		log.Println("Error updating investors balance : ", err)
		return fmt.Errorf("Error updating investor's balance : %s", err)
	}
	return nil
}
