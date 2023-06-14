package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
)

type InvoiceRepository interface {
	Insert(invoice domain.Invoice) error
	FindById(id int) (*domain.Invoice, error)
	UpdateLockStatus(id int) error
	UpdateInvoiceInvestor(id, investor int) error
}

type InvoiceRepositoryDb struct {
	db *pgxpool.Pool
}

func NewInvoiceRepositoryDb(dbclient *pgxpool.Pool) InvoiceRepositoryDb {
	return InvoiceRepositoryDb{db: dbclient}
}

func (repo InvoiceRepositoryDb) Insert(invoice domain.Invoice) error {
	query := `INSERT INTO INVOICE(invoice_number,amount_due,amount_enclosed,duedate,created_on,issuer_id) 
				VALUES($1,$2,$3,$4,$5,$6)`

	_, err := repo.db.Exec(context.Background(), query, invoice.InvoiceNumber, invoice.AmountDue, invoice.AmountEnclosed,
		invoice.DueDate, invoice.CreatedOn, invoice.IssuerId)
	if err != nil {
		return fmt.Errorf("Error inserting invoice : %s", err)
	}
	return nil
}

func (repo InvoiceRepositoryDb) FindById(id int) (*domain.Invoice, error) {
	query := "SELECT * FROM INVOICE WHERE ID=$1"

	row := repo.db.QueryRow(context.Background(), query, id)

	var invoice domain.Invoice
	err := row.Scan(&invoice)
	if err != nil {
		return nil, fmt.Errorf("Error scanning invoice : %s", err)
	}

	return &invoice, nil
}

func (repo InvoiceRepositoryDb) UpdateLockStatus(id int) error {
	query := "UPDATE INVOICE SET is_locked=true WHERE ID=$1"

	_, err := repo.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("Error updating invoice locked status : %s", err)
	}
	return nil
}

func (repo InvoiceRepositoryDb) UpdateInvoiceInvestor(id, investor int) error {
	query := "UPDATE INVOICE SET investor_id=$1 WHERE ID=$2"

	_, err := repo.db.Exec(context.Background(), query, investor, id)
	if err != nil {
		return fmt.Errorf("Error updating invoice investor id : %s", err)
	}
	return nil
}
