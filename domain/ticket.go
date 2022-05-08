package domain

type Ticket struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	TypeTicket  uint   `json:"type_ticket"`
	Active      uint   `json:"active"`
	Price       uint   `json:"price"`
	CreatedAt   string `json:"created_at"`
}
