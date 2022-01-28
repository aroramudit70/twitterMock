package util

import "twitterMock/api/generated/models"

func GetError(code string, msg string) *models.ErrResponse {
	return &models.ErrResponse{
		Code:    code,
		Message: msg,
	}
}

func GetErrorMapped(msg string) *models.ErrResponse {

	switch msg {
	case MongoISEError:
		return GetError(ISECode, msg)
	case DuplicateCodeMsg:
		return GetError(DuplicateCode, msg)
	case NotFoundMsg:
		return GetError(NotFoundCode, msg)
	case UnAuthorizedMsg:
		return GetError(UnAuthorized, msg)
	default:
		return GetError(ISECode, msg)
	}

}
