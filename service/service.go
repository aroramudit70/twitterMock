package service

import (
	"twitterMock/api/generated/models"
	"twitterMock/api/generated/restapi/operations/user"
	"twitterMock/dao"
	"twitterMock/util"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var Success = "Success"
var boolean bool

type ServiceStruct struct {
	Dao    dao.DaoInterface
	logger *zap.Logger
}

type ServiceInterface interface {
	SignUp(params models.Profile) (*string, *models.ErrResponse)
	Follow(UserName, Handle string) (*string, *models.ErrResponse)
	GetFeed(UserName string) (*models.Feed, *models.ErrResponse)
	Tweet(params models.Tweet) (*string, *models.ErrResponse)
	Login(params user.LoginUserParams) (*string, *models.ErrResponse)
	Logout(userName string) *models.ErrResponse
}

func NewServiceInterface(Dao dao.DaoInterface, logger *zap.Logger) ServiceInterface {
	return &ServiceStruct{
		Dao:    Dao,
		logger: logger,
	}
}
func (s *ServiceStruct) SignUp(params models.Profile) (*string, *models.ErrResponse) {
	s.logger.Info("Inside Signup function in Service")
	password := []byte(*params.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, util.GetErrorMapped(err.Error())
	}
	pass := string(hashedPassword)
	params.Password = &pass

	_, err = s.Dao.GetUser(*params.UserName)
	if err != nil && err.Error() == mongo.ErrNoDocuments.Error() {
		params.FollowingList = make([]string, 0)
		params.LoggedIn = &boolean
		err2 := s.Dao.UserSignUp(params)
		if err2 != nil {
			s.logger.Error("Error occoured while calling dao : ", zap.String("Error", err2.Error()))
			return nil, util.GetErrorMapped(util.MongoISEError)
		}
		s.logger.Info("SignUp Service Success")
		return &Success, nil
	} else if err != nil {
		s.logger.Error("Error occoured while calling dao : ", zap.String("Error", err.Error()))
		return nil, util.GetErrorMapped(util.MongoISEError)
	}
	s.logger.Error("Error occoured while calling dao : User Already Exist")
	return nil, util.GetErrorMapped(util.DuplicateCodeMsg)
}

func (s *ServiceStruct) Follow(UserName, Handle string) (*string, *models.ErrResponse) {
	s.logger.Info("Inside Follow function in Service")
	res, err := s.Dao.GetUser(UserName)
	if err != nil && err.Error() == mongo.ErrNoDocuments.Error() {
		s.logger.Error("Error occoured while calling dao : ", zap.String("Error", err.Error()))
		return nil, util.GetErrorMapped(util.NotFoundMsg)
	} else if err != nil {
		s.logger.Error("Error occoured while calling dao : ", zap.String("Error", err.Error()))
		return nil, util.GetErrorMapped(util.MongoISEError)
	}
	if !*res.LoggedIn {
		return nil, util.GetErrorMapped(util.UnAuthorizedMsg)
	}
	err = s.Dao.Follow(UserName, Handle)
	if err != nil {
		s.logger.Error("Error occoured while calling dao : ", zap.String("Error", err.Error()))
		return nil, util.GetErrorMapped(util.MongoISEError)
	}
	s.logger.Info("Feed Service Success")
	return &Success, nil
}
func (s *ServiceStruct) GetFeed(UserName string) (*models.Feed, *models.ErrResponse) {
	s.logger.Info("Inside GetFeed function in Service")
	res, err := s.Dao.GetFeed(UserName)
	if err != nil && err.Error() == mongo.ErrNoDocuments.Error() {
		return nil, util.GetErrorMapped(util.DuplicateCodeMsg)
	} else if err != nil && err.Error() == util.MongoISEError {
		return nil, util.GetErrorMapped(util.MongoISEError)
	} else if err != nil {
		return nil, util.GetErrorMapped(err.Error())
	}
	s.logger.Info("GetFeed Service Success")
	return res, nil
}

func (s *ServiceStruct) Tweet(params models.Tweet) (*string, *models.ErrResponse) {
	s.logger.Info("Inside Tweet function in Service")
	res, err := s.Dao.GetUser(params.UserName)
	if err != nil && err.Error() == mongo.ErrNoDocuments.Error() {
		s.logger.Error("Error occoured while calling dao : ", zap.String("Error", err.Error()))
		return nil, util.GetErrorMapped(util.NotFoundMsg)
	} else if err != nil {
		s.logger.Error("Error occoured while calling dao : ", zap.String("Error", err.Error()))
		return nil, util.GetErrorMapped(util.MongoISEError)
	}
	if !*res.LoggedIn {
		return nil, util.GetErrorMapped(util.UnAuthorizedMsg)
	}
	err = s.Dao.Tweet(params)
	if err != nil {
		s.logger.Error("Error occoured while calling dao : ", zap.String("Error", err.Error()))
		return nil, util.GetErrorMapped(util.MongoISEError)
	}
	s.logger.Info("Tweet Service Success")
	return &Success, nil
}
func (s *ServiceStruct) Login(params user.LoginUserParams) (*string, *models.ErrResponse) {
	s.logger.Info("Inside Login function in Service")
	res, err := s.Dao.GetUser(params.Username)
	if err != nil && err.Error() == mongo.ErrNoDocuments.Error() {
		s.logger.Error("Error occoured while calling dao : ", zap.String("Error", err.Error()))
		return nil, util.GetErrorMapped(util.NotFoundMsg)
	} else if err != nil {
		s.logger.Error("Error occoured while calling dao : ", zap.String("Error", err.Error()))
		return nil, util.GetErrorMapped(util.MongoISEError)
	}
	err = bcrypt.CompareHashAndPassword([]byte(*res.Password), []byte(params.Password))
	if err != nil {
		s.logger.Error("Error occoured while comparing password : ", zap.String("Error", err.Error()))
		return nil, util.GetErrorMapped(util.UnAuthorizedMsg)
	}
	token, err := util.GenerateJWT(res.EmailID, *res.Name, *res.UserName)
	if err != nil {
		s.logger.Error("Error occoured while comparing password : ", zap.String("Error", err.Error()))
		return nil, util.GetErrorMapped(util.UnAuthorizedMsg)
	}
	err = s.Dao.UpdateLogin(*res.UserName, true)
	if err != nil {
		s.logger.Error("Error occoured while calling dao : ", zap.String("Error", err.Error()))
		return nil, util.GetErrorMapped(util.MongoISEError)
	}
	s.logger.Info("LOGIN Service Success")
	return &token, nil
}

func (s *ServiceStruct) Logout(userName string) *models.ErrResponse {
	err := s.Dao.UpdateLogin(userName, false)
	if err != nil {
		s.logger.Error("Error occoured while calling dao : ", zap.String("Error", err.Error()))
		return util.GetErrorMapped(util.MongoISEError)
	}
	s.logger.Info("LOGOUT Service Success")
	return nil
}
