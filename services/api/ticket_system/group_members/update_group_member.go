package ticket_system

import (

	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func UpdateGroupMember(id uint, body models.GroupMember) models.GroupMember {
	db := config.GetDB()
	var group_member models.GroupMember

	db.Where("id = ?", id).First(&group_member)

	if (body.ActiveTicketCount != group_member.ActiveTicketCount){
		group_member.ActiveTicketCount = body.ActiveTicketCount
	} 
	if (body.HierarchyLevel != group_member.HierarchyLevel){
		group_member.HierarchyLevel = body.HierarchyLevel
	} 
	if (body.Status != group_member.Status){
		group_member.Status = body.Status
	} 

	db.Save(&group_member)
	return group_member
}