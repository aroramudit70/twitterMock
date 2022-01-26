package dao

import (
	"twitterMock/api/generated/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type DaoStruct struct {
	db mongo.
}
type DaoInterface interface {
	userSignUp(models.Profile) error
}

func NewDaoInterface() DaoInterface {
	return &DaoStruct{}
}

func (d *DaoStruct) userSignUp(models.Profile) error {

	return nil

}
