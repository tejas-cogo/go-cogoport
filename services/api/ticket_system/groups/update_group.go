package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateGroup(body models.Group) models.Group {
	db := config.GetDB()
	var group models.Group
	db.Where("id = ?", body.ID).Find(&group)

	if body.Name != "" {
		group.Name = body.Name
	}
	if body.Tags != nil {
		group.Tags = body.Tags
	}
	if body.Status != "" {
		group.Status = body.Status
	}

	db.Save(&group)
	return group
}
