package models

import "time"

type AuditLog struct {
	ID			int64		`gorm:"primaryKey"`
	UserID		int64
	Action		string
	Resource	string
	ResourceID	int64
	OldValue	string
	NewValue	string
	CreatedAt	time.Time
}