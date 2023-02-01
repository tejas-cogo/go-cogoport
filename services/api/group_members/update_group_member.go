package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func UpdateGroupMember(body models.GroupMember) (models.GroupMember,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var group_member models.GroupMember

	if body.ID != 0 {
		tx.Where("id = ?", body.ID)
	}

	if err := tx.Find(&group_member).Error; err != nil {
		tx.Rollback()
		return body, errors.New("Group member not found!")
	}

	if body.ActiveTicketCount != 0 {
		group_member.ActiveTicketCount = body.ActiveTicketCount
	}
	if body.HierarchyLevel != 0 {
		group_member.HierarchyLevel = body.HierarchyLevel
	}
	if body.Status != "" {
		group_member.Status = body.Status
	}
	if body.TicketUserID != 0 {
		group_member.HierarchyLevel = body.HierarchyLevel
	}

	if err := tx.Save(&group_member).Error; err != nil {
		tx.Rollback()
		return body, errors.New("Can't save Group member details!")
	}

	tx.Commit()
	
	return group_member, err
}
