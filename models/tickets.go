package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/dariubs/gorm-jsonb"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	TicketUserID uint
	Source       string
	Type         string
	Category     string
	Subcategory  string
	Description  string
	Priority     string
	Tags         pq.StringArray `gorm:"type:text[]"`
	Data         gormjsonb.JSONB
	//pgtype.JSONB    `gorm:"type:json;default:'[];not null'`
	NotificationPreferences pq.StringArray `gorm:"type:text[]"`
	Tat                     time.Duration  `gorm:"type:string"`
	ExpiryDate              time.Time
	Status                  string
}

// type InvoiceDetail struct {
// 	// DataInfoID    uint
// 	InvoiceNumber string
// 	InvoiceUrl    string
// }

// type DataInfo struct {
// 	// invoice invoice_detail `gorm:"type:json;default:{}"`
// 	InvoiceNumber string
// 	// Invoice InvoiceDetail `gorm:"type:json"`
// 	// Priority string
// 	// payment    pq.StringArray `gorm:"type:text[]"`
// 	// attachment pq.StringArray `gorm:"type:text[]"`
// }

type MagicArray []interface{}

func (ma *MagicArray) UnmarshalJSON(b []byte) error {
	if b[0] == '[' {
		return json.Unmarshal(b, (*[]interface{})(ma))
	} else {
		var obj interface{}
		if err := json.Unmarshal(b, &obj); err != nil {
			return err
		}
		*ma = append(*ma, obj)
	}
	return nil

}

type JSONB []interface{}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
