package controller

import (
	"assignmentdua/model"
	"assignmentdua/repository"
	"assignmentdua/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderRepository repository.IOrderRepository
}

func NewOrderController(orderRepository repository.IOrderRepository) *orderController {

	return &orderController{
		orderRepository: orderRepository,
	}
}

func (oc *orderController) Create(ctx *gin.Context) {
	var newOrder model.Order
	err := ctx.ShouldBindJSON(&newOrder)

	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	createdOrder, err := oc.orderRepository.Create(newOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, createdOrder, ""))

}

func (oc *orderController) GetAll(ctx *gin.Context) {

	orders, err := oc.orderRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.CreateResponse(true, orders, ""))
}

func (oc *orderController) Delete(ctx *gin.Context) {

	idString := ctx.Param("id")
	// fmt.Println(idString)

	idInt, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = oc.orderRepository.Delete(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, nil, ""))
}

func (oc *orderController) UpdateOrder(ctx *gin.Context) {
	orderIDStr := ctx.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var updatedOrder model.Order
	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := oc.orderRepository.Update(orderID, updatedOrder); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}
