package base

import (
	"gorm.io/gorm"
	"time"
)

type Audit struct {
	CreatedAt time.Time
	CreatedBy uint `json:"created_by"`
	UpdatedAt time.Time
	UpdatedBy uint `json:"updated_by"`
}

func (a *Audit) BeforeCreate(*gorm.DB) (err error) {
	currentTime := time.Now()
	a.CreatedAt = currentTime
	a.UpdatedAt = currentTime
	return nil
}

func (a *Audit) BeforeUpdate(*gorm.DB) (err error) {
	// Set UpdatedAt to the current time
	a.UpdatedAt = time.Now()
	return nil
}
