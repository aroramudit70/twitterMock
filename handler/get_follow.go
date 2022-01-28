package handler

import (
	"twitterMock/api/generated/restapi/operations/operation"
	"twitterMock/service"
	"twitterMock/util"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type FollowHandler struct {
	Service service.ServiceInterface
	logger  *zap.Logger
}

func NewFollowHandler(Service service.ServiceInterface, logger *zap.Logger) *FollowHandler {
	return &FollowHandler{
		Service: Service,
		logger:  logger,
	}
}

func (h *FollowHandler) Handle(params operation.FollowParams, principal interface{}) middleware.Responder {
	h.logger.Info("Inside Follow handler")
	UserName, err1 := util.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err1 != nil {
		return operation.NewPostTweetInternalServerError().WithPayload(util.GetErrorMapped(util.UnAuthorizedMsg))
	}
	userName := UserName.(string)
	res, err := h.Service.Follow(userName, params.UserHandle)
	if err != nil && err.Code == util.NotFoundCode {
		return operation.NewFollowNotFound().WithPayload(err)
	} else if err != nil && err.Code == util.UnAuthorized {
		return operation.NewFollowUnauthorized().WithPayload(err)
	} else if err != nil {
		return operation.NewFollowInternalServerError().WithPayload(err)
	}
	return operation.NewFollowOK().WithPayload(*res)
}
