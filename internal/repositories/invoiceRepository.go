package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

type InvoiceRepository interface {
	Insert(invoice domain.Invoice) (*domain.Invoice, *appErr.AppError)
	FindById(id int) (*domain.Invoice, *appErr.AppError)
	FindAll() ([]domain.Invoice, *appErr.AppError)
	FindTotalInvestment(id int) (float64, *appErr.AppError)
}

type InvoiceRepositoryDb struct {
	db *pgxpool.Pool
}

func NewInvoiceRepositoryDb(dbclient *pgxpool.Pool) InvoiceRepositoryDb {
	return InvoiceRepositoryDb{db: dbclient}
}

func (repo InvoiceRepositoryDb) Insert(invoice domain.Invoice) (*domain.Invoice, *appErr.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `INSERT INTO INVOICE(invoice_number,amount_due,asking_price,duedate,created_on,issuer_id,is_locked,is_traded) 
				VALUES($1,$2,$3,$4,$5,$6,false,false) RETURNING id`

	var id int
	err := repo.db.QueryRow(ctx, query, invoice.InvoiceNumber, invoice.AmountDue, invoice.AskingPrice,
		invoice.DueDate, invoice.CreatedOn, invoice.IssuerId).Scan(&id)
	if err != nil {
		return nil, appErr.NewUnexpectedError("Error inserting invoice : " + err.Error())
	}

	createdInvoice, err_ := repo.FindById(id)
	if err_ != nil {
		return nil, err_
	}

	return createdInvoice, nil
}

func (repo InvoiceRepositoryDb) FindById(id int) (*domain.Invoice, *appErr.AppError) {
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
			return nil, appErr.NewNotFoundError("Invoice not found ")
		}
		return nil, appErr.NewUnexpectedError("Error querying invoice : " + err.Error())
	}

	return &invoice, nil
}

func (repo InvoiceRepositoryDb) FindAll() ([]domain.Invoice, *appErr.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT id,invoice_number,asking_price,created_on,is_locked,is_traded,issuer_id FROM invoice ORDER BY id"

	invoices := make([]domain.Invoice, 0)
	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, appErr.NewUnexpectedError("Error querying invoices : " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var invoice domain.Invoice
		err := rows.Scan(&invoice.ID, &invoice.InvoiceNumber, &invoice.AskingPrice, &invoice.CreatedOn, &invoice.IsLocked, &invoice.IsTraded, &invoice.IssuerId)
		if err != nil {
			return nil, appErr.NewUnexpectedError("Error scanning invoices : " + err.Error())
		}

		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

func (repo InvoiceRepositoryDb) FindTotalInvestment(id int) (float64, *appErr.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT COALESCE(SUM(bid_amount),0)FROM bids WHERE invoice_id=$1 AND status=1"
	row := repo.db.QueryRow(ctx, query, id)

	var sum float64
	err := row.Scan(&sum)
	if err != nil {
		return 0, appErr.NewUnexpectedError("Error querying invoice sum : " + err.Error())
	}
	return sum, nil
}
