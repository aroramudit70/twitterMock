package handler

import (
	"twitterMock/api/generated/restapi/operations/user"
	"twitterMock/service"
	"twitterMock/util"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type LoginHandler struct {
	Service service.ServiceInterface
	logger  *zap.Logger
}

func NewLoginHandler(Service service.ServiceInterface, logger *zap.Logger) *LoginHandler {
	return &LoginHandler{
		Service: Service,
		logger:  logger,
	}
}

func (h *LoginHandler) Handle(params user.LoginUserParams) middleware.Responder {
	h.logger.Info("Inside Follow handler")
	res, err := h.Service.Login(params)
	if err != nil && err.Code == util.NotFoundCode {
		return user.NewLoginUserNotFound().WithPayload(err)
	} else if err != nil {
		return user.NewLoginUserInternalServerError().WithPayload(err)
	}

	return user.NewLoginUserOK().WithPayload(*res)
}
