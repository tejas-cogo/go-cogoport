package models

import (
	"gorm.io/gorm"
	"github.com/tejas-cogo/go-cogoport/config"
)

var db *gorm.DB

//Model is sample of common table structure

func Init() {

	db := config.GetDB()
	//Printing query
	// db.LogMode(true)

	//Automatically create migration as per model

	// db.Migrator().DropTable(&Group{},&Role{},&TicketUser{},&GroupMember{},&TicketDefaultGroup{},&TicketDefaultTiming{},&TicketDefaultType{},&Ticket{},&TicketActivity{},&TicketReviewer{},&TicketSpectator{},&TicketTask{},&TicketTaskAssignee{},&TicketAudit{})

	// db.Migrator().CreateTable(&Group{})

	// db.Migrator().CreateTable(&Group{},&Role{},&TicketUser{},&GroupMember{},&TicketDefaultGroup{},&TicketDefaultTiming{},&TicketDefaultType{},&Ticket{},&TicketActivity{},&TicketReviewer{},&TicketSpectator{},&TicketTask{},&TicketTaskAssignee{},&TicketAudit{})

	db.Migrator().AutoMigrate(&TicketDefaultTiming{})

	// db.Migrator().AutoMigrate(&Group{},&Role{},&TicketUser{},&GroupMember{},&TicketDefaultGroup{},&TicketDefaultTiming{},&TicketDefaultType{},&Ticket{},&TicketActivity{},&TicketReviewer{},&TicketSpectator{},&TicketTask{},&TicketTaskAssignee{},&TicketAudit{})
}

// GetDB function return the instance of db
