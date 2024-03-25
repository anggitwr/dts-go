package controller

import (
	"finalpro/helper"
	"finalpro/model"
	"finalpro/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	srv service.UserService
}

func NewUserController(srv service.UserService) UserController {
	return &userController{srv}
}

func (c *userController) Register(ctx *gin.Context) {
	data := new(model.User)
	dataReq := new(model.RequestRegister)
	if err := ctx.ShouldBindJSON(dataReq); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	data.Email = dataReq.Email
	data.Password = dataReq.Password
	data.Username = dataReq.Username
	data.Age = dataReq.Age
	err := c.srv.Create(data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, nil, "Register Successfully"))
}
func (c *userController) Login(ctx *gin.Context) {
	data := new(model.RequestLogin)

	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}

	response, err := c.srv.Login(data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, "Login Successfully"))
}

func (c *userController) GetUserByID(ctx *gin.Context) {
	userParamID := ctx.Param("userId")
	userID, err := strconv.Atoi(userParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ID := uint(userID)
	response, err := c.srv.GetUserByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, "OK"))
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	dataReq := new(model.RequestRegister)
	data := new(model.User)

	err := ctx.ShouldBindJSON(dataReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	userParamID := ctx.Param("userId")
	userID, err := strconv.Atoi(userParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	data.Email = dataReq.Email
	data.Password = dataReq.Password
	data.Username = dataReq.Username
	data.Age = dataReq.Age
	ID := uint(userID)
	_, err = c.srv.UpdateUser(*data, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, nil, "Updated user Success"))
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	paramKeyID := ctx.Param("userId")
	userID, _ := strconv.Atoi(paramKeyID)
	ID := uint(userID)
	err := c.srv.DeleteUser(ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, nil, "Deleted User Success"))
}
