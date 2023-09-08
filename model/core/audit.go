package core

import (
	"gorm.io/gorm"
	"time"
)

type Audit struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uint      `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uint      `json:"updated_by"`
}

func (a *Audit) BeforeCreate(*gorm.DB) (err error) {
	currentTime := time.Now()
	a.CreatedAt = currentTime
	a.UpdatedAt = currentTime
	return nil
}

func (a *Audit) BeforeUpdate(*gorm.DB) (err error) {
	a.UpdatedAt = time.Now()
	return nil
}
