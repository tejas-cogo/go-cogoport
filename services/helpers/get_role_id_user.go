package apihelper

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetRoleIdUser(RoleID uuid.UUID, UserID uuid.UUID) uuid.UUID {
	var rubyclient models.RubyClientInput
	var body models.Body

	rubyclient.Endpoint = "partner/list_partner_users"
	body.Filters.RoleIDs = append(body.Filters.RoleIDs, RoleID.String())
	body.Filters.Status = "active"
	body.RmMappingDataRequired = false

	var partner_users models.RestClientOutput
	obj, _ := GetRubyClient(body, rubyclient)

	bodyString := string(obj)

	err := json.Unmarshal([]byte(bodyString), &partner_users)
	if err != nil {
		log.Println(err)
	}

	var user_id_array []string

	db := config.GetDB()
	var ticket models.Ticket
	db.Where("id = ? and status = ?", RoleID, "active").First(&ticket)

	for _, user_details := range partner_users.List {
		if user_details.UserID != ticket.UserID.String() {
			user_id_array = append(user_id_array, user_details.UserID)
		}
	}

	var ticket_reviewer models.TicketReviewer

	// var max models.PartnerUserList

	var users []string

	db.Model(&ticket_reviewer).Where("role_id = ? and status = ?", RoleID, "active").Distinct("user_id").Pluck("user_id", &users)

	if UserID != uuid.Nil {
		if Inslice(UserID.String(), users) {
			users = Remove(users, UserID.String())
		}
	}

	user_id := GetFilteredUser(users, user_id_array)

	return user_id
}

func GetUnifiedRoleIdUser(RoleID uuid.UUID, UserID string) uuid.UUID {

	db2 := config.GetCDB()
	db := config.GetDB()

	var user_id_array []string
	var ticket_reviewer models.TicketReviewer

	db2.Model(&models.PartnerUser{}).Where("role_ids && ? and status = ?", "{"+RoleID.String()+"}", "active").Distinct("user_id").Pluck("user_id", &user_id_array)

	var users []string

	db.Model(&ticket_reviewer).Where("role_id = ? and status = ?", RoleID, "active").Distinct("user_id").Pluck("user_id", &users)

	if users != nil {

		if UserID != "" {
			if Inslice(UserID, users) {
				users = Remove(users, UserID)
			}
		}

		if len(user_id_array) > 0 {
			user_id := GetFilteredUser(users, user_id_array)

			if user_id != uuid.Nil {
				return user_id
			}
		}

	}

	return uuid.Nil
}

func GetFilteredUser(users []string, user_id_array []string) uuid.UUID {

	var err error

	db := config.GetDB()
	var ticket_reviewer models.TicketReviewer

	max := 0

	type Result struct {
		UserID string `json:"user_id"`
		Count  int
	}
	var result []Result

	var user_id uuid.UUID

	if len(users) < len(user_id_array) {
		for _, value := range user_id_array {
			if !Inslice(value, users) {
				user_id, err = uuid.Parse(user_id_array[0])
				if err != nil {
					log.Print(err)
				}
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

	if user_id != uuid.Nil {
		return user_id
	} else {
		return uuid.Nil
	}

}

func Remove(array []string, str string) []string {
	for i, u := range array {
		if u == str && i != len(array)-1 {
			array[i] = array[len(array)-1]
		}
	}
	return array[:len(array)-1]
}
