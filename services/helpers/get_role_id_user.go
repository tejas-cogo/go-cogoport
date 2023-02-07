package apihelper

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetRoleIdUser(RoleID uuid.UUID) uuid.UUID {
	var rubyclient models.RubyClientInput
	var body models.PartnerUserBody

	rubyclient.Endpoint = "partner/list_partner_users"
	body.Filters.RoleIDs = append(body.Filters.RoleIDs, RoleID.String())
	body.Filters.Status = "active"
	body.RmMappingDataRequired = false

	var partner_users models.RubyClientOutput
	obj, _ := RubyClient(body, rubyclient)

	bodyString := string(obj)

	err := json.Unmarshal([]byte(bodyString), &partner_users)
	if err != nil {
		fmt.Println(err, "Error occured")
	}

	var user_id_array []string
	for _, user_details := range partner_users.List {
		user_id_array = append(user_id_array, user_details.UserID)
	}

	var ticket_reviewer models.TicketReviewer

	type Result struct {
		UserID string `json:"user_id"`
		Count  int
	}
	// var max models.PartnerUserList

	var result []Result

	db := config.GetDB()
	db = db.Model(&ticket_reviewer).Where("user_id IN (?) and status = ?", user_id_array, "active").Select("Count(Distinct(ticket_id)) as count,user_id as user_id").Group("user_id").Scan(&result)

	max := 0
	var user_id uuid.UUID

	for _, value := range result {
		if value.Count >= max {
			max = value.Count
			user_id, err = uuid.Parse(value.UserID)
		}
	}

	// a, _ := uuid.Parse(user_id_array[0])
	// fmt.Println(a)
	fmt.Println("--------------------->",RoleID)
	// fmt.Println(user_id)

	// if user_id == "00000000-0000-0000-0000-000000000000"{
	return user_id
	// } else {
		// return a
	// }
}

// rest client leke ruby ki api call krna h. // done
// Listpartner users - rm mapping required false. // done
// users apollo3 - incident client ka user banana h in navigation bar.
// ruby client similar to incident client.
// role id -> partner users list -> get active users from ticket reviewers -> active tickets jiska lowest hoga unko bhejna h
