package controllers

import (
	"net/http"

	"github.com/abe27/crypto/tracker/api/models"
	"github.com/abe27/crypto/tracker/api/services"
	"github.com/gin-gonic/gin"
)

func ShowAllCurrency(c *gin.Context) {
	var r models.Response

	var obj []models.Currency
	// ค้นหาข้อมูล
	db := services.DB
	err := db.Find(&obj).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	r.Message = services.ShowAllDataMessage("Currency")
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func CreateCurrency(c *gin.Context) {
	var r models.Response

	var obj models.Currency
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Message = services.CheckInputRequiredMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	// บันทึกข้อมูล
	db := services.DB
	err = db.Create(&obj).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	r.Message = services.CreateDataSuccessMessage(obj.ID)
	r.Data = &obj
	c.JSON(http.StatusCreated, &r)
}

func ShowCurrencyByID(c *gin.Context) {
	var r models.Response

	var obj models.Currency
	obj.ID = c.Param("id")
	// ค้นหาข้อมูล
	db := services.DB
	err := db.Where("id=?", obj.ID).First(&obj).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(c.Param("id"))
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	r.Message = services.FoundDataMessage(obj.ID)
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func UpdateCurrency(c *gin.Context) {
	var r models.Response

	var obj models.Currency
	obj.ID = c.Param("id")
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Message = services.CheckInputRequiredMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, &r)
		return
	}

	// ค้นหาข้อมูล
	db := services.DB
	err = db.Where("id=?", obj.ID).Updates(&models.Currency{
		Symbol:      obj.Symbol,
		Flag:        obj.Flag,
		Description: obj.Description,
		IsActive:    obj.IsActive,
	}).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	err = db.Where("id=?", obj.ID).First(&obj).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(obj.ID)
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	r.Message = services.UpdateDataMessage(obj.ID)
	r.Data = &obj
	c.JSON(http.StatusOK, &r)
}

func DeleteCurrency(c *gin.Context) {
	var r models.Response

	var obj models.Currency
	obj.ID = c.Param("id")
	// ลบข้อมูล
	db := services.DB
	err := db.Where("id=?", obj.ID).First(&obj).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(obj.ID)
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	err = db.Where("id=?", obj.ID).Delete(&models.Currency{}).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	r.Message = services.DeleteDataMessage(c.Param("id"))
	c.JSON(http.StatusOK, &r)
}
