package ticket_system

import (
	"errors"
	"strings"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListGroup(filters models.FilterGroup) ([]models.GroupWithMember, *gorm.DB, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var groups []models.GroupWithMember

	tx = tx.Model(&models.Group{})

	tx = tx.Select("groups.id, groups.name,groups.status,groups.tags,Count( group_members.id) as count")

	tx = tx.Joins("left join group_members on group_members.group_id = groups.id and group_members.status = ?", "active")

	if filters.GroupMemberID > 0 {
		tx = tx.Where("group_members.id = ?", filters.GroupMemberID)
	}

	if filters.Name != "" {
		filters.Name = "%" + filters.Name + "%"
		tx = tx.Where("name iLike ?", filters.Name)
	}

	if len(filters.Tags) != 0 {
		tx = tx.Where("groups.tags && ?", "{"+strings.Join(filters.Tags, ",")+"}")
	}

	if filters.Status != "" {
		tx = tx.Where("groups.status = ?", filters.Status)
	}



	tx = tx.Order("groups.name desc")


	tx = tx.Group("1,2,3,4")

	tx = tx.Scan(&groups)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return groups, tx, errors.New("Error Occurred!")
	}

	tx.Commit()
	return groups, tx, err
}
