package models

import "time"

type Member struct {
	ID 	        uint     	`json:"id"`
	FullName	string		`sql:"type:varchar(255);not null" json:"full_name"`	
	Username	string		`sql:"type:varchar(50);not null;unique" json:"username"`
	GroupID		int			`sql:"type:int(10)" json:"group_id"`//`sql:"type:int(10) REFERENCES r_groups(id)"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
	DeletedAt 	*time.Time 	`json:"deleted_at"`
}

type MemberModel struct{}

func (m *MemberModel) Find() (list []Member, err error) {
	db := DbConn()
	defer db.Close()
	if err := db.Find(&list).Error; err != nil {
		return list, err
	}
	return list, err
}

func (m *MemberModel) Get(id string) (member Member, err error) {
	db := DbConn()
	defer db.Close()
	if err := db.First(&member, id).Error; err != nil {
		return member, err
	}
	return member, err
}

func (m *MemberModel) Create(data Member) (newMember Member, err error) {
	db := DbConn()
	defer db.Close()
	if err := db.Create(&newMember).Error; err != nil {
		return newMember, err
	}
	return newMember, err
}

func (m *MemberModel) Update(id string, data Member) (member Member, err error) {
	db := DbConn()
	defer db.Close()
	if err := db.First(&member, id).Error; err != nil {
		return member, err
	}
	if err := db.Save(&member).Error; err != nil {
		return member, err
	}
	return member, err
}

func (m *MemberModel) Delete(id string, data Member) (member Member, err error) {
	db := DbConn()
	defer db.Close()
	if err := db.First(&member, id).Error; err != nil {
		return member, err
	}
	if err := db.Delete(&member).Error; err != nil {
		return member, err
	}
	return member, err
}


