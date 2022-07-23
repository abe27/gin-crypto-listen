package controllers

import (
	"net/http"

	"github.com/abe27/crypto/tracker/api/models"
	"github.com/abe27/crypto/tracker/api/services"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	var r models.Response
	r.ID = services.Gid()

	r.Message = services.HelloApi
	c.JSON(http.StatusOK, &r)
}
