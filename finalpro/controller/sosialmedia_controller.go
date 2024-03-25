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

type SocialMediaController interface {
	GetAllSosmed(ctx *gin.Context)
	GetSosmedByID(ctx *gin.Context)
	CreateSosmed(ctx *gin.Context)
	UpdateSosmed(ctx *gin.Context)
	DeleteSosmed(ctx *gin.Context)
}

func NewSosmedController(srv service.SocialMediaService) SocialMediaController {
	return &socialMediaController{srv: srv}
}

type socialMediaController struct {
	srv service.SocialMediaService
}

func (c *socialMediaController) GetAllSosmed(ctx *gin.Context) {
	response, err := c.srv.GetAllSosmeds()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, "OK"))
}

func (c *socialMediaController) GetSosmedByID(ctx *gin.Context) {
	sosmedParamID := ctx.Param("socialMediaId")
	sosmedID, err := strconv.Atoi(sosmedParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ID := uint(sosmedID)
	response, err := c.srv.GetSosmedByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, "OK"))
}

func (c *socialMediaController) CreateSosmed(ctx *gin.Context) {
	data := new(model.SocialMedia)

	err := ctx.ShouldBindJSON(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	data.UserID = uint(userData["id"].(float64))

	_, err = c.srv.CreateSosmed(*data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, nil, "Created Social Media Success"))
}

func (c *socialMediaController) UpdateSosmed(ctx *gin.Context) {
	data := new(model.SocialMedia)

	err := ctx.ShouldBindJSON(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	sosmedParamID := ctx.Param("socialMediaId")
	sosmedID, err := strconv.Atoi(sosmedParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	data.UserID = uint(userData["id"].(float64))

	ID := uint(sosmedID)
	_, err = c.srv.UpdateSosmed(*data, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, nil, "Updated Social Media Success"))
}
func (c *socialMediaController) DeleteSosmed(ctx *gin.Context) {
	paramKeyID := ctx.Param("socialMediaId")
	sosmedID, _ := strconv.Atoi(paramKeyID)
	ID := uint(sosmedID)
	err := c.srv.DeleteSosmed(ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, nil, "Deleted Social Media Success"))
}
