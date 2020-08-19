package theater

import (
	"context"
	"github.com/CastyLab/api.server/app/components"
	"github.com/CastyLab/api.server/app/models"
	"github.com/CastyLab/api.server/grpc"
	"github.com/CastyLab/grpc.proto/proto"
	"github.com/MrJoshLab/go-respond"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"log"
	"net/http"
	"time"
)

func GetMediaSources(ctx *gin.Context) {

	var (
		mediaSources = make([]*proto.MediaSource, 0)
		token = ctx.Request.Header.Get("Authorization")
		mCtx, _ = context.WithTimeout(ctx, 20 * time.Second)
	)

	response, err := grpc.TheaterServiceClient.GetMediaSources(mCtx, &proto.MediaSourceAuthRequest{
		AuthRequest: &proto.AuthenticateRequest{
			Token: []byte(token),
		},
	})

	if err != nil {
		if code, result, ok := components.ParseGrpcErrorResponse(err); !ok {
			ctx.JSON(code, result)
			return
		}
	}

	if response.Result != nil {
		mediaSources = response.Result
	}

	ctx.JSON(respond.Default.SetStatusText("success").
		SetStatusCode(http.StatusOK).
		RespondWithResult(mediaSources))
	return
}

func DeleteMediaSource(ctx *gin.Context)  {

	var (
		rules = govalidator.MapData{
			"source_id": []string{"required"},
		}
		opts = govalidator.Options{
			Request:         ctx.Request,
			Rules:           rules,
			RequiredDefault: true,
		}
		token = ctx.Request.Header.Get("Authorization")
	)

	if validate := govalidator.New(opts).Validate(); validate.Encode() != "" {
		validations := components.GetValidationErrorsFromGoValidator(validate)
		ctx.JSON(respond.Default.ValidationErrors(validations))
		return
	}

	response, err := grpc.TheaterServiceClient.RemoveMediaSource(ctx, &proto.MediaSourceRemoveRequest{
		AuthRequest: &proto.AuthenticateRequest{
			Token: []byte(token),
		},
		MediaSourceId: ctx.PostForm("source_id"),
	})

	if err != nil {
		if code, result, ok := components.ParseGrpcErrorResponse(err); !ok {
			ctx.JSON(code, result)
			return
		}
	}

	if response.Code != http.StatusOK {
		ctx.JSON(respond.Default.SetStatusCode(http.StatusBadRequest).
			SetStatusText("failed").
			RespondWithMessage("Could not delete media source!"))
		return
	}

	ctx.JSON(respond.Default.UpdateSucceeded())
	return
}

func SelectNewMediaSource(ctx *gin.Context)  {

	var (
		rules = govalidator.MapData{
			"source_id": []string{"required"},
		}
		opts = govalidator.Options{
			Request:         ctx.Request,
			Rules:           rules,
			RequiredDefault: true,
		}
		token = ctx.Request.Header.Get("Authorization")
	)

	if validate := govalidator.New(opts).Validate(); validate.Encode() != "" {

		validations := components.GetValidationErrorsFromGoValidator(validate)
		ctx.JSON(respond.Default.ValidationErrors(validations))
		return
	}

	response, err := grpc.TheaterServiceClient.SelectMediaSource(ctx, &proto.MediaSourceAuthRequest{
		AuthRequest: &proto.AuthenticateRequest{
			Token: []byte(token),
		},
		Media: &proto.MediaSource{
			Id: ctx.PostForm("source_id"),
		},
	})

	if err != nil {
		if code, result, ok := components.ParseGrpcErrorResponse(err); !ok {
			ctx.JSON(code, result)
			return
		}
	}

	ctx.JSON(respond.Default.Succeed(response.Result[0]))
	return

}

func ParseMediaSourceUri(ctx *gin.Context)  {

	var (
		rules = govalidator.MapData{
			"source": []string{"required", "media_source_uri"},
		}
		opts = govalidator.Options{
			Request:         ctx.Request,
			Rules:           rules,
			RequiredDefault: true,
		}
	)

	if validate := govalidator.New(opts).Validate(); validate.Encode() != "" {
		validations := components.GetValidationErrorsFromGoValidator(validate)
		ctx.JSON(respond.Default.ValidationErrors(validations))
		return
	}

	mediaSource := models.NewMediaSource(ctx.PostForm("source"))

	err := mediaSource.Parse()
	if err != nil {
		log.Println(err)
		ctx.JSON(respond.Default.
			SetStatusText("Failed!").
			SetStatusCode(http.StatusBadRequest).
			RespondWithMessage("Could not parse data. Please try again later!"))
		return
	}

	if mediaSource.IsUnknown() {
		ctx.JSON(respond.Default.
			SetStatusText("Failed!").
			SetStatusCode(http.StatusBadRequest).
			RespondWithMessage("Invalid media source type!"))
		return
	}

	if mediaSource.IsTorrent() {

		ctx.JSON(respond.Default.SetStatusCode(http.StatusNotAcceptable).
			SetStatusText("Failed").
			RespondWithMessage("Torrent links are not available yet!"))
		return
	}

	ctx.JSON(respond.Default.Succeed(mediaSource.Proto()))
	return
}

func AddNewMediaSource(ctx *gin.Context) {

	var (
		token = ctx.Request.Header.Get("Authorization")
		rules = govalidator.MapData{
			"source": []string{"required", "media_source_uri"},
		}
		opts = govalidator.Options{
			Request:         ctx.Request,
			Rules:           rules,
			RequiredDefault: true,
		}
	)

	if validate := govalidator.New(opts).Validate(); validate.Encode() != "" {
		validations := components.GetValidationErrorsFromGoValidator(validate)
		ctx.JSON(respond.Default.ValidationErrors(validations))
		return
	}

	mediaSource := models.NewMediaSource(ctx.PostForm("source"))

	err := mediaSource.Parse()
	if err != nil {
		ctx.JSON(respond.Default.
			SetStatusText("Failed!").
			SetStatusCode(http.StatusBadRequest).
			RespondWithMessage("Could not parse data. Please try again later!"))
		return
	}

	if mediaSource.IsUnknown() {
		ctx.JSON(respond.Default.
			SetStatusText("Failed!").
			SetStatusCode(http.StatusBadRequest).
			RespondWithMessage("Invalid media source type!"))
		return
	}

	// check for index files inside of torrent
	if mediaSource.IsTorrent() {
		ctx.JSON(respond.Default.SetStatusCode(http.StatusNotAcceptable).
			SetStatusText("Failed").
			RespondWithMessage("Torrent links are not available yet!"))
		return
	}

	protoMsg := mediaSource.Proto()

	if title := ctx.PostForm("title"); title != "" {
		protoMsg.Title = title
	}

	response, err := grpc.TheaterServiceClient.AddMediaSource(ctx, &proto.MediaSourceAuthRequest{
		AuthRequest: &proto.AuthenticateRequest{
			Token: []byte(token),
		},
		Media: protoMsg,
	})

	if err != nil {
		if code, result, ok := components.ParseGrpcErrorResponse(err); !ok {
			ctx.JSON(code, result)
			return
		}
	}

	ctx.JSON(respond.Default.Succeed(response.Result[0]))
	return
}