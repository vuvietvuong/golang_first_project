package usecase

import (
	"context"
	"golang_first_pj/domain"
	"golang_first_pj/pkg/crypto"
	"golang_first_pj/pkg/jwt"
	"golang_first_pj/repository"
	"gorm.io/gorm"
	"os"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

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
		jwt, err := jwt.GenerateJWT(user)
		if err != nil {
			return nil, err
		}
		return &jwt, nil
	}
	return nil, nil
}
