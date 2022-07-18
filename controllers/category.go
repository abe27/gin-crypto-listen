package controllers

import (
	"net/http"

	"github.com/abe27/crypto/tracker/api/models"
	"github.com/abe27/crypto/tracker/api/services"
	"github.com/gin-gonic/gin"
)

func ShowAllCategories(c *gin.Context) {
	var r models.Response
	r.ID = services.Gid()
	var obj []models.Category
	db := services.DB
	err := db.Find(&obj).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage("Category")
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}
	r.Success = true
	r.Message = services.ShowAllDataMessage("Category")
	r.Data = obj
	c.JSON(http.StatusOK, &r)
}

func CreateCategory(c *gin.Context) {
	var r models.Response
	r.ID = services.Gid()
	var obj models.Category
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
	r.Success = true
	r.Message = services.CreateDataSuccessMessage(obj.Name)
	r.Data = obj
	c.JSON(http.StatusCreated, &r)
}

func ShowCategoryByID(c *gin.Context) {
	var r models.Response
	r.ID = services.Gid()
	var obj models.Category
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

	r.Success = true
	r.Message = services.FoundDataMessage(obj.ID)
	r.Data = obj
	c.JSON(http.StatusOK, &r)
}

func UpdateCategory(c *gin.Context) {
	var r models.Response
	r.ID = services.Gid()
	var obj models.Category
	err := c.ShouldBind(&obj)
	if err != nil {
		r.Message = services.CheckInputRequiredMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, &r)
	}
	// ค้นหาข้อมูล
	obj.ID = c.Param("id")
	db := services.DB

	err = db.Where("id=?", obj.ID).Updates(&models.Category{
		Name:        obj.Name,
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

	r.Success = true
	r.Message = services.UpdateDataMessage(obj.ID)
	r.Data = obj
	c.JSON(http.StatusOK, &r)
}

func DeleteCategory(c *gin.Context) {
	var r models.Response
	r.ID = services.Gid()
	var obj models.Category
	// ค้นหาข้อมูล
	db := services.DB
	err := db.Where("id=?", c.Param("id")).First(&obj).Error
	if err != nil {
		r.Message = services.NotFoundDataMessage(c.Param("id"))
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	err = db.Where("id=?", c.Param("id")).Delete(&models.Category{}).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	r.Success = true
	r.Message = services.DeleteDataMessage(c.Param("id"))
	c.JSON(http.StatusOK, &r)
}
