package repositories

import (
	"context"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

var mutexMap sync.Map
var dbTimeOut = time.Second * 3

type BidRepository interface {
	ProcessBid(ctx *context.Context, bid domain.Bid, restBalance float64) (*domain.Bid, *appErr.AppError)
	ProcessApproveBid(ctx *context.Context, invoiceid, issuerid int, amount float64) *appErr.AppError
	ProcessCancelBid(ctx *context.Context, invoiceid int) *appErr.AppError
	GetAll(ctx *context.Context, invoiceId int) ([]domain.Bid, *appErr.AppError)
	GetBid(ctx *context.Context, id int) (*domain.Bid, *appErr.AppError)
}

type BidRepositoryDb struct {
	db *pgxpool.Pool
}

func NewBidRepositoryDb(dbclient *pgxpool.Pool) BidRepositoryDb {
	return BidRepositoryDb{dbclient}
}

func (repo BidRepositoryDb) ProcessBid(ctx *context.Context, bid domain.Bid, restBalance float64) (*domain.Bid, *appErr.AppError) {
	invoiceMutex, _ := mutexMap.LoadOrStore(bid.InvoiceId, &sync.Mutex{})
	mutex := invoiceMutex.(*sync.Mutex)

	mutex.Lock()
	defer mutex.Unlock()

	tx, err := repo.db.BeginTx(*ctx, pgx.TxOptions{})
	if err != nil {
		return nil, appErr.NewUnexpectedError(err.Error())
	}

	defer tx.Rollback(*ctx)

	var id int

	err = tx.QueryRow(*ctx, "SELECT SAVE_BID($1,$2,$3,$4)", bid.InvoiceId, bid.BidAmount, bid.CreatedAt, bid.InvestorId).Scan(&id)
	if err != nil {
		return nil, appErr.NewUnexpectedError("Error inserting bid : " + err.Error())
	}

	_, err = tx.Exec(*ctx, "CALL UPDATE_INVESTOR_BALANCE($1,$2)", bid.BidAmount, bid.InvestorId)
	if err != nil {
		return nil, appErr.NewUnexpectedError("Error updating investor's balance : " + err.Error())
	}

	if restBalance <= bid.BidAmount {
		_, err = tx.Exec(*ctx, "CALL UPDATE_INVOICE_STATUS($1,true)", bid.InvoiceId)
		if err != nil {
			return nil, appErr.NewUnexpectedError("Error updating invoice lock status : " + err.Error())
		}
	}

	err = tx.Commit(*ctx)
	if err != nil {
		return nil, appErr.NewUnexpectedError(err.Error())
	}

	createdBid, err_ := repo.GetBid(ctx, id)
	if err_ != nil {
		return nil, err_
	}

	return createdBid, nil
}

func (repo BidRepositoryDb) ProcessApproveBid(ctx *context.Context, invoiceid, issuerid int, amount float64) *appErr.AppError {
	//begin tx
	tx, err := repo.db.BeginTx(*ctx, pgx.TxOptions{})
	if err != nil {
		return appErr.NewUnexpectedError(err.Error())
	}
	defer tx.Rollback(*ctx)

	//update is_approved in bid
	_, err = tx.Exec(*ctx, "CALL UPDATE_BID_APPROVAL_STATUS($1)", invoiceid)
	if err != nil {
		return appErr.NewUnexpectedError("Error updating approval status of bid : " + err.Error())
	}

	//update issuer balance
	_, err = tx.Exec(*ctx, "CALL UPDATE_ISSUER_BALANCE($1,$2)", amount, issuerid)
	if err != nil {
		return appErr.NewUnexpectedError("Error updating issuer's balance : " + err.Error())
	}

	//update is_traded in invoice
	_, err = tx.Exec(*ctx, "CALL UPDATE_INVOICE_TRADED_STATUS($1)", invoiceid)
	if err != nil {
		return appErr.NewUnexpectedError("Error updating invoice is_traded status : " + err.Error())
	}

	if err = tx.Commit(*ctx); err != nil {
		return appErr.NewUnexpectedError(err.Error())
	}
	return nil
}

func (repo BidRepositoryDb) ProcessCancelBid(ctx *context.Context, invoiceid int) *appErr.AppError {
	tx, err := repo.db.BeginTx(*ctx, pgx.TxOptions{})
	if err != nil {
		return appErr.NewUnexpectedError(err.Error())
	}

	defer tx.Rollback(*ctx)

	//update status of bid to 0 where invoiceid = __
	_, err = tx.Exec(*ctx, "CALL UPDATE_BID_STATUS($1)", invoiceid)
	if err != nil {
		return appErr.NewUnexpectedError("Error updating bid status : " + err.Error())
	}

	//update invoice is_locked status to false
	_, err = tx.Exec(*ctx, "CALL UPDATE_INVOICE_STATUS($1,false)", invoiceid)
	if err != nil {
		return appErr.NewUnexpectedError("Error updating invoice is_locked status : " + err.Error())
	}

	//update balances of investors who bid on that invoice
	_, err = tx.Exec(*ctx, "CALL UPDATE_ALL_INVESTORS_BALANCES($1)", invoiceid)
	if err != nil {
		return appErr.NewUnexpectedError("Error updating investor's balances : " + err.Error())
	}

	if err = tx.Commit(*ctx); err != nil {
		return appErr.NewUnexpectedError(err.Error())
	}

	return nil
}

func (repo BidRepositoryDb) GetAll(ctx *context.Context, invoiceId int) ([]domain.Bid, *appErr.AppError) {
	var bids []domain.Bid
	rows, err := repo.db.Query(*ctx, "SELECT * FROM GET_ALL_BIDS_BY_INVOICE($1)", invoiceId)
	defer rows.Close()

	if err != nil {
		return nil, appErr.NewUnexpectedError("Error querying bids : " + err.Error())
	} else {
		for rows.Next() {
			var bid domain.Bid
			err := rows.Scan(&bid.ID, &bid.BidAmount, &bid.CreatedAt, &bid.IsApproved, &bid.InvoiceId, &bid.InvestorId, &bid.Status)
			if err != nil {
				return nil, appErr.NewUnexpectedError("Error scanning bids : " + err.Error())
			}

			bids = append(bids, bid)
		}
		return bids, nil
	}
}

func (repo BidRepositoryDb) GetBid(ctx *context.Context, id int) (*domain.Bid, *appErr.AppError) {
	var bid domain.Bid
	row := repo.db.QueryRow(*ctx, "SELECT * FROM GET_BID_BY_ID($1)", id)

	err := row.Scan(&bid.ID, &bid.BidAmount, &bid.CreatedAt, &bid.IsApproved, &bid.InvoiceId, &bid.InvestorId, &bid.Status)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, appErr.NewNotFoundError("Bid not found ")
		}
		return nil, appErr.NewUnexpectedError("Error querying bid : " + err.Error())
	}
	return &bid, nil
}
