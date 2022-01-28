package handler

import (
	"twitterMock/api/generated/restapi/operations/operation"
	"twitterMock/service"
	"twitterMock/util"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type FeedHandler struct {
	Service service.ServiceInterface
	logger  *zap.Logger
}

func NewFeedHandler(Service service.ServiceInterface, logger *zap.Logger) *FeedHandler {
	return &FeedHandler{
		Service: Service,
		logger:  logger,
	}
}

func (h *FeedHandler) Handle(params operation.GetFeedParams, principal interface{}) middleware.Responder {
	h.logger.Info("Inside Feed handler")
	UserName, err1 := util.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err1 != nil {
		return operation.NewPostTweetInternalServerError().WithPayload(util.GetErrorMapped(util.UnAuthorizedMsg))
	}
	userName := UserName.(string)
	res, err := h.Service.GetFeed(userName)
	if err != nil && err.Code == util.NotFoundCode {
		return operation.NewGetFeedNotFound().WithPayload(err)
	} else if err != nil && err.Code == util.UnAuthorized {
		return operation.NewGetFeedUnauthorized().WithPayload(err)
	} else if err != nil {
		return operation.NewGetFeedInternalServerError().WithPayload(err)
	}
	return operation.NewGetFeedOK().WithPayload(*res)
}
