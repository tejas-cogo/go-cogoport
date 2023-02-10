package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/morkid/paginate"
	models "github.com/tejas-cogo/go-cogoport/models"

	role_service "github.com/tejas-cogo/go-cogoport/services/api/ticket_default_roles"
	service "github.com/tejas-cogo/go-cogoport/services/api/ticket_default_types"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

func ListTicketType(c *gin.Context) {
	var filters models.TicketDefaultType

	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}

	ser, db := service.ListTicketType(filters)
	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		pg := paginate.New()
		c.JSON(c.Writer.Status(), pg.Response(db, c.Request, &ser))
	}
}

func ListTicketDefaultType(c *gin.Context) {
	var filters models.TicketDefaultFilter

	err := c.Bind(&filters)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}

	ser, db := service.ListTicketDefaultType(filters)

	if c.Writer.Status() == 400 {
		c.JSON(c.Writer.Status(), "Not Found")
	} else {
		data := paginate.New().With(db).Request(c.Request).Response(&ser)
		items, _ := json.Marshal(data.Items)
		var output []models.TicketDefault

		err := json.Unmarshal([]byte(items), &output)
		if err != nil {
			print(err)
			c.JSON(400, err)
		}

		var users []string
		var roles []string
		var closures []string

		for j := 0; j < len(output); j++ {
			var f models.TicketDefaultRole
			f.TicketDefaultTypeID = output[j].ID
			output[j].TicketDefaultRole, _ = role_service.ListTicketDefaultRole(f)

			for i := 0; i < len(output[j].TicketDefaultRole); i++ {

				if output[j].TicketDefaultRole[i].UserID != uuid.Nil {
					users = append(users, output[j].TicketDefaultRole[i].UserID.String())
				} else {
					roles = append(roles, output[j].TicketDefaultRole[i].RoleID.String())
				}
			}

			for i := 0; i < len(output[j].ClosureAuthorizer); i++ {
				if output[j].ClosureAuthorizer[i] != "" {
					if !helpers.Inslice(output[j].ClosureAuthorizer[i], closures) {
						closures = append(closures, output[j].ClosureAuthorizer[i])
					}
				}
			}
		}

		user_data := helpers.GetUnifiedUserData(users)
		closure_data := helpers.GetUnifiedUserData(closures)
		role_data := helpers.GetUnifiedAuthRoleData(roles)

		for j := 0; j < len(output); j++ {
			var f models.TicketDefaultRole
			f.TicketDefaultTypeID = output[j].ID
			output[j].TicketDefaultRole, _ = role_service.ListTicketDefaultRole(f)

			if len(output[j].ClosureAuthorizer) != 0 {
				for i := 0; i < len(output[j].ClosureAuthorizer); i++ {

					if output[j].ClosureAuthorizer[i] != "" {

						fmt.Println("here ")

						for k := 0; k < len(closure_data); k++ {
							fmt.Println("here ", output[j].ClosureAuthorizer[i])
							fmt.Println("here kjdehc", closure_data[k])

							if output[j].ClosureAuthorizer[i] == closure_data[k].ID.String() {

								var user models.User
								user.ID = closure_data[k].ID
								user.Name = closure_data[k].Name
								user.Email = closure_data[k].Email
								user.MobileNumber = closure_data[k].MobileNumber

								output[j].ClosureAuthorizerData = append(output[j].ClosureAuthorizerData, user)
								break

							}
						}
					}
				}
			}

			for i := 0; i < len(output[j].TicketDefaultRole); i++ {

				if output[j].TicketDefaultRole[i].UserID != uuid.Nil {
					for k := 0; k < len(user_data); k++ {
						if user_data[k].ID == output[j].TicketDefaultRole[i].UserID {
							output[j].TicketDefaultRole[i].User.ID = user_data[k].ID
							output[j].TicketDefaultRole[i].User.Name = user_data[k].Name
							output[j].TicketDefaultRole[i].User.Email = user_data[k].Email
							output[j].TicketDefaultRole[i].User.MobileNumber = user_data[k].MobileNumber
							break
						}
					}
				} else {
					for k := 0; k < len(role_data); k++ {
						if role_data[k].ID == output[j].TicketDefaultRole[i].RoleID {
							output[j].TicketDefaultRole[i].Role.ID = role_data[k].ID
							output[j].TicketDefaultRole[i].Role.Name = role_data[k].Name
							output[j].TicketDefaultRole[i].Role.StakeholderID = role_data[k].StakeholderID
							output[j].TicketDefaultRole[i].Role.Status = role_data[k].Status
							break
						}
					}
				}
			}
		}

		data.Items = output

		c.JSON(c.Writer.Status(), data)
	}
}

func CreateTicketDefaultType(c *gin.Context) {
	var ticket_default_type models.TicketDefaultType
	err := c.Bind(&ticket_default_type)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	ser, err := service.CreateTicketDefaultType(ticket_default_type)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func DeleteTicketDefaultType(c *gin.Context) {
	var body models.TicketDefaultType
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	id := body.ID
	ser, err := service.DeleteTicketDefaultType(id)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}

func UpdateTicketDefaultType(c *gin.Context) {
	var body models.TicketDefaultType
	err := c.Bind(&body)
	if err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}
	ser, err := service.UpdateTicketDefaultType(body)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(c.Writer.Status(), ser)
	}
}
