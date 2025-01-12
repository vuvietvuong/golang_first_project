package controller

import (
	"fmt"
	"peanut/config"
	"peanut/domain"
	"peanut/pkg/response"
	"peanut/repository"
	"peanut/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Usecase usecase.UserUsecase
}

func NewUserController(db *gorm.DB) *UserController {
	fmt.Println(config.IsDevelopment())
	return &UserController{
		Usecase: usecase.NewUserUsecase( repository.NewUserRepo(db)),
	}
}

func (c *UserController) GetUsers(ctx *gin.Context) {

}

func (c *UserController) GetUser(ctx *gin.Context) {

}

func (c *UserController) CreateUser(ctx *gin.Context) {
	user := domain.User{}
	if !bindJSON(ctx, &user) {
		return
	}

	err := c.Usecase.CreateUser(ctx, user)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, nil)
}
