package apihelper

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetManagerRmId(UserID uuid.UUID) []string {
	var rubyclient models.RubyClientInput
	var body models.Body
	var partner_users_rm_mapping models.RestClientOutput

	rubyclient.Endpoint = "partner/list_partner_user_rm_mappings"
	body.Filters.UserID = UserID.String()
	body.Filters.Status = "active"

	obj, _ := RubyClient(body, rubyclient)

	bodyString := string(obj)

	err := json.Unmarshal([]byte(bodyString), &partner_users_rm_mapping)
	if err != nil {
		fmt.Println(err, "Error occured")
	}
	var user_id_array []string
	for _, user_details := range partner_users_rm_mapping.List {
		user_id_array = append(user_id_array, user_details.ReportingManagerID)
	}

	return user_id_array
}
