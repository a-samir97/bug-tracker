package models

import "time"

type Base struct {
	Id        int
	CreatedAt string
	UpdateAt  string
}

func GenerateISOString() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
}

// before creating the object, generate id, created_at and update_at fields
func (base *Base) BeforeCreate() error {
	timeNow := GenerateISOString()
	base.CreatedAt, base.UpdateAt = timeNow, timeNow

	return nil
}

// After creating the object, update updated_at field
func (base *Base) AfterUpdate() error {
	base.UpdateAt = GenerateISOString()
	return nil
}
