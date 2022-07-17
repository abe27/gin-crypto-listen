package services

import (
	"net/http"
	"strings"

	"github.com/abe27/crypto/tracker/api/models"
	"github.com/gin-gonic/gin"
)

func AuthorizationRequired(c *gin.Context) {
	// secret_key := os.Getenv("SECRET_KEY")
	var r models.Response
	r.ID = Gid()
	s := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")

	if token == "" {
		r.Message = RequiredAuthenticationMessage
		c.JSON(http.StatusUnauthorized, &r)
		c.Abort()
		return
	}

	// Check Token On DB
	db := DB
	var jwtToken models.JwtToken
	err := db.Where("key=?", token).Find(&jwtToken).Error
	if err != nil {
		r.Message = SystemErrorMessage
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}

	if jwtToken.ID == "" {
		r.Message = NotFoundTokenMessage
		c.AbortWithStatusJSON(http.StatusUnauthorized, &r)
		return
	}
	_, er := ValidateToken(jwtToken.Token)
	if er != nil {
		r.Message = TokenExpiredMessage
		db.Delete(&jwtToken)
		c.AbortWithStatusJSON(http.StatusInternalServerError, &r)
		return
	}
	c.Next()
}
