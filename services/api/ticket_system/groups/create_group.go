package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

type GroupService struct {
	Group models.Group
}

func CreateGroup(group models.Group) (models.Group,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	group.Status = "active"

	stmt := validate(group)
	if stmt != "validated" {
		return group, errors.New(stmt)
	}

	if err := tx.Create(&group).Error; err != nil {
		tx.Rollback()
		return group, errors.New("Error Occurred!")
	}

	tx.Commit()

	return group, err
}

func validate(group models.Group) string {
	if group.Name == "" {
		return ("Group Name Is Required!")
	}

	if len(group.Name) < 2 || len(group.Name) > 40 {
		return ("Name field must be between 2-40 chars!")
	}

	return ("validated")
}
