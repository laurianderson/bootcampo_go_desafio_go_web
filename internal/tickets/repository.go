package tickets

import (
	"context"
	"fmt"
	"github.com/laurianderson/bootcamp_go_desafio_go_web/internal/domain"
)

//Interface
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

//Struct
type repository struct {
	db []domain.Ticket
}

//Builder
func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

//Get all tickets
func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

//Get ticket by destination
func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

//Get total tickets by destination
func (r *repository) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {

    var ticketsDest []domain.Ticket

    if len(r.db) == 0 {
        return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
    }

    for _, t := range r.db {
        if t.Country == destination {
            ticketsDest = append(ticketsDest, t)
        }
    }

    return ticketsDest, nil
}

//Average destination
func (r *repository) AverageDestination(ctx context.Context, destination string) (float64, error) {

    var ticketsDest []domain.Ticket

    if len(r.db) == 0 {
        return 0, fmt.Errorf("empty list of tickets")
    }

    for _, t := range r.db {
        if t.Country == destination {
            ticketsDest = append(ticketsDest, t)
        }
    }

    if len(ticketsDest) == 0 {
        return 0, fmt.Errorf("empty list of tickets")
    }

    return float64(len(ticketsDest)) / float64(len(r.db)), nil
}