package store

import "github.com/hossein1990/master_ticket/domain"

type Ticket struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	TypeTicket  uint   `json:"type_ticket"`
	Active      uint   `json:"active"`
	Price       uint   `json:"price"`
	CreatedAt   string `json:"created_at"`
}

func mapFromTicketEntity(tiket domain.Ticket) Ticket {
	return Ticket{
		Id:          tiket.Id,
		Name:        tiket.Name,
		Description: tiket.Description,
		StartDate:   tiket.StartDate,
		TypeTicket:  tiket.TypeTicket,
		Active:      tiket.Active,
		Price:       tiket.Price,
		CreatedAt:   tiket.CreatedAt,
	}
}
func mapToTicketEntity(tiket Ticket) domain.Ticket {
	return domain.Ticket{
		Id:          tiket.Id,
		Name:        tiket.Name,
		Description: tiket.Description,
		StartDate:   tiket.StartDate,
		TypeTicket:  tiket.TypeTicket,
		Active:      tiket.Active,
		Price:       tiket.Price,
		CreatedAt:   tiket.CreatedAt,
	}
}
