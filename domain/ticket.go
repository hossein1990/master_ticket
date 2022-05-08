package domain

type Ticket struct {
	Id          uint   `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	StartDate   string `db:"start_date"`
	TypeTicket  uint   `db:"type_ticket"`
	Active      uint   `db:"active"`
	Price       uint   `db:"price"`
	CreatedAt   string `db:"created_at"`
}
