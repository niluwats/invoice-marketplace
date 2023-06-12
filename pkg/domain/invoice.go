package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Invoice struct {
	ID                int       `json:"id"`
	InvoiceNumber     string    `json:"invoice_number"`
	CustomerFirstName string    `json:"customer_first_name"`
	CustomerLastName  string    `json:"customer_last_name"`
	AmountDue         float64   `json:"amount_due"`
	AmountEnclosed    float64   `json:"amount_enclosed"`
	DueDate           time.Time `json:"duedate"`
	IsLocked          bool      `json:"is_locked"`
	IsTraded          bool      `json:"is_traded"`
	InvestorId        int       `json:"investor_id"`
}

type InvoiceRepositoryDb struct {
	db *pgxpool.Pool
}

func NewInvoiceRepositoryDb(dbclient *pgxpool.Pool) InvoiceRepositoryDb {
	return InvoiceRepositoryDb{db: dbclient}
}

func (repo InvoiceRepositoryDb) Insert(invoice Invoice) error {
	query := `INSERT INTO INVOICE(invoice_number,customer_first_name,customer_last_name,
			amount_due,amount_enclosed,duedate) VALUES($1,$2,$3,$4,$5,$6)`

	_, err := repo.db.Exec(context.Background(), query, invoice.InvoiceNumber, invoice.CustomerFirstName, invoice.CustomerLastName, invoice.AmountDue, invoice.AmountEnclosed, invoice.DueDate)
	if err != nil {
		return fmt.Errorf("Error inserting invoice : %v", err)
	}
	return nil
}
