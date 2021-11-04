package controller

import (
        "github.com/Kassarapron/newproject/entity"
        "github.com/gin-gonic/gin"
        "net/http"
)

// GET /Company
func ListCompany(c *gin.Context) {
	var Company []entity.Company
	if err := entity.DB().Raw("SELECT * FROM Company").Scan(&Company).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": Company})
}


// GET /Company /id
func ListCompany(c *gin.Context) {
	var Company entity.Company
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Company WHERE id = ?",id).Scan(&Company).Error; err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": Company})
}

// POST /Company
func CreateCompany(c *gin.Context) {
	var Company entity.Company
	if err := c.ShouldBindJSON(&Company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if err := entity.DB().Create(&Company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Company})
}

// PATCH /Company
func UpdateCompany(c *gin.Context) {
	var Company entity.Company
	if err := c.ShouldBindJSON(&Company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if tx := entity.DB().Where("id = ?", Company.ID).First(&Company); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Company not found"})
			return
	if err := entity.DB().Save(&Company).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": v})
}

// DELETE /Company/:id
func DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Company WHERE id = ?", id); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Company not found"})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": Company})
}