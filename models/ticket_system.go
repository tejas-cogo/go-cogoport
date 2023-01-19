package models

import (
	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type PartnerUserRmMapping struct {
	ID                 uuid.UUID
	PartnerID          uuid.UUID
	UserID             uuid.UUID
	ReportingManagerID uuid.UUID
	Status             string
}

type PartnerUser struct {
	ID        uuid.UUID
	PartnerID uuid.UUID
	UserID    uuid.UUID
	Status    string
}

type Filter struct {
	gorm.Model
	Ticket              Ticket
	TicketUser          TicketUser
	Group               Group
	GroupMember         GroupMember
	Role                Role
	TicketActivity      TicketActivity
	Activity            Activity
	TicketAudit         TicketAudit
	TicketDefaultGroup  TicketDefaultGroup
	TicketDefaultTiming TicketDefaultTiming
	TicketDefaultType   TicketDefaultType
	TicketReviewer      TicketReviewer
	TicketSpectator     TicketSpectator
	TicketTask          TicketTask
	TicketToken         TicketToken
}

//Model is sample of common table structure

func Init() {

	db := config.GetDB()
	//Printing query
	// db.LogMode(true)

	//Automatically create migration as per model

	// db.Migrator().DropTable(&Group{},&Role{},&TicketUser{},&GroupMember{},&TicketDefaultGroup{},&TicketDefaultTiming{},&TicketDefaultType{},&Ticket{},&TicketActivity{},&TicketReviewer{},&TicketSpectator{},&TicketTask{},&TicketTaskAssignee{},&TicketAudit{})

	// db.Migrator().CreateTable(&TicketToken{})

	// db.Migrator().CreateTable(&Group{},&Role{},&TicketUser{},&GroupMember{},&TicketDefaultGroup{},&TicketDefaultTiming{},&TicketDefaultType{},&Ticket{},&TicketActivity{},&TicketReviewer{},&TicketSpectator{},&TicketTask{},&TicketTaskAssignee{},&TicketAudit{})

	db.Migrator().AutoMigrate(&TicketDefaultGroup{})

	// db.Migrator().AutoMigrate(&Group{}, &Role{}, &TicketUser{}, &TicketDefaultGroup{}, &GroupMember{}, &TicketDefaultTiming{}, &TicketDefaultType{}, &Ticket{}, &TicketActivity{}, &TicketReviewer{}, &TicketSpectator{}, &TicketTask{}, &TicketTaskAssignee{}, &TicketAudit{})
}

// GetDB function return the instance of db
