package handlers

import (
	"net/http"
	"strconv"

	"gmdb/services"
	"gmdb/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Global service instance (in production, inject via dependency injection)
var actorService *services.ActorService

// InitActorHandlers initializes the handlers with required dependencies
func InitActorHandlers(service *services.ActorService) {
	actorService = service
}

// HandleGetActors retrieves all actors with optional pagination
func HandleGetActors(c *gin.Context) {
	// Parse query parameters
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	// Call service
	actors, err := actorService.GetAllActors(limit, offset)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Actors retrieved successfully", actors)
}

// HandleGetActor retrieves a single actor by ID
func HandleGetActor(c *gin.Context) {
	// Parse UUID from path
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid actor ID")
		return
	}

	// Call service
	actor, err := actorService.GetActor(id)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "actor not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(c, status, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Actor retrieved successfully", actor)
}

// HandleCreateActor creates a new actor
func HandleCreateActor(c *gin.Context) {
	var req services.CreateActorRequest

	// Bind JSON input
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Call service (handles all business logic)
	actor, err := actorService.CreateActor(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Actor created successfully", actor)
}

// HandleUpdateActor updates an existing actor
func HandleUpdateActor(c *gin.Context) {
	// Parse UUID from path
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid actor ID")
		return
	}

	var req services.CreateActorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Call service
	actor, err := actorService.UpdateActor(id, req)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "actor not found" {
			status = http.StatusNotFound
		} else if err.Error() != "failed to update actor" {
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(c, status, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Actor updated successfully", actor)
}

// HandleDeleteActor soft deletes an actor
func HandleDeleteActor(c *gin.Context) {
	// Parse UUID from path
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid actor ID")
		return
	}

	// Call service
	if err := actorService.DeleteActor(id); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "actor not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(c, status, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Actor deleted successfully", nil)
}
