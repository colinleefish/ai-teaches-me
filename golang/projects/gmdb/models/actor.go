package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Actor struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	BirthDate *time.Time     `json:"birth_date"`
	Biography string         `json:"biography" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships - no foreign key constraints
	Movies []Movie `json:"movies" gorm:"many2many:movie_actors;constraint:OnDelete:SET NULL"`
	Awards []Award `json:"awards" gorm:"foreignKey:ActorID;constraint:OnDelete:SET NULL"`
}

type Actor struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()`
}
