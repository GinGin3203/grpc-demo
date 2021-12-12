package model

import "time"

type UserRecord struct {
	ID            int32
	Role          string
	Name          string
	LastUpdatedAt time.Time
}
