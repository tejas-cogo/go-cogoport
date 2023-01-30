package ticket_system

import (
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketUser(filters models.TicketUserFilter) ([]models.TicketUser, *gorm.DB) {
	db := config.GetDB()

	var ticket_user []models.TicketUser

	if filters.GroupUnassigned == true {
		var ticket_users []uint
		db.Model(&models.GroupMember{}).Where("status = ?", "active").Where("group_name != ?", "Default").Distinct("ticket_user_id").Pluck("TicketUserId", &ticket_users)
		fmt.Println("dnck", ticket_users)
		if len(ticket_users) != 0 {
			db = db.Not("id IN ?", ticket_users)
		}
	}

	if filters.ID != 0 {
		db = db.Where("id = ?", filters.ID)
	}

	if filters.NotPresentID != 0 {
		db = db.Where("id != ?", filters.NotPresentID)
	}

	if filters.SystemUserID != "" {
		db = db.Where("system_user_id = ?", filters.SystemUserID)
	}

	if filters.Name != "" {
		filters.Name = "%" + filters.Name + "%"
		db = db.Where("name iLIKE ?", filters.Name)
	}

	if filters.Email != "" {
		filters.Email = "%" + filters.Email + "%"
		db = db.Where("email iLIKE ?", filters.Email)
	}

	if filters.MobileNumber != "" {
		db = db.Where("mobile_number = ?", filters.MobileNumber)
	}

	if filters.Type != "" {
		db = db.Where("type = ?", filters.Type)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	if filters.RoleID > 1 {
		db = db.Where("role_id = ?", filters.RoleID)
	}

	if filters.RoleUnassigned == true {
		db = db.Where("role_id = ?", 1)
	}

	if filters.AgentRmID != "" {

		db2 := config.GetCDB()
		var partner_user_rm_mapping []models.PartnerUserRmMapping
		var partner_user_rm_ids []string
		var ticket_reviewer []models.TicketReviewer
		var ticket_user []models.TicketUser
		var ticket_id []uint
		var ticket_users []uint

		db2.Where("reporting_manager_id = ? and status = ?", filters.AgentRmID, "active").Distinct("user_id").Find(&partner_user_rm_mapping).Pluck("user_id", &partner_user_rm_ids)
		fmt.Println("partner_user_rm_ids", partner_user_rm_ids)

		db.Where("system_user_id IN ?", partner_user_rm_ids).Distinct("id").Find(&ticket_user).Pluck("id", &ticket_users)

		db.Where("ticket_user_id IN ?", ticket_users).Distinct("ticket_id").Order("ticket_id").Find(&ticket_reviewer).Pluck("ticket_id", &ticket_id)

	}

	db = db.Preload("Role").Find(&ticket_user)

	return ticket_user, db
}
