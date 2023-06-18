package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
)

type InvoiceRepository interface {
	Insert(invoice domain.Invoice) error
	FindById(id int) (*domain.Invoice, error)
	FindTotalInvestment(id int) (float64, error)
}

type InvoiceRepositoryDb struct {
	db *pgxpool.Pool
}

func NewInvoiceRepositoryDb(dbclient *pgxpool.Pool) InvoiceRepositoryDb {
	return InvoiceRepositoryDb{db: dbclient}
}

func (repo InvoiceRepositoryDb) Insert(invoice domain.Invoice) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `INSERT INTO INVOICE(invoice_number,amount_due,asking_price,duedate,created_on,issuer_id,is_locked,is_traded) 
				VALUES($1,$2,$3,$4,$5,$6,false,false)`

	_, err := repo.db.Exec(ctx, query, invoice.InvoiceNumber, invoice.AmountDue, invoice.AskingPrice,
		invoice.DueDate, invoice.CreatedOn, invoice.IssuerId)
	if err != nil {
		return fmt.Errorf("Error inserting invoice : %s", err)
	}
	return nil
}

func (repo InvoiceRepositoryDb) FindById(id int) (*domain.Invoice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT 
				invoice.id,
				invoice.invoice_number,
				invoice.amount_due,
				invoice.asking_price,
				invoice.created_on,
				invoice.duedate,
				invoice.is_locked,
				invoice.is_traded,
				invoice.issuer_id,
				CASE 
					WHEN invoice.is_traded=true THEN ARRAY_AGG(bids.investor_id)
					ELSE NULL
				END AS investors
				FROM invoice LEFT JOIN bids ON invoice.id = bids.invoice_id 
				WHERE invoice.id = $1 GROUP BY invoice.id`

	row := repo.db.QueryRow(ctx, query, id)

	var invoice domain.Invoice
	err := row.Scan(&invoice.ID, &invoice.InvoiceNumber, &invoice.AmountDue, &invoice.AskingPrice,
		&invoice.CreatedOn, &invoice.DueDate, &invoice.IsLocked,
		&invoice.IsTraded, &invoice.IssuerId, &invoice.InvestorIds)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("No invoice found : %s", err)
		}
		return nil, fmt.Errorf("Error scanning invoice : %s", err)
	}

	return &invoice, nil
}

func (repo InvoiceRepositoryDb) FindTotalInvestment(id int) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT COALESCE(SUM(bid_amount),0)FROM bids WHERE invoice_id=$1"
	row := repo.db.QueryRow(ctx, query, id)

	var sum float64
	err := row.Scan(&sum)
	if err != nil {
		return 0, err
	}
	return sum, nil
}
