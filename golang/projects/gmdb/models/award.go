package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Award struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Category    string         `json:"category"`
	Year        int            `json:"year"`
	MovieID     *uuid.UUID     `json:"movie_id"`
	ActorID     *uuid.UUID     `json:"actor_id"`
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships - award can belong to either movie or actor (no FK constraints)
	Movie *Movie `json:"movie,omitempty" gorm:"foreignKey:MovieID;constraint:OnDelete:SET NULL"`
	Actor *Actor `json:"actor,omitempty" gorm:"foreignKey:ActorID;constraint:OnDelete:SET NULL"`
}

// BeforeCreate hook to generate UUIDv7
func (a *Award) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == (uuid.UUID{}) {
		a.ID = uuid.Must(uuid.NewV7())
	}
	return
}
