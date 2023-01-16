package models

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"gorm.io/gorm"
)

var db *gorm.DB

//Model is sample of common table structure

func Init() {

	db := config.GetDB()
	//Printing query
	// db.LogMode(true)

	//Automatically create migration as per model

	// db.Migrator().DropTable(&Group{},&Role{},&TicketUser{},&GroupMember{},&TicketDefaultGroup{},&TicketDefaultTiming{},&TicketDefaultType{},&Ticket{},&TicketActivity{},&TicketReviewer{},&TicketSpectator{},&TicketTask{},&TicketTaskAssignee{},&TicketAudit{})

	// db.Migrator().CreateTable(&TicketToken{})

	// db.Migrator().CreateTable(&Group{},&Role{},&TicketUser{},&GroupMember{},&TicketDefaultGroup{},&TicketDefaultTiming{},&TicketDefaultType{},&Ticket{},&TicketActivity{},&TicketReviewer{},&TicketSpectator{},&TicketTask{},&TicketTaskAssignee{},&TicketAudit{})

	db.Migrator().AutoMigrate(&TicketUser{})

	// db.Migrator().AutoMigrate(&Group{}, &Role{}, &TicketUser{}, &TicketDefaultGroup{}, &GroupMember{}, &TicketDefaultTiming{}, &TicketDefaultType{}, &Ticket{}, &TicketActivity{}, &TicketReviewer{}, &TicketSpectator{}, &TicketTask{}, &TicketTaskAssignee{}, &TicketAudit{})
}

// GetDB function return the instance of db
