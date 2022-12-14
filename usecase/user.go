package usecase

import (
	"context"
	user2 "golang_first_pj/domain"
	"golang_first_pj/domain/model"
	"golang_first_pj/pkg/crypto"
	"golang_first_pj/repository"

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

func NewUserUsecase(db *gorm.DB) UserUsecase {
	return &userUsecase{
		UserRepo: repository.NewUserRepo(db),
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
