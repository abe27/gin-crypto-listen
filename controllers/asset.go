package controllers

import (
	"net/http"

	"github.com/abe27/crypto/tracker/api/models"
	"github.com/abe27/crypto/tracker/api/services"
	"github.com/gin-gonic/gin"
)

func ShowAllAsset(c *gin.Context) {
	var r models.Response
	var obj models.Asset
	// Fetch all assets
	db := services.DB
	err := db.Find(&obj).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = nil
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}
	r.Message = services.ShowAllDataMessage("Asset")
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func CreateAsset(c *gin.Context) {
	var r models.Response
	var obj models.Asset
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Message = services.CheckInputRequiredMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, &r)
		return
	}

	// Create a new asset
	db := services.DB
	err = db.Create(&obj).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}
	r.Message = services.CreateDataSuccessMessage("Asset")
	r.Data = &obj
	c.JSON(http.StatusCreated, &r)
}

func ShowAssetByID(c *gin.Context) {
	var r models.Response
	var obj models.Asset
	// Fetch By ID
	db := services.DB
	err := db.Where("id=?", c.Param("id")).First(&obj).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(c.Param("id"))
		r.Data = nil
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}
	r.Message = services.FoundDataMessage(c.Param("id"))
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func UpdateAsset(c *gin.Context) {
	var r models.Response
	var obj models.Asset
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Message = services.CheckInputRequiredMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, &r)
		return
	}

	// Initialize DB
	db := services.DB
	// Fetch asset by ID
	err = db.Where("id=?", c.Param("id")).First(&models.Asset{}).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(c.Param("id"))
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}
	// Update the asset
	err = db.Where("id=?", c.Param("id")).Updates(&models.Asset{
		CategoryID:  obj.CategoryID,
		CryptoID:    obj.CryptoID,
		Description: obj.Description,
		IsActive:    obj.IsActive,
	}).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	r.Message = services.UpdateDataMessage(c.Param("id"))
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func DeleteAsset(c *gin.Context) {
	var r models.Response
	// Delete the asset By ID
	db := services.DB
	err := db.Where("id=?", c.Param("id")).First(&models.Asset{}).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(c.Param("id"))
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	err = db.Where("id=?", c.Param("id")).Delete(&models.Asset{}).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	r.Message = services.DeleteDataMessage(c.Param("id"))
	r.Data = nil
	c.JSON(http.StatusOK, &r)
}
