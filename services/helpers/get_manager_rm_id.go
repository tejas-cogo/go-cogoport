package apihelper

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetManagerRmId(UserID uuid.UUID) []string {
	var rubyclient models.RubyClientInput
	var body models.Body
	var partner_users_rm_mapping models.RestClientOutput

	rubyclient.Endpoint = "partner/list_partner_user_rm_mappings"
	body.Filters.UserID = UserID.String()
	body.Filters.Status = "active"

	obj, _ := GetRubyClient(body, rubyclient)

	bodyString := string(obj)

	err := json.Unmarshal([]byte(bodyString), &partner_users_rm_mapping)
	if err != nil {
		log.Println(err)
	}
	var user_id_array []string
	for _, user_details := range partner_users_rm_mapping.List {
		user_id_array = append(user_id_array, user_details.ReportingManagerID)
	}

	return user_id_array
}

func GetUnifiedManagerRmId(UserID uuid.UUID) []string {

	db2 := config.GetCDB()

	var user_id_array []string

	db2.Model(&models.PartnerUser{}).Where("user_id In (?) and status = ?", UserID, "active").Distinct("manager_id").Pluck("manager_id", &user_id_array)

	return user_id_array
}
