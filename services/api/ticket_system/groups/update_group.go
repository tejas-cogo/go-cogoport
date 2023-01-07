package ticket_system

import (
	"fmt"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateGroup(id uint, body models.Group) models.Group {
	db := config.GetDB()
	var group models.Group
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&group)

	if (body.Name != group.Name){
		group.Name = body.Name
	} 
	// if (body.Tags != nil){
	// 	group.Tags = body.Tags
	// } 
	if (body.Status != group.Status){
		group.Status = body.Status
	} 

	db.Save(&group)
	return group
}