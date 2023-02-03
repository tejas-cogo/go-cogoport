package migrations

import (
	"github.com/tejas-cogo/go-cogoport/config"
	models "github.com/tejas-cogo/go-cogoport/models"
)

func Init() {
	// var models = []interface{}{&Group{}, &Product{}, &Order{}}

	// db.Migrator().Automigrate(models...)

	db := config.GetDB()
	//Printing query
	// db.LogMode(true)

	//Automatically create migration as per model

	// db.Migrator().DropTable(&Group{},&Role{},&TicketUser{},&GroupMember{},&TicketDefaultGroup{},&TicketDefaultTiming{},&TicketDefaultType{},&Ticket{},&TicketActivity{},&TicketReviewer{},&TicketSpectator{},&TicketTask{},&TicketTaskAssignee{},&TicketAudit{})

	// db.Migrator().CreateTable(&TicketToken{})

	// db.Migrator().CreateTable(&Group{},&Role{},&TicketUser{},&GroupMember{},&TicketDefaultGroup{},&TicketDefaultTiming{},&TicketDefaultType{},&Ticket{},&TicketActivity{},&TicketReviewer{},&TicketSpectator{},&TicketTask{},&TicketTaskAssignee{},&TicketAudit{})

	// db.Migrator().AutoMigrate(&Ticket{})

	db.Migrator().AutoMigrate(&models.Group{}, &models.Role{}, &models.TicketUser{}, &models.TicketDefaultGroup{}, &models.GroupMember{}, &models.TicketDefaultTiming{}, &models.TicketDefaultType{}, &models.Ticket{}, &models.TicketActivity{}, &models.TicketReviewer{}, &models.TicketSpectator{}, &models.TicketTask{}, &models.TicketTaskAssignee{}, &models.TicketAudit{})
}

// GetDB function return the instance of db
