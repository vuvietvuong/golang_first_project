package response

import (
	"errors"
	"golang-basic/config"
	"golang-basic/domain"
	"golang-basic/pkg/apierrors"
	"golang-basic/pkg/i18n"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Error(ctx *gin.Context, err error) {
	var debugMsg string
	if gin.Mode() != gin.ReleaseMode {
		debugMsg = err.Error()
	}
	errType := apierrors.ErrType(err)

	ctx.JSON(errType.HTTPCode(), domain.ErrorResponse{
		Code:         errType.Code(),
		DebugMessage: debugMsg,
		ErrorDetails: errorDetails(err, ctx.Query("locale")),
	})
}

func errorDetails(err error, locale string) (details []domain.ErrorDetail) {
	// Set default locale
	if locale == "" {
		locale = config.Locale
	}
	var vErrs validator.ValidationErrors
	if errors.As(err, &vErrs) {
		trans := i18n.GetTrans(locale)
		for _, err := range vErrs {
			details = append(details, domain.ErrorDetail{
				Field:        err.Field(),
				ErrorCode:    err.Tag(),
				ErrorMessage: err.Translate(trans),
			})
		}
	}
	return
}
