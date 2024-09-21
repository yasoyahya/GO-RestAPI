package categorycontroller

import (
	"GO-RESTAPI-GIN/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var category []models.Category

	models.DB.Find(&category)

	c.JSON(http.StatusOK, gin.H{"Category": category})
}
