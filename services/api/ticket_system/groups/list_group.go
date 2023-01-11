package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListGroup(filters models.Group,tags string) []models.Group{
	db := config.GetDB()

	var groups []models.Group

	if (filters.Name != ""){
		db = db.Where("name = ?", filters.Name)
	} 

	if (tags != ""){
		db = db.Where("? Like ANY(tags)", tags)
	} 

	if (filters.Status != ""){
		db = db.Where("status = ?", filters.Status)
	}else{
		db = db.Where("status = ?", "active")
	} 

	db.Find(&groups)

	return groups
}