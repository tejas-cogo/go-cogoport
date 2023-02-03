package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

type TicketDefaultGroupService struct {
	TicketDefaultGroup models.TicketDefaultGroup
}

func CreateTicketDefaultGroup(ticket_default_group models.TicketDefaultGroup) (models.TicketDefaultGroup,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error
	var existed_default_group models.TicketDefaultGroup

	stmt := validations.ValidateTicketDefaultGroup(ticket_default_group)
	if stmt != "validated" {
		return ticket_default_group, errors.New(stmt)
	}
	
	ticket_default_group.Status = "active"

	tx.Where("ticket_default_type_id = ? and group_id = ? and group_member_id = ? and level = ? and status = ?", ticket_default_group.TicketDefaultTypeID, ticket_default_group.GroupID, ticket_default_group.GroupMemberID, ticket_default_group.Level, "active").First(&existed_default_group)

	if existed_default_group.ID > 0 {
		if err := tx.Model(&ticket_default_group).Where("id = ?", existed_default_group.ID).Update("status","inactive").Error; err != nil {
			tx.Rollback()
			return ticket_default_group, errors.New(err.Error())
		}

		if err := tx.Where("id = ?", existed_default_group.ID).Delete(&ticket_default_group).Error; err != nil {
			tx.Rollback()
			return ticket_default_group, errors.New(err.Error())
		}
	} 

	if err := tx.Create(&ticket_default_group).Error; err != nil {
		tx.Rollback()
		return ticket_default_group, errors.New(err.Error())
	}

	tx.Commit()

	return ticket_default_group, err
}
