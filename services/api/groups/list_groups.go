package api

import (
	"strings"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListGroup(filters models.FilterGroup) ([]models.GroupWithMember, *gorm.DB, error) {
	db := config.GetDB()

	var err error

	var groups []models.GroupWithMember

	db = db.Model(&models.Group{})

	db = db.Select("groups.id, groups.name,groups.status,groups.tags,Count( group_members.id) as count")

	db = db.Joins("left join group_members on group_members.group_id = groups.id and group_members.status = ?", "active")

	if filters.GroupMemberID > 0 {
		db = db.Where("group_members.id = ?", filters.GroupMemberID)
	}

	if filters.Name != "" {
		filters.Name = "%" + filters.Name + "%"
		db = db.Where("name iLike ?", filters.Name)
	}

	if len(filters.Tags) != 0 {
		db = db.Where("groups.tags && ?", "{"+strings.Join(filters.Tags, ",")+"}")
	}

	if filters.Status != "" {
		db = db.Where("groups.status = ?", filters.Status)
	}

	db = db.Order("groups.name desc")

	db = db.Group("1,2,3,4")

	db = db.Scan(&groups)


	return groups, db, err
}
