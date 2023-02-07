package apihelper

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
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
	// var partner_users map[string]interface{}
	var partner_users models.RubyClientOutput
	obj, _ := RubyClient(body, rubyclient)
	// fmt.Println(obj)
	// var new models.RubyClientOutput
	// err := json.Unmarshal([]byte(obj), &partner_users)

	bodyString := string(obj)

	err := json.Unmarshal([]byte(bodyString), &partner_users)

	// for k, v := range partner_users {
	// 	// fmt.Print("-------------------------------------------------------------------------------------------------------")
	// 	// fmt.Print(k, v)
	// 	switch c := v.(type) {
	// 	default:
	// 		// fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
	// 	}
	// }
	if err != nil {
		fmt.Println(err, "Error occured")
	}

	fmt.Println(partner_users)

	return auth_role.StakeholderId

}

// rest client leke ruby ki api call krna h. // done
// Listpartner users - rm mapping required false. // done
// users apollo3 - incident client ka user banana h in navigation bar.
// ruby client similar to incident client.
// role id -> partner users list -> get active users from ticket reviewers -> active tickets jiska lowest hoga unko bhejna h
