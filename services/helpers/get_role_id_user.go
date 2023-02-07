package apihelper

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetRoleIdUser(RoleID uuid.UUID) uuid.UUID {
	var auth_role models.AuthRole
	var rubyclient models.RubyClientInput
	var body models.PartnerUserBody

	rubyclient.Endpoint = "partner/list_partner_users"
	body.Filters.RoleIDs = append(body.Filters.RoleIDs, RoleID.String())
	body.Filters.Status = "active"
	body.RmMappingDataRequired = false
	body.PageLimit = 1

	var partner_users models.RubyClientOutput
	obj, _ := RubyClient(body, rubyclient)

	bodyString := string(obj)

	err := json.Unmarshal([]byte(bodyString), &partner_users)
	if err != nil {
		fmt.Println(err, "Error occured")
	}

	var user_id []string
	for _, user_details := range partner_users.List {
		user_id = append(user_id, user_details.UserID)
	}

	var ticket_reviewer models.TicketReviewer

	type Result struct {
		UserID string
		Count  int
	}
	// var max models.PartnerUserList

	var result Result

	db := config.GetDB()
	db.Model(&ticket_reviewer).Where("user_id = ? and status = ?", user_id, "active").Select("Count(Distinct(ticket_id)) as count,user_id as user_id").Group("user_id").Scan(&result)

	fmt.Println("----------------------")
	fmt.Println(result)
	return auth_role.StakeholderId

}

// rest client leke ruby ki api call krna h. // done
// Listpartner users - rm mapping required false. // done
// users apollo3 - incident client ka user banana h in navigation bar.
// ruby client similar to incident client.
// role id -> partner users list -> get active users from ticket reviewers -> active tickets jiska lowest hoga unko bhejna h
