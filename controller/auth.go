package controller

import (
	"github.com/gin-gonic/gin"
	"golang-basic/domain"
	"golang-basic/pkg/response"
	"golang-basic/usecase"
	"gorm.io/gorm"
	"net/http"
)

type AuthController struct {
	Usecase usecase.AuthUsecase
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		Usecase: usecase.NewAuthUsecase(db),
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	user := domain.Auth{}

	if !bindJSON(ctx, &user) {
		return
	}

	jwt, _ := c.Usecase.Login(ctx, user)
	if jwt != nil {
		response.OK(ctx, jwt)
	} else {
		response.WithStatusCode(ctx, http.StatusUnauthorized, nil)
	}
}
