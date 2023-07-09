package models

type PaymentStatus string

const (
	Approved = "approved"
	Rejected = "rejected"
)

func (s PaymentStatus) String() string {
	switch s {
	case Approved:
		return Approved
	case Rejected:
		return Rejected
	}
	return "unknown"
}

type OrderStatus string

const (
	Pending   = "pending"
	Received  = "received"
	Preparing = "preparing"
	Ready     = "ready"
	Finished  = "finished"
	Cancelled = "cancelled"
)

func (s OrderStatus) String() string {
	switch s {
	case Pending:
		return Pending
	case Received:
		return Received
	case Preparing:
		return Preparing
	case Ready:
		return Ready
	case Finished:
		return Finished
	}
	return "unknown"
}
