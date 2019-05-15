package models

import (
	"time"
	"math/rand"
	"unsafe"

	"../db"
)

type RGroup struct {
	ID        uint       `json:"id"`
	Code      string     `sql:"type:varchar(50);not null;unique" json:"code"`
	Name      string     `sql:"type:varchar(50);not null" json:"name"`
	Size      int        `json:"size"`
	CreatedBy uint       `sql:"not null" json:"created_by"`
	Admin     uint       `sql:"not null" json:"admin"`
	Members	  []Member	 `gorm:"foreignkey:GroupID;association_foreignkey:ID" json:"members"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type GroupModel struct{}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func GenerateGroupCode(n int) string {
	b := make([]byte, n)
	var src = rand.NewSource(time.Now().UnixNano())
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return *(*string)(unsafe.Pointer(&b))
}

func (m *GroupModel) Find(list *[]RGroup) (err error) {
	if err := db.DB.Find(list).Error; err != nil {
		return err
	}
	return nil
}

func (m *GroupModel) Get(id string, rGroup *RGroup) (err error) {
	if err := db.DB.First(rGroup, id).Error; err != nil {
		return err
	}

	var members []Member
	
	if err := db.DB.Where("group_id = ?", id).Find(&members).Error; err != nil {
		return err
	}

	rGroup.Members = members

	return nil
}

func (m *GroupModel) Create(rGroup *RGroup) (err error) {
	rGroup.Code = GenerateGroupCode(14)

	if err := db.DB.Create(rGroup).Error; err != nil {
		return err
	}
	return nil
}