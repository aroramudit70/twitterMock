package module

import (
	"net/http"
	"twitterMock/api/generated/restapi"
	"twitterMock/api/generated/restapi/operations"
	"twitterMock/dao"
	"twitterMock/handler"
	"twitterMock/service"

	"twitterMock/util"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var HandlerModule = wire.NewSet(
	NewAPI,
	handler.NewFeedHandler,
	handler.NewFollowHandler,
	handler.NewSignUpHandler,
	handler.NewTweetHandler,
	handler.NewLoginHandler,
	handler.NewLogoutHandler,
	NewHandler,
)
var ServiceModule = wire.NewSet(
	service.NewServiceInterface,
)

var DaoModulle = wire.NewSet(
	dao.NewDaoInterface,
)

// NewAPI ...
func NewAPI(logger *zap.Logger) (*operations.TwitterMockAPI, error) {
	// Load the OpenAPI2 spec
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}

	api := operations.NewTwitterMockAPI(swaggerSpec)
	api.ServeError = errors.ServeError
	api.Logger = func(s string, i ...interface{}) {
		if len(i) == 0 {
			logger.Sugar().Info(s)
		} else {
			logger.Sugar().Infof(s, i)
		}
	}
	// api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	return api, nil
}

func NewHandler(
	logger *zap.Logger,
	api *operations.TwitterMockAPI,
	feed *handler.FeedHandler,
	follow *handler.FollowHandler,
	signup *handler.SignUpHandler,
	tweet *handler.TweetHandler,
	signin *handler.LoginHandler,
	signout *handler.LogoutHandler,

) (http.Handler, error) {
	// Handlers
	api.OperationFollowHandler = follow
	api.OperationGetFeedHandler = feed
	api.OperationPostTweetHandler = tweet
	api.UserSignupHandler = signup
	api.UserLoginUserHandler = signin
	api.UserLogoutUserHandler = signout
	api.BearerAuth = util.ValidateHeader
	// Middlewares
	handler := api.Serve(nil)
	return handler, nil
}
