package server

import (
	"github.com/gin-gonic/gin"
	"github.com/laurianderson/bootcamp_go_desafio_go_web/cmd/api/handlers"
	"github.com/laurianderson/bootcamp_go_desafio_go_web/internal/domain"
	"github.com/laurianderson/bootcamp_go_desafio_go_web/internal/tickets"
)

//Struc
type Router struct {
	engine *gin.Engine
	bd     *[]domain.Ticket
}

//Builder
func NewRouter(engine *gin.Engine, tickets *[]domain.Ticket) *Router {
	return &Router{
        engine: engine,
        bd:     tickets,
    }
}

//Map routes
func (r *Router) MapRoutes() {
	r.engine.Use(gin.Recovery())
	r.engine.Use(gin.Logger())

	r.SetupRoutes()
}

//Set up routes
func (r *Router) SetupRoutes() {

	rp := tickets.NewRepository(*r.bd)
	s := tickets.NewService(rp)
	ct := handlers.NewControllerTicket(s)

	tckGroup := r.engine.Group("/tickets")

	tckGroup.GET("", ct.GetAll())
	tckGroup.GET("/getByCountry/:dest", ct.GetTicketsByCountry())
	tckGroup.GET("/getTotalByCountry/:dest", ct.GetTotalTicketsByCountry())
	tckGroup.GET("/getAverage/:dest", ct.AverageDestination())

}