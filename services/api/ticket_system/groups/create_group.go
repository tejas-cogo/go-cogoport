package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type GroupService struct {
	Group models.Group
}

func CreateGroup(group models.Group) models.Group {
	db := config.GetDB()
	//  result := map[string]interface{}{}
	db.Create(&group)
	return group
}
