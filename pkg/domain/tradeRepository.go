package domain

type TradeRepository interface {
	Insert(trade Trade) error
	UpdateApproval(id int) error
}
