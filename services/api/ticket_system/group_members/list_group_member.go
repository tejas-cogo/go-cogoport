package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func ListGroupMember(filters models.GroupMember) []models.GroupMember{
	db := config.GetDB()

	var group_members []models.GroupMember

	result := map[string]interface{}{}

	if (filters.GroupId != 0){
		db = db.Where("group_id = ?", filters.GroupId)
	} 

	if (filters.Status != ""){
		db = db.Where("status = ?", filters.Status)
	}else{
		db = db.Where("status = ?", "active")
	} 

	db.Order("HierarchyLevel desc").Order("ActiveTicketCount asc")

	db.Find(&group_members).Take(&result)

	return group_members
}