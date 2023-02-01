package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
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

	stmt := validations.ValidateGroup(group)
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
