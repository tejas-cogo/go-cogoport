package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
)

func ListTicket(c *gin.Context) {
	var filters models.Ticket
	var sort models.Sort

	ID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[id]"))
	filters.ID = uint(ID)

	filters.Source = c.Request.URL.Query().Get("filters[source]")

	filters.Type = c.Request.URL.Query().Get("filters[type]")

	filters.Priority = c.Request.URL.Query().Get("filters[priority]")

	filters.CreatedAt, _ = time.Parse("2006-01-02", c.Request.URL.Query().Get("filters[created_at]"))

	filters.ExpiryDate, _ = time.Parse("2006-01-02", c.Request.URL.Query().Get("filters[expiry_date]"))

	filters.Status = c.Request.URL.Query().Get("filters[status]")

	sort.SortBy = c.Request.URL.Query().Get("sort_by")
	sort.SortType = c.Request.URL.Query().Get("sort_type")

	// filters.Tags[0] = c.Request.URL.Query().Get("filters[tags]")
	// c.JSON(200, pg.Response(model, c.Request, &[]Article{}))
	ser, db := service.ListTicket(filters, sort)
	pg := paginate.New() 
	fmt.Println("count", db.RowsAffected, "count")
	c.JSON(200, pg.Response(db, c.Request, &ser))
}

func GetTicketStats(c *gin.Context) {
	var stats models.TicketStat

	id := c.Request.URL.Query().Get("filters[agent_id]")
	rm_id := c.Request.URL.Query().Get("filters[agent_rm_id]")

	if id != "" {
		stats.AgentID, _ = uuid.Parse(id)
	}
	if rm_id != "" {
		stats.AgentRmID, _ = uuid.Parse(rm_id)
	}

	fmt.Println("rm", stats.AgentRmID, "rm")

	c.JSON(200, service.GetTicketStats(stats))
}

func ListTicketDetail(c *gin.Context) {
	var filters models.TicketDetail

	ID, _ := strconv.Atoi(c.Request.URL.Query().Get("filters[id]"))
	filters.TicketID = uint(ID)

	c.JSON(200, service.ListTicketDetail(filters))
}

func CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	c.BindJSON(&ticket)
	c.JSON(200, service.CreateTicket(ticket))
}

// func DeleteTicket(c *gin.Context) {
// 	var body models.Ticket
// 	c.BindJSON(&body)
// 	id := body.ID
// 	c.JSON(200, service.DeleteTicket(id))
// }

func UpdateTicket(c *gin.Context) {
	var body models.Ticket
	c.BindJSON(&body)
	c.JSON(200, service.UpdateTicket(body))
}
