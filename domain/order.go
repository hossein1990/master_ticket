package domain

type Order struct {
	Id          uint   `json:"id"`
	UserId      uint   `json:"user_id"`
	TicketId    uint   `json:"ticket_id"`
	Price       uint   `json:"price"`
	Ticket      string `json:"ticket"`
	Transaction string `json:"transaction"`
	State       uint   `json:"state"`
}
