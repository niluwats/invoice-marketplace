package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
)

var dbTimeOut = time.Second * 3

type BidRepository interface {
	ProcessBid(bid domain.Bid, restBalance float64) error
	ProcessApproveBid(invoiceid, issuerid int, amount float64) error
	GetAll(invoiceId int) ([]domain.Bid, error)
}

type BidRepositoryDb struct {
	db *pgxpool.Pool
}

func NewBidRepositoryDb(dbclient *pgxpool.Pool) BidRepositoryDb {
	return BidRepositoryDb{dbclient}
}

func (repo BidRepositoryDb) ProcessBid(bid domain.Bid, restBalance float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	query := `INSERT INTO bids(invoice_id,bid_amount,timestamp,investor_id,is_approved) 
			  VALUES($1,$2,$3,$4,$5)`

	_, err = tx.Exec(ctx, query, bid.InvoiceId, bid.BidAmount, bid.TimeStamp, bid.InvestorId, false)
	if err != nil {
		log.Println("Error inserting bid : ", err)
		return fmt.Errorf("Error inserting bid : %s", err)
	}

	query = "UPDATE investors SET balance=balance-$1 WHERE ID=$2"

	_, err = tx.Exec(ctx, query, bid.BidAmount, bid.InvestorId)
	if err != nil {
		log.Println("Error updating investor's balance : ", err)
		return fmt.Errorf("Error updating investor's balance : %s", err)
	}

	if restBalance <= bid.BidAmount {
		query = "UPDATE INVOICE SET is_locked=true WHERE ID=$1"

		_, err = tx.Exec(ctx, query, bid.InvoiceId)
		if err != nil {
			return fmt.Errorf("Error updating invoice lock status : %s", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repo BidRepositoryDb) ProcessApproveBid(invoiceid, issuerid int, amount float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	//begin tx
	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	//update is_approved in bid
	query := "UPDATE bids SET is_approved=true WHERE invoice_id=$1"

	_, err = tx.Exec(ctx, query, invoiceid)
	if err != nil {
		log.Println("Error updating approval status of bid : ", err)
		return fmt.Errorf("Error updating approval status of bid : %s", err)
	}

	//update issuer balance
	query = `UPDATE investors SET balance=balance+$1 FROM issuers 
			WHERE issuers.investor_id=investors.id and issuers.id=$2`

	_, err = tx.Exec(ctx, query, amount, issuerid)
	if err != nil {
		log.Println("Error updating issuer's balance : ", err)
		return fmt.Errorf("Error updating issuer's balance : %s", err)
	}

	//update is_traded in invoice
	query = "UPDATE invoice SET is_traded=true WHERE id=$1"
	_, err = tx.Exec(ctx, query, invoiceid)
	if err != nil {
		log.Println("Error updating invoice is_traded status : ", err)
		return fmt.Errorf("Error updating invoice is_traded status : %s", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (repo BidRepositoryDb) GetAll(invoiceId int) ([]domain.Bid, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT * FROM bids WHERE invoice_id=$1"

	var bids []domain.Bid
	rows, err := repo.db.Query(ctx, query, invoiceId)
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("Error querying bids :%s", err)
	} else {
		for rows.Next() {
			var bid domain.Bid
			err := rows.Scan(&bid.ID, &bid.BidAmount, &bid.TimeStamp, &bid.IsApproved, &bid.InvoiceId, &bid.InvestorId)
			if err != nil {
				return nil, fmt.Errorf("Error scanning bids :%s", err)
			}
			bids = append(bids, bid)
		}
		return bids, nil
	}
}
