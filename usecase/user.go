package usecase

import (
	"context"
	user2 "golang-basic/domain"
	"golang-basic/domain/model"
	"golang-basic/pkg/crypto"
	"golang-basic/repository"

	"gorm.io/gorm"
)

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUser(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, u user2.CreateUserReq) error
}

type userUsecase struct {
	UserRepo repository.UserRepo
}

func NewUserUsecase(r repository.UserRepo) UserUsecase {
	return &userUsecase{
		UserRepo: r,
	}
}

func (uc *userUsecase) GetUsers(ctx context.Context) (users []model.User, err error) {
	return
}

func (uc *userUsecase) GetUser(ctx context.Context, id int) (user *model.User, err error) {
	return
}

func (uc *userUsecase) CreateUser(ctx context.Context, u user2.CreateUserReq) (err error) {
	hash := crypto.HashString(u.Password)
	u.Password = hash
	_, err = uc.UserRepo.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return
}
