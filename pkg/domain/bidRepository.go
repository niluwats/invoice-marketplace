package domain

type BidRepository interface {
	Insert(bid Bid) error
	UpdateApproval(id int) error
}
