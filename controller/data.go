package controller

import (
	"gin/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func GetDate(c *gin.Context) {

	var data entity.Data

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM datas WHERE id = ? ", id).Scan(&data).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": data})

}

func ListUser(c *gin.Context) {

	var dates []entity.Data

	if err := entity.DB().Raw("SELECT * FROM datas").Scan(&dates).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": dates})

}
