package store

import "github.com/hossein1990/master_ticket/domain"

type Order struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	UserId      uint   `json:"user_id"`
	TicketId    uint   `json:"ticket_id"`
	Price       uint   `json:"price"`
	Ticket      string `json:"ticket"`
	Transaction string `json:"transaction"`
	State       uint   `json:"state"`
}

func mapFromOrderEntity(order domain.Order) Order {
	return Order{
		Id:          order.Id,
		UserId:      order.UserId,
		TicketId:    order.TicketId,
		Price:       order.Price,
		Ticket:      order.Ticket,
		Transaction: order.Transaction,
		State:       order.State,
	}
}
func mapTOrderEntity(order Order) domain.Order {
	return domain.Order{
		Id:          order.Id,
		UserId:      order.UserId,
		TicketId:    order.TicketId,
		Price:       order.Price,
		Ticket:      order.Ticket,
		Transaction: order.Transaction,
		State:       order.State,
	}
}
