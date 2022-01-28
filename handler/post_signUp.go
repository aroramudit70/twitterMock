package handler

import (
	"twitterMock/api/generated/restapi/operations/user"
	"twitterMock/service"
	"twitterMock/util"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type SignUpHandler struct {
	Service service.ServiceInterface
	logger  *zap.Logger
}

func NewSignUpHandler(Service service.ServiceInterface, logger *zap.Logger) *SignUpHandler {
	return &SignUpHandler{
		Service: Service,
		logger:  logger,
	}
}

func (h *SignUpHandler) Handle(params user.SignupParams) middleware.Responder {
	h.logger.Info("Inside SignUp handler")
	res, err := h.Service.SignUp(*params.Body)
	if err != nil && err.Code == util.ISECode {
		return user.NewSignupInternalServerError().WithPayload(err)
	} else if err != nil && err.Code == util.DuplicateCode {
		return user.NewSignupConflict().WithPayload(err)
	}
	return user.NewSignupOK().WithPayload(*res)
}
