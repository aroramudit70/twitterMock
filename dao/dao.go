package dao

import (
	"context"
	"errors"
	"time"
	"twitterMock/api/generated/models"

	"twitterMock/util"

	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type DaoStruct struct {
	db     *mongo.Database
	logger *zap.Logger
}
type DaoInterface interface {
	UserSignUp(models.Profile) error
	GetUser(UserName string) (*models.Profile, error)
	Follow(UserName, Handle string) error
	GetFeed(UserName string) (*models.Feed, error)
	Tweet(params models.Tweet) error
	UpdateLogin(userName string, state bool) error
}

func NewDaoInterface(db *mongo.Database, logger *zap.Logger) DaoInterface {
	return &DaoStruct{
		db:     db,
		logger: logger,
	}
}

func (d *DaoStruct) UserSignUp(params models.Profile) error {
	d.logger.Info("Inside UserSignUp function in Dao")
	_, err := d.db.Collection(util.CollectionUser).InsertOne(context.TODO(), params)
	if err != nil {
		d.logger.Error("Error occoured while saving data : ", zap.String("Error", err.Error()))
		return err
	}
	d.logger.Info("UserSignUP Dao Success")
	return nil
}
func (d *DaoStruct) GetUser(UserName string) (*models.Profile, error) {
	d.logger.Info("Inside GetUser function in Dao")
	filter := bson.M{
		"username": UserName,
	}
	d.logger.Debug("Filters for DB query: ", zap.Any("Filters", filter))
	res := &models.Profile{}
	err := d.db.Collection(util.CollectionUser).FindOne(context.TODO(), filter).Decode(res)
	if err != nil {
		d.logger.Error("Error occoured while DB Transaction : ", zap.String("Error", err.Error()))
		return nil, err
	}
	d.logger.Info("GetUser Dao Success")
	return res, nil
}

func (d *DaoStruct) Follow(UserName, Handle string) error {
	d.logger.Info("Inside Follow function in Dao")
	filter := bson.M{
		"username": UserName,
	}
	update := bson.M{
		"$push": bson.M{"followinglist": Handle},
	}
	d.logger.Debug("Filters for DB query: ", zap.Any("Filters", filter))
	_, err := d.db.Collection(util.CollectionUser).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		d.logger.Error("Error occoured while DB Transaction : ", zap.String("Error", err.Error()))
		return err
	}
	d.logger.Info("Follow Dao Success")
	return nil
}

func (d *DaoStruct) GetFeed(UserName string) (*models.Feed, error) {
	d.logger.Info("Inside GetFeed function in Dao")
	user, err := d.GetUser(UserName)
	if err != nil {
		d.logger.Error("Error occoured while DB Transaction : ", zap.String("Error", err.Error()))
		return nil, err
	}
	if !*user.LoggedIn {
		return nil, errors.New(util.UnAuthorizedMsg)
	}

	filter := bson.M{
		"username": bson.M{
			"$in": user.FollowingList,
		},
	}
	d.logger.Debug("Filters for DB query: ", zap.Any("Filters", filter))
	res := models.Feed{}
	cur, err := d.db.Collection(util.CollectionTweet).Find(context.TODO(), filter)
	if err != nil {
		d.logger.Error("Error occoured while DB Transaction : ", zap.String("Error", err.Error()))
		return nil, errors.New(util.MongoISEError)
	}
	for cur.Next(context.TODO()) {
		model := &models.Tweet{}
		err := cur.Decode(model)
		if err != nil {
			d.logger.Error("Error occoured while DB Transaction : ", zap.String("Error", err.Error()))
			return nil, err
		}
		res = append(res, model)
	}
	d.logger.Info("GetFeed Dao Success")
	return &res, nil
}
func (d *DaoStruct) Tweet(params models.Tweet) error {
	d.logger.Info("Inside Tweet function in Dao")
	params.Time = strfmt.DateTime(time.Now())
	_, err := d.db.Collection(util.CollectionTweet).InsertOne(context.TODO(), params)
	if err != nil {
		d.logger.Error("Error occoured while DB Transaction : ", zap.String("Error", err.Error()))
		return err
	}
	d.logger.Info("Tweet Dao Success")
	return nil
}

func (d *DaoStruct) UpdateLogin(userName string, state bool) error {
	d.logger.Info("Inside UpdateLogin function in Dao")
	filter := bson.M{
		"username": userName,
	}
	update := bson.M{
		"$set": bson.M{
			"loggedin": state},
	}
	d.logger.Debug("Filters for DB query: ", zap.Any("Filters", filter))
	_, err := d.db.Collection(util.CollectionUser).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		d.logger.Error("Error occoured while DB Transaction : ", zap.String("Error", err.Error()))
		return err
	}
	d.logger.Info("Follow UpdateLogin Success")
	return nil
}
