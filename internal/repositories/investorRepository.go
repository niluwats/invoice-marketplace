package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
)

type InvestorRepository interface {
	FindById(id int) (*domain.Investor, error)
	FindAll() ([]domain.Investor, error)
	FindByEmail(email string) (*domain.Investor, error)
	Save(investor domain.Investor) error
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

	query := "SELECT id,first_name,last_name,balance,is_issuer FROM investors WHERE ID=$1"

	var investor domain.Investor
	row := repo.db.QueryRow(ctx, query, id)

	err := row.Scan(&investor.ID, &investor.FirstName, &investor.LastName, &investor.Balance, &investor.IsIssuer)
	if err != nil {
		log.Println("Error scanning investor : ", err)
		return nil, fmt.Errorf("Error scanning investor : %s", err)
	}

	return &investor, nil
}

func (repo InvestorRepositoryDb) FindAll() ([]domain.Investor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT id,first_name,last_name,balance FROM investors"

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

func (repo InvestorRepositoryDb) FindByEmail(email string) (*domain.Investor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT id,password,is_issuer FROM investors WHERE email=$1"

	var user domain.Investor
	row := repo.db.QueryRow(ctx, query, email)

	err := row.Scan(&user.ID, &user.Password, &user.IsIssuer)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("Email not found")
		}
		log.Println("Error scanning investor : ", err)
		return nil, fmt.Errorf("Error scanning investor : %s", err)
	}

	return &user, nil
}

func (repo InvestorRepositoryDb) Save(investor domain.Investor) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	user, err := repo.FindByEmail(investor.Email)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("Email already taken ")
	}

	query := `INSERT INTO investors(first_name,last_name,balance,email,password,is_active) 
				VALUES($1,$2,$3,$4,$5,true)`

	_, err = repo.db.Exec(ctx, query, investor.FirstName, investor.LastName, investor.Balance, investor.Email, investor.Password)
	if err != nil {
		return fmt.Errorf("Error inserting invoice : %s", err)
	}
	return nil
}
