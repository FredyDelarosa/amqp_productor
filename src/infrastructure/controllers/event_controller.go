package controllers

import (
	"net/http"
	"sensor/src/application"
	"sensor/src/domain/entities"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	useCase *application.CreateEventUseCase
}

func NewEventController(useCase *application.CreateEventUseCase) *EventController {
	return &EventController{useCase: useCase}
}

func (ctrl *EventController) CreateEvent(c *gin.Context) {
	var event entities.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.useCase.Execute(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error procesando el evento"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Evento registrado exitosamente"})
}
