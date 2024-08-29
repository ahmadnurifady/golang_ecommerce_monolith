package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"golang-gorm/internal/domain/dto"
	"log"
	"net/http"
)

type HandlerAuth interface {
	GoogleAuth(ctx *gin.Context)
	GoogleAuthCallback(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Route()
}

type handlerAuth struct {
	rg *gin.RouterGroup
}

func (h handlerAuth) Route() {
	ag := h.rg.Group("/auth")
	ag.GET("/:provider", h.GoogleAuth)
	ag.GET("/google/callback", h.GoogleAuthCallback)
	ag.GET("/logout", h.Logout)
}

func (h handlerAuth) GoogleAuth(ctx *gin.Context) {
	provider := ctx.Param("provider")

	q := ctx.Request.URL.Query()
	q.Add("provider", provider)

	ctx.Request.URL.RawQuery = q.Encode()

	req := ctx.Request
	res := ctx.Writer
	gothic.BeginAuthHandler(res, req)
}

func (h handlerAuth) GoogleAuthCallback(ctx *gin.Context) {

	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		log.Println(err.Error())
	}

	ctx.JSON(http.StatusOK, dto.BaseResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "SUCCESS LOGIN GOOGLE",
		Data:            user,
	})
}

func (h handlerAuth) Logout(ctx *gin.Context) {
	gothic.Logout(ctx.Writer, ctx.Request)
	ctx.JSON(http.StatusOK, dto.BaseResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "SUCCESS LOGOT GOOGLE",
	})
}

func NewHandlerAuth(rg *gin.RouterGroup) HandlerAuth {
	return &handlerAuth{rg: rg}
}
