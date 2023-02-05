package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	// validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

type TokenUserService struct {
	TokenUser models.TicketUser
}

func CreateTokenUser(token_user models.TicketUser) (models.TicketUser, error) {
	db := config.GetDB()

	token_user.Status = "active"
	var exist_user models.TicketUser
	var err error
	db.Where("system_user_id = ? and status = ?", token_user.SystemUserID, "active").First(&exist_user)

	if exist_user.ID <= 0 {

		// stmt := validations.ValidateTokenUser(token_user)

		// if stmt != "validated" {
		// 	return token_user, errors.New(stmt)
		// }
		tx := db.Begin()
		if err := tx.Create(&token_user).Error; err != nil {
			tx.Rollback()
			return token_user, errors.New(err.Error())
		}
		tx.Commit()
		return token_user, err

	} else {

		return exist_user, err
	}
}
