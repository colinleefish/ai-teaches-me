package models

import "github.com/google/uuid"

// MovieActor represents the join table for many-to-many relationship
// GORM will create this table automatically with the many2many tag
type MovieActor struct {
	MovieID uuid.UUID `gorm:"primaryKey;type:uuid"`
	ActorID uuid.UUID `gorm:"primaryKey;type:uuid"`
}
