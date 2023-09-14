package productcontroller

import (
	"encoding/json"
	"github.com/MDSPROID/Todo-App-Golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"gorm.io/gorm"
	// "fmt"
)

func Index(c *gin.Context){
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products":products})
}

func Show(c *gin.Context){
	var product []models.Product
	id := c.Param("id") // "id" berdasarkan nama param yang ada di route
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
				return
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data":product})
}

func Create(c *gin.Context){
	var product models.Product
	
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product":product})
}

func Update(c *gin.Context){
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat update produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbaharui"})
}

func Delete(c *gin.Context){
	var product models.Product

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})

}