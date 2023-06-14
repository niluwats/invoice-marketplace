package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/niluwats/invoice-marketplace/internal/domain"
)

type BidRepository interface {
	Insert(bid domain.Bid) error
	UpdateApproval(id int) error
	GetAll(invoiceId int) ([]domain.Bid, error)
}

type BidRepositoryDb struct {
	db *pgxpool.Pool
}

func NewBidRepositoryDb(dbclient *pgxpool.Pool) BidRepositoryDb {
	return BidRepositoryDb{dbclient}
}

func (repo BidRepositoryDb) Insert(bid domain.Bid) error {
	query := `INSERT INTO bids(invoice_id,bid_amount,timestamp,investor_id,is_approved) 
			  VALUES($1,$2,$3,$4,$5)`

	_, err := repo.db.Exec(context.Background(), query, bid.InvoiceId, bid.BidAmount, bid.TimeStamp, bid.InvestorId, false)
	if err != nil {
		log.Println("Error inserting bid : ", err)
		return fmt.Errorf("Error inserting bid : %s", err)
	}
	return nil
}

func (repo BidRepositoryDb) UpdateApproval(id int) error {
	query := "UPDATE bids SET is_approved=true WHERE id=$1"

	_, err := repo.db.Exec(context.Background(), query, id)
	if err != nil {
		log.Println("Error updating approval status of bid : ", err)
		return fmt.Errorf("Error updating approval status of bid : %s", err)
	}
	return nil
}

func (repo BidRepositoryDb) GetAll(invoiceId int) ([]domain.Bid, error) {
	query := "SELECT * FROM bids WHERE invoice_id=$1"

	var bids []domain.Bid
	rows, err := repo.db.Query(context.Background(), query, invoiceId)
	if err != nil {
		return nil, fmt.Errorf("Error querying bids :%s", err)
	} else {
		for rows.Next() {
			var bid domain.Bid
			err := rows.Scan(&bid)
			if err != nil {
				return nil, fmt.Errorf("Error scanning bids :%s", err)
			}
			bids = append(bids, bid)
		}
		return bids, nil
	}
}
