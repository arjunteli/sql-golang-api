package main

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetPersonInfo(c *gin.Context) {

	personIDParam := c.Param("person_id")

	personID, err := strconv.Atoi(personIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return
	}

	personInfo, err := FetchPersonInfo(personID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		} else {

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	c.JSON(http.StatusOK, personInfo)
}

func CreatePerson(c *gin.Context) {
	var newPerson CreatePersonRequest

	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := InsertNewPerson(newPerson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create person"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person created successfully"})
}
