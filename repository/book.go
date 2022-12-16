package repository

import (
	"context"
	"fmt"
	user2 "golang-basic/domain"
	"golang-basic/domain/model"
	"gorm.io/gorm"
)

type BookRepo interface {
	GetBooks(ctx context.Context, userId int) ([]model.Book, error)
	GetBook(ctx context.Context, userId int, id int) (*model.Book, error)
	CreateBook(ctx context.Context, u user2.CreateBookReq, userId int) (*model.Book, error)
	EditBook(ctx context.Context, u user2.CreateBookReq, userId int, id int) (*model.Book, error)
	DeleteBook(ctx context.Context, userId int, id int) (*model.Book, error)
}

type bookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) BookRepo {
	return &bookRepo{DB: db}
}
func (r *bookRepo) GetBooks(ctx context.Context, userId int) (books []model.Book, err error) {
	result := r.DB.Where("user_id= ?", userId).Find(&books)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.GetBooks] failed: %w", result.Error)
		return nil, err
	}
	return books, nil
}
func (r *bookRepo) GetBook(ctx context.Context, userId int, id int) (book *model.Book, err error) {
	result := r.DB.Where("user_id= ? and id=?", userId, id).First(book)
	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.GetBook] failed: %w", result.Error)
		return nil, err
	}
	return book, nil
}
func (r *bookRepo) CreateBook(ctx context.Context, u user2.CreateBookReq, userId int) (book *model.Book, err error) {
	book = &model.Book{
		Name:   u.Name,
		UserId: userId,
	}
	result := r.DB.Create(book)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.CreateBook] failed: %w", result.Error)
		return nil, err
	}
	return
}
func (r *bookRepo) EditBook(ctx context.Context, u user2.CreateBookReq, userId int, id int) (book *model.Book, err error) {
	book = &model.Book{
		Name: u.Name,
	}
	result := r.DB.Where("user_id= ? and id=?", userId, id).Updates(book)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.EditBook] failed: %w", result.Error)
		return nil, err
	}
	return
}
func (r *bookRepo) DeleteBook(ctx context.Context, userId int, id int) (book *model.Book, err error) {
	book = &model.Book{}
	result := r.DB.Where("user_id= ? and id=?", userId, id).Delete(book)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Book.GetBook] failed: %w", result.Error)
		return nil, err
	}
	return
}
