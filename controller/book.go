package controller

import (
	"github.com/gin-gonic/gin"
	jwt2 "github.com/golang-jwt/jwt/v4"
	user2 "golang-basic/domain"
	"golang-basic/pkg/jwt"
	"golang-basic/pkg/response"
	"golang-basic/usecase"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BookController struct {
	Usecase usecase.BookUsecase
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{
		Usecase: usecase.NewBookUsecase(db),
	}
}
func (c *BookController) GetBooks(ctx *gin.Context) {
	token, _ := jwt.GetToken(ctx)
	claims := token.Claims.(jwt2.MapClaims)

	userId := int(claims["id"].(float64))

	books, err := c.Usecase.GetBooks(ctx, userId)
	if checkError(ctx, err) {
		return
	} else {
		response.OK(ctx, books)
	}
}
func (c *BookController) GetBook(ctx *gin.Context) {
	token, _ := jwt.GetToken(ctx)
	claims := token.Claims.(jwt2.MapClaims)

	userId := int(claims["id"].(float64))

	id, error := strconv.Atoi(ctx.Param("id"))

	if error != nil {
		return
	}
	book, err := c.Usecase.GetBook(ctx, userId, id)
	if checkError(ctx, err) {
		return
	} else {
		response.OK(ctx, book)
	}
}
func (c *BookController) CreateBook(ctx *gin.Context) {
	token, _ := jwt.GetToken(ctx)
	claims := token.Claims.(jwt2.MapClaims)

	userId := int(claims["id"].(float64))
	book := user2.CreateBookReq{}
	if !bindJSON(ctx, &book) {
		return
	}

	err := c.Usecase.CreateBook(ctx, book, userId)
	if checkError(ctx, err) {
		return
	}

	response.WithStatusCode(ctx, http.StatusCreated, nil)
}
func (c *BookController) EditBook(ctx *gin.Context) {
	token, _ := jwt.GetToken(ctx)
	claims := token.Claims.(jwt2.MapClaims)

	userId := int(claims["id"].(float64))
	id, error := strconv.Atoi(ctx.Param("id"))

	if error != nil {
		return
	}
	book := user2.CreateBookReq{}
	if !bindJSON(ctx, &book) {
		return
	}

	err := c.Usecase.EditBook(ctx, book, userId, id)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, nil)
}
func (c *BookController) DeleteBook(ctx *gin.Context) {
	token, _ := jwt.GetToken(ctx)
	claims := token.Claims.(jwt2.MapClaims)

	userId := int(claims["id"].(float64))
	id, error := strconv.Atoi(ctx.Param("id"))

	if error != nil {
		return
	}

	err := c.Usecase.DeleteBook(ctx, userId, id)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, nil)
}
