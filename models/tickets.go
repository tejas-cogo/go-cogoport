package models

import (
	"time"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	TicketUserId            uint   
	Source                  string 
	Type                    string 
	Category                string 
	Subcategory             string 
	Description             string
	Priority                string         
	Tags                    pq.StringArray `gorm:"type:[]string"`
	Data                    pq.StringArray `gorm:"json:[]string"`
	NotificationPreferences pq.StringArray `gorm:"json:[]string"`
	Tat                     time.Time
	ExpiryDate              time.Time
	Status                  string
}
