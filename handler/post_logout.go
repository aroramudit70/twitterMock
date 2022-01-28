package handler

import (
	"twitterMock/api/generated/restapi/operations/user"
	"twitterMock/service"
	"twitterMock/util"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type LogoutHandler struct {
	Service service.ServiceInterface
	logger  *zap.Logger
}

func NewLogoutHandler(Service service.ServiceInterface, logger *zap.Logger) *LogoutHandler {
	return &LogoutHandler{
		Service: Service,
		logger:  logger,
	}
}

func (h *LogoutHandler) Handle(params user.LogoutUserParams, principal interface{}) middleware.Responder {
	h.logger.Info("Inside Feed handler")
	UserName, err1 := util.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err1 != nil {
		return user.NewLogoutUserDefault(500)
	}
	userName := UserName.(string)
	err := h.Service.Logout(userName)
	if err != nil && err.Code == util.ISECode {
		return user.NewLogoutUserDefault(500)
	} else if err != nil && err.Code == util.NotFoundCode {
		return user.NewLogoutUserDefault(400)
	}
	return user.NewLogoutUserDefault(200)
}
