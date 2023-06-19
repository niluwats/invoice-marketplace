package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

var dbTimeOut = time.Second * 3

type BidRepository interface {
	ProcessBid(bid domain.Bid, restBalance float64) (*domain.Bid, *appErr.AppError)
	ProcessApproveBid(invoiceid, issuerid int, amount float64) *appErr.AppError
	GetAll(invoiceId int) ([]domain.Bid, *appErr.AppError)
	GetBid(id int) (*domain.Bid, *appErr.AppError)
}

type BidRepositoryDb struct {
	db *pgxpool.Pool
}

func NewBidRepositoryDb(dbclient *pgxpool.Pool) BidRepositoryDb {
	return BidRepositoryDb{dbclient}
}

func (repo BidRepositoryDb) ProcessBid(bid domain.Bid, restBalance float64) (*domain.Bid, *appErr.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, appErr.NewUnexpectedError(err.Error())
	}

	defer tx.Rollback(ctx)

	var id int
	query := `INSERT INTO bids(invoice_id,bid_amount,timestamp,investor_id,is_approved) 
			  VALUES($1,$2,$3,$4,$5) RETURNING id`

	err = tx.QueryRow(ctx, query, bid.InvoiceId, bid.BidAmount, bid.TimeStamp, bid.InvestorId, false).Scan(&id)
	if err != nil {
		return nil, appErr.NewUnexpectedError("Error inserting bid : " + err.Error())
	}

	query = "UPDATE investors SET balance=balance-$1 WHERE ID=$2"

	_, err = tx.Exec(ctx, query, bid.BidAmount, bid.InvestorId)
	if err != nil {
		return nil, appErr.NewUnexpectedError("Error updating investor's balance : " + err.Error())
	}

	if restBalance <= bid.BidAmount {
		query = "UPDATE INVOICE SET is_locked=true WHERE ID=$1"

		_, err = tx.Exec(ctx, query, bid.InvoiceId)
		if err != nil {
			return nil, appErr.NewUnexpectedError("Error updating invoice lock status : " + err.Error())
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, appErr.NewUnexpectedError(err.Error())
	}

	createdBid, err_ := repo.GetBid(id)
	if err_ != nil {
		return nil, err_
	}

	return createdBid, nil
}

func (repo BidRepositoryDb) ProcessApproveBid(invoiceid, issuerid int, amount float64) *appErr.AppError {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	//begin tx
	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return appErr.NewUnexpectedError(err.Error())
	}
	defer tx.Rollback(ctx)

	//update is_approved in bid
	query := "UPDATE bids SET is_approved=true WHERE invoice_id=$1"

	_, err = tx.Exec(ctx, query, invoiceid)
	if err != nil {
		return appErr.NewUnexpectedError("Error updating approval status of bid : " + err.Error())
	}

	//update issuer balance
	query = `UPDATE investors SET balance=balance+$1 FROM issuers 
			WHERE issuers.investor_id=investors.id and issuers.id=$2`

	_, err = tx.Exec(ctx, query, amount, issuerid)
	if err != nil {
		return appErr.NewUnexpectedError("Error updating issuer's balance : " + err.Error())
	}

	//update is_traded in invoice
	query = "UPDATE invoice SET is_traded=true WHERE id=$1"
	_, err = tx.Exec(ctx, query, invoiceid)
	if err != nil {
		return appErr.NewUnexpectedError("Error updating invoice is_traded status : " + err.Error())
	}

	if err = tx.Commit(ctx); err != nil {
		return appErr.NewUnexpectedError(err.Error())
	}
	return nil
}

func (repo BidRepositoryDb) GetAll(invoiceId int) ([]domain.Bid, *appErr.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT * FROM bids WHERE invoice_id=$1"

	var bids []domain.Bid
	rows, err := repo.db.Query(ctx, query, invoiceId)
	defer rows.Close()

	if err != nil {
		return nil, appErr.NewUnexpectedError("Error querying bids : " + err.Error())
	} else {
		for rows.Next() {
			var bid domain.Bid
			err := rows.Scan(&bid.ID, &bid.BidAmount, &bid.TimeStamp, &bid.IsApproved, &bid.InvoiceId, &bid.InvestorId)
			if err != nil {
				return nil, appErr.NewUnexpectedError("Error scanning bids : " + err.Error())
			}
			bids = append(bids, bid)
		}
		return bids, nil
	}
}

func (repo BidRepositoryDb) GetBid(id int) (*domain.Bid, *appErr.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT * FROM bids WHERE ID=$1"

	var bid domain.Bid
	row := repo.db.QueryRow(ctx, query, id)

	err := row.Scan(&bid.ID, &bid.BidAmount, &bid.TimeStamp, &bid.IsApproved, &bid.InvoiceId, &bid.InvestorId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, appErr.NewNotFoundError("Bid not found ")
		}
		return nil, appErr.NewUnexpectedError("Error querying bid : " + err.Error())
	}

	return &bid, nil
}
