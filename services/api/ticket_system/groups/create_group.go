package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type GroupService struct {
	Group models.Group
}

func CreateGroup(group models.Group) string {
	db := config.GetDB()
	//  result := map[string]interface{}{}
	group.Status = "active"
	db.Create(&group)
	return "Successfully created"
}
