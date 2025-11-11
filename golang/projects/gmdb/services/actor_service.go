package services

import (
	"errors"
	"time"

	"gmdb/models"
	"gmdb/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ActorService struct {
	db *gorm.DB
}

// NewActorService creates a new actor service instance
func NewActorService(db *gorm.DB) *ActorService {
	return &ActorService{db: db}
}

// CreateActorRequest represents the input for creating an actor
type CreateActorRequest struct {
	Name      string     `json:"name" binding:"required"`
	BirthDate *time.Time `json:"birth_date"`
	Biography string     `json:"biography"`
}

// ActorResponse represents the output format for an actor
type ActorResponse struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	BirthDate *time.Time `json:"birth_date"`
	Biography string     `json:"biography"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// CreateActor handles the business logic for creating a new actor
func (s *ActorService) CreateActor(req CreateActorRequest) (*ActorResponse, error) {
	// Business validation
	if err := s.validateCreateActor(req); err != nil {
		return nil, err
	}

	// Create actor model
	actor := models.Actor{
		ID:        utils.NewUUIDv7(),
		Name:      req.Name,
		BirthDate: req.BirthDate,
		Biography: req.Biography,
	}

	// Save to database
	if err := s.db.Create(&actor).Error; err != nil {
		return nil, errors.New("failed to create actor")
	}

	// Transform to response
	return s.toResponse(actor), nil
}

// GetActor retrieves an actor by ID
func (s *ActorService) GetActor(id uuid.UUID) (*ActorResponse, error) {
	var actor models.Actor

	if err := s.db.First(&actor, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("actor not found")
		}
		return nil, errors.New("failed to retrieve actor")
	}

	return s.toResponse(actor), nil
}

// GetAllActors retrieves all actors with optional filters
func (s *ActorService) GetAllActors(limit, offset int) ([]*ActorResponse, error) {
	var actors []models.Actor

	query := s.db.Model(&models.Actor{})
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&actors).Error; err != nil {
		return nil, errors.New("failed to retrieve actors")
	}

	// Transform to responses
	responses := make([]*ActorResponse, len(actors))
	for i, actor := range actors {
		responses[i] = s.toResponse(actor)
	}

	return responses, nil
}

// UpdateActor updates an existing actor
func (s *ActorService) UpdateActor(id uuid.UUID, req CreateActorRequest) (*ActorResponse, error) {
	var actor models.Actor

	// Check if actor exists
	if err := s.db.First(&actor, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("actor not found")
		}
		return nil, errors.New("failed to retrieve actor")
	}

	// Business validation
	if err := s.validateCreateActor(req); err != nil {
		return nil, err
	}

	// Update fields
	actor.Name = req.Name
	actor.BirthDate = req.BirthDate
	actor.Biography = req.Biography

	if err := s.db.Save(&actor).Error; err != nil {
		return nil, errors.New("failed to update actor")
	}

	return s.toResponse(actor), nil
}

// DeleteActor soft deletes an actor
func (s *ActorService) DeleteActor(id uuid.UUID) error {
	result := s.db.Delete(&models.Actor{}, "id = ?", id)
	if result.Error != nil {
		return errors.New("failed to delete actor")
	}
	if result.RowsAffected == 0 {
		return errors.New("actor not found")
	}
	return nil
}

// Business logic validation
func (s *ActorService) validateCreateActor(req CreateActorRequest) error {
	if req.Name == "" {
		return errors.New("actor name is required")
	}
	if len(req.Name) < 2 {
		return errors.New("actor name must be at least 2 characters")
	}
	if req.BirthDate != nil && req.BirthDate.After(time.Now()) {
		return errors.New("birth date cannot be in the future")
	}
	return nil
}

// Transform model to response DTO
func (s *ActorService) toResponse(actor models.Actor) *ActorResponse {
	return &ActorResponse{
		ID:        actor.ID,
		Name:      actor.Name,
		BirthDate: actor.BirthDate,
		Biography: actor.Biography,
		CreatedAt: actor.CreatedAt,
		UpdatedAt: actor.UpdatedAt,
	}
}
