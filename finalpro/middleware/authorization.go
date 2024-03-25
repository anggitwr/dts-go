package middleware

import (
	"finalpro/helper"
	"finalpro/lib"
	"finalpro/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := lib.GetDB()
		userId, err := strconv.Atoi(c.Param("userId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, "Invalid Path Variable"))
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := model.User{}

		err = db.Select("id").First(&User, uint(userId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, helper.NewResponse(http.StatusNotFound, nil, "Data doesn't exist"))
			return
		}

		if User.ID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.NewResponse(http.StatusUnauthorized, nil, "You are not allowed to access this data"))
			return
		}
		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := lib.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, "Invalid Path Variable"))
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := model.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, helper.NewResponse(http.StatusNotFound, nil, "Data doesn't exist"))
			return
		}

		if Photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.NewResponse(http.StatusUnauthorized, nil, "You are not allowed to access this data"))
			return
		}
		c.Next()
	}
}
func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := lib.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, "Invalid Path Variable"))
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := model.Comment{}

		err = db.Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, helper.NewResponse(http.StatusNotFound, nil, "Data doesn't exist"))
			return
		}

		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.NewResponse(http.StatusUnauthorized, nil, "You are not allowed to access this data"))
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := lib.GetDB()
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := model.SocialMedia{}

		err = db.Select("user_id").First(&SocialMedia, uint(socialMediaId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Data Not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		if SocialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
