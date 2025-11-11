package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Movie struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Year        int            `json:"year"`
	Director    string         `json:"director"`
	Genre       string         `json:"genre"`
	Description string         `json:"description" gorm:"type:text"`
	Rating      float64        `json:"rating" gorm:"type:decimal(3,1)"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships - no foreign key constraints
	Actors []Actor `json:"actors" gorm:"many2many:movie_actors;constraint:OnDelete:SET NULL"`
	Awards []Award `json:"awards" gorm:"foreignKey:MovieID;constraint:OnDelete:SET NULL"`
}

// BeforeCreate hook to generate UUIDv7
func (m *Movie) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == (uuid.UUID{}) {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return
}
