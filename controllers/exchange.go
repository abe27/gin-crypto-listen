package controllers

import (
	"net/http"

	"github.com/abe27/crypto/tracker/api/models"
	"github.com/abe27/crypto/tracker/api/services"
	"github.com/gin-gonic/gin"
)

func ShowAllExchange(c *gin.Context) {
	db := services.DB
	var r models.Response
	var exchange []models.Exchange
	r.ID = services.Gid()
	err := db.Find(&exchange).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage("Exchange")
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	r.Message = services.ShowAllDataMessage("Exchange")
	r.Data = &exchange
	c.JSON(http.StatusOK, &r)
}

func CreateExchange(c *gin.Context) {
	var r models.Response
	r.ID = services.Gid()
	var obj models.Exchange
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Message = services.CheckInputRequiredMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, &r)
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

	r.Message = services.CreateDataSuccessMessage(obj.Name)
	r.Data = &obj
	c.JSON(http.StatusCreated, &r)
}

func ShowExchangeByID(c *gin.Context) {
	var r models.Response
	r.ID = services.Gid()
	var obj models.Exchange
	obj.ID = c.Param("id")

	// ค้นหาข้อมูล
	db := services.DB
	err := db.Where("id=?", obj.ID).First(&obj).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(obj.ID)
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}
	r.Message = services.FoundDataMessage(obj.Name)
	r.Data = &obj
	c.JSON(http.StatusFound, &r)
}

func UpdateExchange(c *gin.Context) {
	var r models.Response
	var obj models.Exchange
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Message = services.CheckInputRequiredMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, &r)
		return
	}
	// ค้นหาข้อมูล
	obj.ID = c.Param("id")
	db := services.DB
	err = db.Where("id=?", obj.ID).Updates(&models.Exchange{
		Name:        obj.Name,
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

	err = db.First(&obj).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(obj.ID)
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	r.Message = services.UpdateDataMessage(obj.Name)
	r.Data = &obj
	c.JSON(http.StatusAccepted, &r)
}

func DeleteExchange(c *gin.Context) {
	var r models.Response
	r.ID = services.Gid()
	var obj models.Exchange
	obj.ID = c.Param("id")
	// ลบข้อมูล
	db := services.DB

	err := db.First(&obj).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(obj.ID)
		r.Data = nil
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	err = db.Where("id=?", obj.ID).Delete(&models.Exchange{}).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
	}

	r.Message = services.DeleteDataMessage(obj.Name)
	r.Data = nil
	c.JSON(http.StatusOK, &r)
}
