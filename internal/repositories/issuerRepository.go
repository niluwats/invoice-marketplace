package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
)

type IssuerRepository interface {
	FindById(id int) (*domain.Issuer, error)
	FindAll() ([]domain.Issuer, error)
	UpdateBalance(id int, amount float64) error
}

type IssuerRepositoryDb struct {
	db *pgxpool.Pool
}

func NewIssuerRepositoryDb(dbclient *pgxpool.Pool) IssuerRepositoryDb {
	return IssuerRepositoryDb{db: dbclient}
}

func (repo IssuerRepositoryDb) FindById(id int) (*domain.Issuer, error) {
	query := "SELECT * FROM ISSUERS WHERE ID=$1"

	var issuer domain.Issuer
	row := repo.db.QueryRow(context.Background(), query, id)

	err := row.Scan(&issuer)
	if err != nil {
		return nil, fmt.Errorf("Error scanning issuer : %v", err)
	}

	return &issuer, nil
}

func (repo IssuerRepositoryDb) FindAll() ([]domain.Issuer, error) {
	query := "SELECT * FROM ISSUERS"

	issuers := make([]domain.Issuer, 0)
	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("Error querying issuers : %v", err)
	}

	for rows.Next() {
		var issuer domain.Issuer
		err := rows.Scan(&issuer)
		if err != nil {
			return nil, fmt.Errorf("Error scanning issuers : %v", err)
		}

		issuers = append(issuers, issuer)
	}

	return issuers, nil
}

func (repo IssuerRepositoryDb) UpdateBalance(id int, amount float64) error {
	query := "UPDATE ISSUERS SET balance=balance+$1 WHERE ID=$2"

	_, err := repo.db.Exec(context.Background(), query, amount, id)
	if err != nil {
		log.Println("Error updating issuer's balance : ", err)
		return fmt.Errorf("Error updating issuer's balance : %v", err)
	}
	return nil
}
