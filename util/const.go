package util

const (
	CollectionUser    = "user"
	MySecretKeyForJWT = "123123123123123"
	CollectionTweet   = "tweet"
	MongoISEError     = "ERROR DURING DB TRANSACTION"
	ISECode           = "500"
	ISEMessage        = "Internal Server Error"
	DuplicateCode     = "409"
	DuplicateCodeMsg  = "UserName already taken. Try Different Name"
	NotFoundMsg       = "User Not Found"
	NotFoundCode      = "404"
	UnAuthorized      = "401"
	UnAuthorizedMsg   = "Login Failed/expired. Please Login Again"
)
