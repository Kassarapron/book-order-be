package controller

import (
        "backend/entity"
        "github.com/gin-gonic/gin"
        "net/http"
)

// GET /BookOrder
func ListBookOrder(c *gin.Context) {
	var BookOrder []entity.BookOrder
	if err := entity.DB().Raw("SELECT * FROM BookOrder").Scan(&BookOrder).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": BookOrder})
}


// GET /BookOrder /id
func ListBookOrder(c *gin.Context) {
	var BookOrder entity.BookOrder
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM BookOrder WHERE id = ?",id).Scan(&BookOrder).Error; err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": BookOrder})
}

// POST /BookOrder
func CreateBookOrder(c *gin.Context) {
	var BookOrder entity.BookOrder
	if err := c.ShouldBindJSON(&BookOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if err := entity.DB().Create(&BookOrder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": BookOrder})
}

// PATCH /BookOrder
func UpdateBookOrder(c *gin.Context) {
	var BookOrder entity.BookOrder
	if err := c.ShouldBindJSON(&BookOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if tx := entity.DB().Where("id = ?", BookOrder.ID).First(&BookOrder); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "BookOrder not found"})
			return
	}
	if err := entity.DB().Save(&BookOrder).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": BookOrder})
}

// DELETE /BookOrder/:id
func DeleteBookOrder(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM BookOrder WHERE id = ?", id); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "BookOrder not found"})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": BookOrder})
}