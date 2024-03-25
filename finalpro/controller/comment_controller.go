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

type CommentController interface {
	GetAllComments(ctx *gin.Context)
	GetCommentByID(ctx *gin.Context)
	CreateComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

func NewCommentController(srv service.CommentService) CommentController {
	return &commentController{srv: srv}
}

type commentController struct {
	srv service.CommentService
}

func (c *commentController) GetAllComments(ctx *gin.Context) {
	response, err := c.srv.GetAllComments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, "OK"))
}

func (c *commentController) GetCommentByID(ctx *gin.Context) {
	commentParamID := ctx.Param("commentId")
	commentID, err := strconv.Atoi(commentParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ID := uint(commentID)
	response, err := c.srv.GetCommentByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, "OK"))
}

func (c *commentController) CreateComment(ctx *gin.Context) {
	data := new(model.Comment)

	err := ctx.ShouldBindJSON(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	data.UserID = uint(userData["id"].(float64))

	_, err = c.srv.CreateComment(*data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, nil, "Created Comment Success"))
}

func (c *commentController) UpdateComment(ctx *gin.Context) {
	data := new(model.Comment)

	err := ctx.ShouldBindJSON(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	commentParamID := ctx.Param("commentId")
	commentID, err := strconv.Atoi(commentParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	data.UserID = uint(userData["id"].(float64))

	ID := uint(commentID)
	_, err = c.srv.UpdateComment(*data, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, nil, "Updated Comment Success"))
}

func (c *commentController) DeleteComment(ctx *gin.Context) {
	paramKeyID := ctx.Param("commentId")
	commentID, _ := strconv.Atoi(paramKeyID)
	ID := uint(commentID)
	err := c.srv.DeleteComment(ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, nil, "Deleted Comment Success"))
}
