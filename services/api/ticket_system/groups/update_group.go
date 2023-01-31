package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func UpdateGroup(body models.Group) (models.Group,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var group models.Group

	if body.ID != 0 {
		tx.Where("id = ?", body.ID).Find(&group)
	}

	if body.Name != "" {
		group.Name = body.Name
	}
	if body.Tags != nil {
		group.Tags = body.Tags
	}
	if body.Status != "" {
		group.Status = body.Status
	}

	if err := tx.Save(&group).Error; err != nil {
		tx.Rollback()
		return body, errors.New("Error Occurred!")
	}

	tx.Commit()
	
	return group, err
}
