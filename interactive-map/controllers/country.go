package controllers

import (
	"fmt"
	"interactive-map/database"
	"interactive-map/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetCountries(c *gin.Context) {
	var user models.User
	claims := c.MustGet("user").(*jwt.RegisteredClaims)
	if err := database.DB.First(&user, claims.Subject).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var countries []models.CountryStatus
	if err := database.DB.Where("user_id = ?", user.ID).Find(&countries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch countries"})
		return
	}

	c.JSON(http.StatusOK, countries)
}

func UpdateCountryStatus(c *gin.Context) {
	var input models.CountryStatus
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	var user models.User
	claims := c.MustGet("user").(*jwt.RegisteredClaims)
	if err := database.DB.First(&user, claims.Subject).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	input.UserID = user.ID
	if err := database.DB.Save(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save country status"})
		return
	}

	c.JSON(http.StatusOK, input)
}

// DELETE /countries/:countryISO
func DeleteCountry(c *gin.Context) {
	var user models.User
	claims := c.MustGet("user").(*jwt.RegisteredClaims)
	if err := database.DB.First(&user, claims.Subject).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	countryISO := c.Param("countryISO")
	if err := database.DB.Where("user_id = ? AND country_iso = ?", user.ID, countryISO).Delete(&models.CountryStatus{}).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete country"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Country deleted"})
}
