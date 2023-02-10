package apihelper

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetUserData(IDs pq.StringArray) []models.UserData {
	var rubyclient models.RubyClientInput

	type Filters struct {
		ID pq.StringArray `json:"id"`
	}

	type Body struct {
		Filters Filters `json:"filters"`
	}

	type Response struct {
		List       []models.UserData
		Page       uint `json:"page"`
		Total      uint `json:"total"`
		TotalCount uint `json:"total_count"`
		PageLimit  uint `json:"page_limit"`
	}

	var user Response
	var body Body

	rubyclient.Endpoint = "user/list_users"
	body.Filters.ID = IDs
	obj, _ := GetRubyClient(body, rubyclient)

	bodyString := string(obj)

	err := json.Unmarshal([]byte(bodyString), &user)
	if err != nil {
		log.Println(err)
	}

	var users []models.UserData
	for _, user_details := range user.List {
		users = append(users, user_details)
	}
	return users
}

func GetUnifiedUserData(IDs pq.StringArray) []models.UserData {

	db2 := config.GetCDB()

	var user_data []models.UserData

	var user_ids []uuid.UUID

	for _, u := range IDs {
		data, err := uuid.Parse(u)
		if err != nil {
			log.Println(err)
		}
		user_ids = append(user_ids, data)
	}

	db2.Model(&models.User{}).Where("id in (?)", user_ids).Find(&user_data)

	return user_data
}
