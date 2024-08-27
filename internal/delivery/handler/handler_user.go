package handler

import (
	"github.com/gin-gonic/gin"
	"golang-gorm/internal/domain/dto"
	"golang-gorm/internal/usecase"
	"net/http"
)

type HandlerUser interface {
	CreateUserHandler(ctx *gin.Context)
	GetUserByIdHandler(ctx *gin.Context)
	Route()
}

type handlerUser struct {
	uc usecase.UsecaseUser
	rg *gin.RouterGroup
}

func (h handlerUser) CreateUserHandler(ctx *gin.Context) {
	var request dto.UserRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.BaseResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: err.Error(),
			Data:            nil,
		})
	}

	result, err := h.uc.CreateUserUsecase(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.BaseResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: err.Error(),
			Data:            nil,
		})
	}
	ctx.JSON(http.StatusCreated, dto.BaseResponse{
		ResponseCode:    http.StatusCreated,
		ResponseMessage: "CREATED",
		Data:            result,
	})
}

func (h handlerUser) GetUserByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := h.uc.FindUserByIdUsecase(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.BaseResponse{
			ResponseCode:    http.StatusBadRequest,
			ResponseMessage: err.Error(),
			Data:            nil,
		})
	}

	ctx.JSON(http.StatusOK, dto.BaseResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "OK",
		Data:            result,
	})
}

func (h handlerUser) Route() {
	ug := h.rg.Group("/user")
	ug.POST("/", h.CreateUserHandler)
	ug.GET("/userId", h.GetUserByIdHandler)
}

func NewHandlerUser(uc usecase.UsecaseUser, rg *gin.RouterGroup) HandlerUser {
	return &handlerUser{
		uc: uc,
		rg: rg,
	}
}
