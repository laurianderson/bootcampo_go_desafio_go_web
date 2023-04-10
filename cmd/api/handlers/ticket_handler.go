package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/laurianderson/bootcamp_go_desafio_go_web/internal/tickets"
)

type  ControllerTicket struct {
	service tickets.Service
}

func NewControllerTicket(s tickets.Service) *ControllerTicket {
	return &ControllerTicket{
		service: s,
	}
}
//Get all tickets
func(ct *ControllerTicket) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
        tickets, err := ct.service.GetAll(ctx)
        if err!= nil {
            ctx.String(http.StatusInternalServerError, err.Error(), nil)
            return
        }
        ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": tickets})
    }
}

//Get tickets by country
func (ct *ControllerTicket) GetTicketsByCountry() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		destination := ctx.Param("dest")

		tickets, err := ct.service.GetTotalTickets(ctx, destination)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": tickets})
	}
}

//Average destination
func (s *ControllerTicket) AverageDestination() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		destination := ctx.Param("dest")

		avg, err := s.service.AverageDestination(ctx, destination)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": avg})
	}
}
