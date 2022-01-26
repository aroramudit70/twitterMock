package service

import (
	"twitterMock/api/generated/models"
	"twitterMock/dao"
	// "logur.dev/adapter/zap"
)

type ServiceStruct struct {
	dao dao.DaoInterface
	// logger *zap.logger
}

type ServiceInterface interface {
	SignUp(params *models.Profile) (*string, *models.ErrResponse)
}

func NewServiceInterface() ServiceInterface {
	return &ServiceStruct{}
}
func (s *ServiceStruct) SignUp(params *models.Profile) (*string, *models.ErrResponse) {
	return nil, nil
}
