package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/laurianderson/bootcamp_go_desafio_go_web/cmd/api/server"
	"github.com/laurianderson/bootcamp_go_desafio_go_web/internal/domain"
)

func main() {

	//Load tickets.csv
	list, err := LoadTicketsFromFile("/Users/landerson/Documents/meli_bootcamp/go/bootcamp_go_desafio_go_web/tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	//Ping
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	router := server.NewRouter(r, &list)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}

}

//Open tickets.csv file
func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
