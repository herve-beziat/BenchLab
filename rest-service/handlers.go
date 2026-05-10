package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herve-beziat/BenchLab/internal/store"
)

// createSensor gère le POST /sensors
func createSensor(s *store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sensor store.Sensor
		if err := c.ShouldBindJSON(&sensor); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		created, err := s.CreateSensor(sensor)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, created)
	}
}
