package domain

type Orders struct {
	Id          uint   `db:"id"`
	UserId      uint   `db:"user_id"`
	TicketId    uint   `db:"ticket_id"`
	Price       uint   `db:"price"`
	Ticket      string `db:"ticket"`
	Transaction string `db:"transaction"`
	State       uint   `db:"state"`
}
