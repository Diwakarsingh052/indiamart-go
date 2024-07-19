package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"small-app/models"
)

func (h *handler) Signup(c *gin.Context) {
	// Declare a variable to hold decoded data from request body
	var newUser models.NewUser
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		slog.Error("json validation error", slog.String("TRACE ID", "fake-id"), slog.String("Error", err.Error()))
		// Respond with a 400 Bad Request status code and error message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})

		// Stop further execution
		return
	}

	// Create a new validator instance
	validate := validator.New()
	// Validate the newUser struct using the validate instance
	err = validate.Struct(newUser)

	// Check if validation encountered errors
	if err != nil {
		slog.Error("validation failed", slog.String("TRACE ID", "fake-id"),
			slog.String("ERROR", err.Error()))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return

	}

	uData, err := h.CreateUser(newUser)
	// If user fetch operation fails, respond with an error
	if err != nil {

		slog.Error("error in creating the user", slog.String("Trace ID", "fake-id"),
			slog.String("Error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Problem in creating user"})
		return
	}

	c.JSON(http.StatusOK, uData)

}
