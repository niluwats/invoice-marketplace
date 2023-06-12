package domain

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Investor struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int64  `json:"balance"`
}

type InvestorRepositoryDb struct {
	db *pgxpool.Pool
}

func NewInvestorRepositoryDb(dbclient *pgxpool.Pool) InvestorRepositoryDb {
	return InvestorRepositoryDb{db: dbclient}
}

func (repo InvestorRepositoryDb) GetById(id int) (*Investor, error) {
	query := "SELECT * FROM INVESTORS WHERE ID=$1"

	var investor Investor
	row := repo.db.QueryRow(context.Background(), query, id)

	err := row.Scan(&investor)
	if err != nil {
		log.Println("Error scanning investor : ", err)
		return nil, fmt.Errorf("Error scanning investor : %v", err)
	}

	return &investor, nil
}

func (repo InvestorRepositoryDb) GetAll() ([]Investor, error) {
	query := "SELECT * FROM INVESTORS"

	investors := make([]Investor, 0)
	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		log.Println("Error querying investors : ", err)
		return nil, fmt.Errorf("Error querying investors : %v", err)
	}

	for rows.Next() {
		var investor Investor
		err := rows.Scan(&investor)
		if err != nil {
			log.Println("Error scanning investors : ", err)
			return nil, fmt.Errorf("Error scanning investors : %v", err)
		}

		investors = append(investors, investor)
	}

	return investors, nil
}
