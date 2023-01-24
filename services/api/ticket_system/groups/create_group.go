package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type GroupService struct {
	Group models.Group
}

func CreateGroup(group models.Group) (string, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	group.Status = "active"

	stmt := validate(group)
	if stmt != "validated" {
		return stmt, err
	}

	if err := tx.Create(&group).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err
	}

	tx.Commit()

	return "Successfully Created!", err
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
