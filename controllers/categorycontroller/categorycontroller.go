package categorycontroller

import (
	"GO-RESTAPI-GIN/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var category []models.Category

	models.DB.Find(&category)

	// c.JSON(http.StatusOK, gin.H{"Category": category})

	type Response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"Data"`
	}

	response := Response{
		Status:  "success",
		Message: "data retrieved successfully",
		Data:    category,
	}

	c.JSON(http.StatusOK, response)

}

func Show(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := models.DB.First(&category, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	// c.JSON(http.StatusOK, gin.H{"Category": category})

	type Response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"Data"`
	}

	response := Response{
		Status:  "success",
		Message: "data retrieved successfully",
		Data:    category,
	}

	c.JSON(http.StatusOK, response)
}

func Create(c *gin.Context) {

	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	models.DB.Create(&category)
	c.JSON(http.StatusOK, gin.H{"product": category})
}

func Update(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&category).Where("id = ? ", id).Updates(&category).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {
	var product models.Category

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})

}
