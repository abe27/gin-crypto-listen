package controllers

import (
	"net/http"
	"strings"

	"github.com/abe27/crypto/tracker/api/models"
	"github.com/abe27/crypto/tracker/api/services"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var r models.Response
	var u models.User

	// Map ข้อมูล
	err := c.ShouldBind(&u)
	if err != nil {
		r.Message = services.CheckInputRequiredMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, &r)
		return
	}

	passwd, _ := services.HashPassword(u.Password)
	u.Password = passwd
	db := services.DB
	// ตรวจสอบข้อผิดพลาดและบันทึกข้อมูล
	err = db.Create(&u).Error
	if err != nil {
		r.Message = services.DataIsDuplicateMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusBadRequest, &r)
		return
	}

	r.Message = services.RegisterSuccessMessage
	r.Data = &u
	c.JSON(http.StatusCreated, &r)
}

func SignIn(c *gin.Context) {
	var r models.Response
	var u models.User

	username := c.PostForm("email")

	// ตรวจสอบข้อมูลผู้ใช้งาน
	db := services.DB
	err := db.Where("email=?", username).First(&u).Error
	if err != nil {
		r.Message = services.NotFoundUserMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusNotFound, &r)
		return
	}

	// Compare HashPassword
	IsSuccess := services.CheckPasswordHash(c.PostForm("password"), u.Password)
	if !IsSuccess {
		r.Message = services.PasswordIsNotMatchMessage
		r.Data = nil
		c.AbortWithStatusJSON(http.StatusUnauthorized, &r)
		return
	}

	// ทำการ Generate Token
	header, tokenType, token, er := services.CreateToken(u.ID)
	if er != nil {
		r.Message = services.SystemErrorMessage
		r.Data = er
		c.AbortWithStatusJSON(http.StatusBadRequest, &r)
		return
	}

	var auth models.AuthResponse
	auth.ID = services.Gid()
	auth.Header = header
	auth.Type = tokenType
	auth.Token = token
	r.Message = services.SigInSuccessMessage
	r.Data = &auth
	c.JSON(http.StatusOK, &r)
}

func SignOut(c *gin.Context) {
	var r models.Response

	s := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	if token == "" {
		r.Message = services.AuthenticateRequiredTokenMessage
		c.AbortWithStatusJSON(http.StatusUnauthorized, &r)
		return
	}
	// Delete Token On DB
	db := services.DB
	err := db.Where("key=?", token).Delete(&models.JwtToken{}).Error
	if err != nil {
		r.Message = services.SystemErrorMessage
		r.Data = err
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	r.Message = services.UserLeaveMessage
	r.Data = nil
	c.JSON(http.StatusOK, &r)
}
