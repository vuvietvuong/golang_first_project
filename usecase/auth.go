package usecase

import (
	"context"
	"golang-basic/domain"
	"golang-basic/pkg/crypto"
	"golang-basic/pkg/jwt"
	"golang-basic/repository"

	"gorm.io/gorm"
)

type AuthUsecase interface {
	Login(ctx context.Context, u domain.Auth) (*string, error)
}
type authUsecase struct {
	UserRepo repository.UserRepo
}

func NewAuthUsecase(db *gorm.DB) AuthUsecase {
	return &authUsecase{
		UserRepo: repository.NewUserRepo(db),
	}
}
func (uc *authUsecase) Login(ctx context.Context, u domain.Auth) (*string, error) {
	user, err := uc.UserRepo.GetUserByGmail(ctx, u.Email)

	if err != nil {
		return nil, err
	}
	match := crypto.DoMatch(user.Password, u.Password)
	if match {
		generateJWT, err := jwt.GenerateJWT(user)
		if err != nil {
			return nil, err
		}
		return &generateJWT, nil
	}
	return nil, nil
}
