package controllers

import (
	"net/http"

	"github.com/abe27/crypto/tracker/api/models"
	"github.com/abe27/crypto/tracker/api/services"
	"github.com/gin-gonic/gin"
)

func ShowAllCryptoCurrency(c *gin.Context) {
	var r models.Response
	var obj []models.Cryptocurrency
	// Fetch All Data
	db := services.DB
	err := db.Find(&obj).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}
	r.Message = services.ShowAllDataMessage("Crypto Currency")
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func CreateCryptoCurrency(c *gin.Context) {
	var r models.Response
	var obj models.Cryptocurrency
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Message = services.CheckInputRequiredMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	// Create a new currency
	db := services.DB
	err = db.Create(&obj).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}
	r.Message = services.CreateDataSuccessMessage("Crypto Currency")
	r.Data = &obj
	c.JSON(http.StatusCreated, &r)
}

func ShowCryptoCurrencyByID(c *gin.Context) {
	var r models.Response
	var obj models.Cryptocurrency
	// Fetch data by ID
	db := services.DB
	err := db.Where("id", c.Param("id")).First(&obj).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(c.Param("id"))
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	r.Message = services.FoundDataMessage(c.Param("id"))
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func UpdateCryptoCurrency(c *gin.Context) {
	var r models.Response
	var obj models.Cryptocurrency
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Message = services.CheckInputRequiredMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, &r)
		return
	}

	// Update Data By ID
	db := services.DB
	err = db.Where("id", c.Param("id")).First(&models.Cryptocurrency{}).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(c.Param("id"))
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	err = db.Where("id", c.Param("id")).Updates(&models.Cryptocurrency{
		Name:        obj.Name,
		Symbol:      obj.Symbol,
		Address:     obj.Address,
		Flag:        obj.Flag,
		Description: obj.Description,
		IsActive:    obj.IsActive,
	}).Error

	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
	}
	r.Message = services.UpdateDataMessage(c.Param("id"))
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func DeleteCryptoCurrency(c *gin.Context) {
	var r models.Response
	var obj models.Cryptocurrency
	// Update Data By ID
	db := services.DB
	err := db.Where("id", c.Param("id")).First(&obj).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(c.Param("id"))
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	err = db.Delete(&obj).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
	}
	r.Message = services.DeleteDataMessage(c.Param("id"))
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}
