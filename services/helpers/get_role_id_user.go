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
	var body models.Body

	rubyclient.Endpoint = "partner/list_partner_users"
	body.Filters.RoleIDs = append(body.Filters.RoleIDs, RoleID.String())
	body.Filters.Status = "active"
	body.RmMappingDataRequired = false

	var partner_users models.RestClientOutput
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

	max := 0
	var users []string
	var user_id uuid.UUID

	db.Model(&ticket_reviewer).Where("role_id = ? and status = ?", RoleID, "active").Distinct("user_id").Pluck("user_id", &users)

	if len(users) < len(user_id_array) {
		for _, value := range user_id_array {
			if !Inslice(value, users) {
				user_id, err = uuid.Parse(user_id_array[0])
			}
		}
	} else {
		db.Model(&ticket_reviewer).Where("user_id IN (?) and status = ?", user_id_array, "active").Select("Count(Distinct(ticket_id)) as count,user_id as user_id").Group("user_id").Order("count desc").Scan(&result)
		for _, value := range result {
			if value.Count >= max {
				max = value.Count
				user_id, err = uuid.Parse(value.UserID)
			}
		}

	}

	return user_id
}
