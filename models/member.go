package models

import (
	"time"

	"../db"
)

type Member struct {
	ID        uint       `json:"id"`
	FullName  string     `sql:"type:varchar(255);not null" json:"full_name"`
	Username  string     `sql:"type:varchar(50);not null;unique" json:"username"`
	GroupID   int        `sql:"type:int(10)" json:"group_id"`
	Email     string     `sql:"type:varchar(50);not null;unique" json:"email"`
	GroupID   int        `sql:"type:int(10)" json:"group_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type MemberModel struct{}

func (m *MemberModel) Find(list *[]Member) (err error) {
	if err := db.DB.Find(list).Error; err != nil {
		return err
	}
	return nil
}

func (m *MemberModel) Get(id string, member *Member) (err error) {
	if err := db.DB.First(member, id).Error; err != nil {
		return err
	}
	return nil
}

func (m *MemberModel) Create(member *Member) (err error) {
	if err := db.DB.Create(member).Error; err != nil {
		return err
	}
	return nil
}

func (m *MemberModel) GetMembers(groupID int, list *[]Member) (err error) {
	if err := db.DB.Where("group_id = ?", groupID).Find(list).Error; err != nil {
		return err
	}
	return nil
}

func (m *MemberModel) AddToGroup(id int, groupID int, member *Member) (err error) {
	if err := db.DB.First(member, id).Error; err != nil {
		return err
	}
	if err := db.DB.Model(&member).Update("group_id", groupID).Error; err != nil {
		return err
	}
	return nil
}

func (m *MemberModel) JoinGroup(id int, groupCode string, member *Member) (err error) {
	var rGroup RGroup

	//Verify member is not in group yet

	if err := db.DB.First(member, id).Error; err != nil {
		return err
	}
	if err := db.DB.Where("code = ?", groupCode).First(&rGroup).Error; err != nil {
		return err
	}
	if err := db.DB.Model(&member).Update("group_id", int(rGroup.ID)).Error; err != nil {
		return err
	}
	if err := db.DB.Model(&rGroup).Update("size", rGroup.Size+1).Error; err != nil {
		return err
	}

	return nil
}

/*func (m *MemberModel) Update(id string, data Member) (member Member, err error) {
	if err := db.DB.First(&member, id).Error; err != nil {
		return member, nil
	}
	if err := db.DB.Save(&member).Error; err != nil {
		return member, err
	}
	return member, nil
}

func (m *MemberModel) Delete(id string, data Member) (member Member, err error) {
	if err := db.DB.First(&member, id).Error; err != nil {
		return member, err
	}
	if err := db.DB.Delete(&member).Error; err != nil {
		return member, err
	}
	return member, nil
}*/
