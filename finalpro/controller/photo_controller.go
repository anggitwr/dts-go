package controller

import (
	"finalpro/helper"
	"finalpro/model"
	"finalpro/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type PhotoController interface {
	GetAllPhotos(ctx *gin.Context)
	GetPhotoByID(ctx *gin.Context)
	CreatePhoto(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

func NewPhotoController(srv service.PhotoService) PhotoController {
	return &photoController{srv: srv}
}

type photoController struct {
	srv service.PhotoService
}

func (c *photoController) GetAllPhotos(ctx *gin.Context) {
	response, err := c.srv.GetAllPhotos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, "OK"))
}

func (c *photoController) GetPhotoByID(ctx *gin.Context) {
	photoParamID := ctx.Param("photoId")
	photoID, err := strconv.Atoi(photoParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ID := uint(photoID)
	response, err := c.srv.GetPhotoByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, "OK"))
}

func (c *photoController) CreatePhoto(ctx *gin.Context) {
	dataReq := new(model.RequestPhoto)
	data := new(model.Photo)
	err := ctx.ShouldBindJSON(dataReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	data.UserID = uint(userData["id"].(float64))
	data.Title = dataReq.Title
	data.Caption = dataReq.Caption
	data.PhotoURL = dataReq.PhotoURL
	_, err = c.srv.CreatePhoto(*data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, nil, "Created Photo Success"))
}

func (c *photoController) UpdatePhoto(ctx *gin.Context) {
	dataReq := new(model.RequestPhoto)
	data := new(model.Photo)

	err := ctx.ShouldBindJSON(dataReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	photoParamID := ctx.Param("photoId")
	photoID, err := strconv.Atoi(photoParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	data.UserID = uint(userData["id"].(float64))
	data.Title = dataReq.Title
	data.Caption = dataReq.Caption
	data.PhotoURL = dataReq.PhotoURL
	ID := uint(photoID)
	_, err = c.srv.UpdatePhoto(*data, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, nil, "Updated Photo Success"))
}

func (c *photoController) DeletePhoto(ctx *gin.Context) {
	paramKeyID := ctx.Param("photoId")
	pID, _ := strconv.Atoi(paramKeyID)
	ID := uint(pID)
	err := c.srv.DeletePhoto(ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, nil, "Deleted Photo Success"))
}
