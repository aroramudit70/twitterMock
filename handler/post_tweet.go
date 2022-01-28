package handler

import (
	"twitterMock/api/generated/restapi/operations/operation"
	"twitterMock/service"
	"twitterMock/util"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type TweetHandler struct {
	Service service.ServiceInterface
	logger  *zap.Logger
}

func NewTweetHandler(Service service.ServiceInterface, logger *zap.Logger) *TweetHandler {
	return &TweetHandler{
		Service: Service,
		logger:  logger,
	}
}

func (h *TweetHandler) Handle(params operation.PostTweetParams, principal interface{}) middleware.Responder {
	h.logger.Info("Inside Tweet handler")
	h.logger.Debug("Header value", zap.Any("Header Details: ", params.HTTPRequest.Header.Get("Authorization")))
	UserName, err1 := util.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err1 != nil {
		return operation.NewPostTweetInternalServerError().WithPayload(util.GetErrorMapped(util.UnAuthorizedMsg))
	}
	params.Body.UserName = UserName.(string)
	res, err := h.Service.Tweet(*params.Body)
	if err != nil && err.Code == util.NotFoundCode {
		return operation.NewPostTweetNotFound().WithPayload(err)
	} else if err != nil && err.Code == util.UnAuthorized {
		return operation.NewPostTweetUnauthorized().WithPayload(err)
	} else if err != nil && err.Code == util.ISECode {
		return operation.NewPostTweetInternalServerError().WithPayload(err)
	}
	return operation.NewPostTweetOK().WithPayload(*res)
}
