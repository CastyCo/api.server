package user

import (
	"context"
	"github.com/CastyLab/api.server/app/components"
	"github.com/CastyLab/api.server/app/components/recaptcha"
	"github.com/CastyLab/api.server/grpc"
	"github.com/CastyLab/grpc.proto"
	"github.com/CastyLab/grpc.proto/messages"
	"github.com/MrJoshLab/go-respond"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"time"
)

// Create a new user
func Create(ctx *gin.Context)  {

	var (
		rules = govalidator.MapData{
			"fullname":                 []string{"min:4", "max:30"},
			"password":                 []string{"required", "min:4", "max:30"},
			"password_confirmation":    []string{"required", "min:4", "max:30"},
			"username":                 []string{"required", "between:3,20"},
			"email":                    []string{"required", "email"},
			"g-recaptcha-response":     []string{"required"},
		}
		opts = govalidator.Options{
			Request:         ctx.Request,
			Rules:           rules,
			RequiredDefault: true,
		}
		mCtx, _ = context.WithTimeout(ctx, 10 * time.Second)
	)

	if validate := govalidator.New(opts).Validate(); validate.Encode() != "" {

		validations := components.GetValidationErrorsFromGoValidator(validate)
		ctx.JSON(respond.Default.ValidationErrors(validations))
		return
	}

	if ctx.PostForm("password") != ctx.PostForm("password_confirmation") {
		ctx.JSON(respond.Default.ValidationErrors(map[string] interface{} {
			"password": []string {
				"Passwords are not match!",
			},
		}))
		return
	}

	if success, err := recaptcha.Verify(ctx); err != nil || !success {
		ctx.JSON(respond.Default.ValidationErrors(map[string] interface{} {
			"recaptcha": []string {
				"Captcha is invalid!",
			},
		}))
		return
	}

	response, err := grpc.UserServiceClient.CreateUser(mCtx, &proto.CreateUserRequest{
		User: &messages.User{
			Fullname: ctx.PostForm("fullname"),
			Username: ctx.PostForm("username"),
			Email:    ctx.PostForm("email"),
			Password: ctx.PostForm("password"),
		},
	})

	if response != nil && response.Code == 420 {

		valErrs := make(map[string] interface{})
		for _, verr := range response.ValidationError {
			valErrs[verr.Field] = verr.Errors
		}

		ctx.JSON(respond.Default.ValidationErrors(valErrs))
		return
	}

	if err != nil || response == nil || response.Code != http.StatusOK {
		ctx.JSON(respond.Default.SetStatusCode(420).
			SetStatusText("failed").
			RespondWithMessage("Could not create user."))
		return
	}

	ctx.JSON(respond.Default.Succeed(map[string] interface{} {
		"token": string(response.Token),
		"refreshed_token": string(response.Token),
		"type": "bearer",
	}))
	return
}
