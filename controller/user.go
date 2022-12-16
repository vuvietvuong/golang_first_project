package controller

import (
	"fmt"
	"golang-basic/config"
	user2 "golang-basic/domain"
	"golang-basic/pkg/response"
	"golang-basic/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Usecase usecase.UserUsecase
}

func NewUserController(db *gorm.DB) *UserController {
	fmt.Println(config.IsDevelopment())
	return &UserController{
		Usecase: usecase.NewUserUsecase(db),
	}
}

func (c *UserController) GetUsers(ctx *gin.Context) {

}

func (c *UserController) GetUser(ctx *gin.Context) {

}

func (c *UserController) CreateUser(ctx *gin.Context) {
	user := user2.CreateUserReq{}
	if !bindJSON(ctx, &user) {
		return
	}

	err := c.Usecase.CreateUser(ctx, user)
	if checkError(ctx, err) {
		return
	}

	response.WithStatusCode(ctx, http.StatusCreated, nil)
}
