package tickets

import (
	"context"
	"github.com/laurianderson/bootcamp_go_desafio_go_web/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context ,destination string) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
    rp Repository
}

func NewService(rp Repository) Service {
    return &service{
        rp: rp,
    }
}

//Get all tickets
func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
    return s.rp.GetAll(ctx)
}

//Get ticket by destination
func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	return s.rp.GetTicketByDestination(ctx, destination)
}

//Get total tickets by destination
func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
    return s.rp.GetTotalTickets(ctx, destination)
}

//Average destination
func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
    return s.rp.AverageDestination(ctx, destination)
}