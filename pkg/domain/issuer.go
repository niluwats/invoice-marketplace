package domain

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Issuer struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int64  `json:"balance"`
}

type IssuerRepositoryDb struct {
	db *pgxpool.Pool
}

func NewIssuerRepositoryDb(dbclient *pgxpool.Pool) IssuerRepositoryDb {
	return IssuerRepositoryDb{db: dbclient}
}

func (repo IssuerRepositoryDb) GetById(id int) (*Issuer, error) {
	query := "SELECT * FROM ISSUERS WHERE ID=$1"

	var issuer Issuer
	row := repo.db.QueryRow(context.Background(), query, id)

	err := row.Scan(&issuer)
	if err != nil {
		return nil, fmt.Errorf("Error scanning issuer : %v", err)
	}

	return &issuer, nil
}

func (repo IssuerRepositoryDb) GetAll() ([]Issuer, error) {
	query := "SELECT * FROM ISSUERS"

	issuers := make([]Issuer, 0)
	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("Error querying issuers : %v", err)
	}

	for rows.Next() {
		var issuer Issuer
		err := rows.Scan(&issuer)
		if err != nil {
			return nil, fmt.Errorf("Error scanning issuers : %v", err)
		}

		issuers = append(issuers, issuer)
	}

	return issuers, nil
}
