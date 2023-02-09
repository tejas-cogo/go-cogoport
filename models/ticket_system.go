package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
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
	ManagerID uuid.UUID
	RoleIDs   pq.StringArray `gorm:"type:text[]"`
	Status    string
}
type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	MobileNumber string
}
type UserData struct {
	ID           uuid.UUID
	Name         string `json:"name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
}

type Filter struct {
	Ticket              Ticket
	TicketUser          TicketUser
	TicketActivity      TicketActivity
	Activity            Activity
	TicketAudit         TicketAudit
	TicketDefaultRole   TicketDefaultRole
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

	// db.Migrator().CreateTable(&TicketDefaultRole{})

	// db.Migrator().CreateTable(&Group{},&Role{},&TicketUser{},&GroupMember{},&TicketDefaultGroup{},&TicketDefaultTiming{},&TicketDefaultType{},&Ticket{},&TicketActivity{},&TicketReviewer{},&TicketSpectator{},&TicketTask{},&TicketTaskAssignee{},&TicketAudit{})

	db.Migrator().AutoMigrate(&TicketDefaultType{})

	// db.Migrator().AutoMigrate(&TicketDefaultRole{}, &TicketDefaultTiming{}, &TicketDefaultType{}, &Ticket{}, &TicketActivity{}, &TicketReviewer{}, &TicketSpectator{}, &TicketTask{}, &TicketTaskAssignee{}, &TicketAudit{}, &TicketTask{}, &TicketTaskAssignee{})
}

// GetDB function return the instance of db
