// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"twitterMock/api/generated/restapi/operations"
	"twitterMock/api/generated/restapi/operations/operation"
	"twitterMock/api/generated/restapi/operations/user"
)

//go:generate swagger generate server --target ../../generated --name TwitterMock --spec ../../swagger/server.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.TwitterMockAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TwitterMockAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.BearerAuth == nil {
		api.BearerAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.OperationFollowHandler == nil {
		api.OperationFollowHandler = operation.FollowHandlerFunc(func(params operation.FollowParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operation.Follow has not yet been implemented")
		})
	}
	if api.OperationGetFeedHandler == nil {
		api.OperationGetFeedHandler = operation.GetFeedHandlerFunc(func(params operation.GetFeedParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operation.GetFeed has not yet been implemented")
		})
	}
	if api.UserLoginUserHandler == nil {
		api.UserLoginUserHandler = user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.LoginUser has not yet been implemented")
		})
	}
	if api.UserLogoutUserHandler == nil {
		api.UserLogoutUserHandler = user.LogoutUserHandlerFunc(func(params user.LogoutUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.LogoutUser has not yet been implemented")
		})
	}
	if api.OperationPostTweetHandler == nil {
		api.OperationPostTweetHandler = operation.PostTweetHandlerFunc(func(params operation.PostTweetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operation.PostTweet has not yet been implemented")
		})
	}
	if api.UserSignupHandler == nil {
		api.UserSignupHandler = user.SignupHandlerFunc(func(params user.SignupParams) middleware.Responder {
			return middleware.NotImplemented("operation user.Signup has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
