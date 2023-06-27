package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

type InvoiceRepository interface {
	Insert(ctx *context.Context, invoice domain.Invoice) (*domain.Invoice, *appErr.AppError)
	FindById(ctx *context.Context, id int) (*domain.Invoice, *appErr.AppError)
	FindAll(ctx *context.Context) ([]domain.Invoice, *appErr.AppError)
	FindTotalInvestment(ctx *context.Context, id int) (float64, *appErr.AppError)
	FindIfExistsByNo(ctx *context.Context, invoiceNo string) bool
}

type InvoiceRepositoryDb struct {
	db *pgxpool.Pool
}

func NewInvoiceRepositoryDb(dbclient *pgxpool.Pool) InvoiceRepositoryDb {
	return InvoiceRepositoryDb{db: dbclient}
}

func (repo InvoiceRepositoryDb) Insert(ctx *context.Context, invoice domain.Invoice) (*domain.Invoice, *appErr.AppError) {
	var id int
	err := repo.db.QueryRow(*ctx, "SELECT SAVE_INVOICE($1,$2,$3,$4,$5,false,false,$6)", invoice.InvoiceNumber, invoice.AmountDue, invoice.AskingPrice,
		invoice.CreatedOn, invoice.DueDate, invoice.IssuerId).Scan(&id)
	if err != nil {
		return nil, appErr.NewUnexpectedError("Error inserting invoice : " + err.Error())
	}

	createdInvoice, err_ := repo.FindById(ctx, id)
	if err_ != nil {
		return nil, err_
	}

	return createdInvoice, nil
}

func (repo InvoiceRepositoryDb) FindById(ctx *context.Context, id int) (*domain.Invoice, *appErr.AppError) {
	row := repo.db.QueryRow(*ctx, "SELECT * FROM GET_INVOICE_BY_ID($1)", id)

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

func (repo InvoiceRepositoryDb) FindAll(ctx *context.Context) ([]domain.Invoice, *appErr.AppError) {
	invoices := make([]domain.Invoice, 0)
	rows, err := repo.db.Query(*ctx, "SELECT * FROM GET_INVOICES()")
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

func (repo InvoiceRepositoryDb) FindTotalInvestment(ctx *context.Context, id int) (float64, *appErr.AppError) {
	row := repo.db.QueryRow(*ctx, "SELECT * FROM GET_TOTAL_INVESTMENT($1)", id)

	var sum float64
	err := row.Scan(&sum)
	if err != nil {
		return 0, appErr.NewUnexpectedError("Error querying invoice sum : " + err.Error())
	}
	return sum, nil
}

func (repo InvoiceRepositoryDb) FindIfExistsByNo(ctx *context.Context, invoiceNo string) bool {
	row := repo.db.QueryRow(*ctx, "SELECT * FROM GET_INVOICE_BY_INVOICENUMBER($1)", invoiceNo)

	var scannedInvNo string
	err := row.Scan(&scannedInvNo)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false
		}
		return false
	}

	return true
}
