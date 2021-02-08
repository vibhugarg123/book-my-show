package constants

const (
	FAILED_STARTING_APPLICATION = "failed to start the application"
	FAILED_MYSQL_CONNECTION     = "failed to connect with my sql connection"
	EMPTY_REQUEST_BODY          = "request body is empty"
	EMAIL_ID_MANDATORY          = "email id is missing in request"
	PASSWORD_MANDATORY          = "password is missing in request"
	FIRST_NAME_MANDATORY        = "first name is missing in request"
	FAILED_GET_DB_CALL          = "failed fetching result from database"
	USER_ALREADY_EXISTS         = "user with email id- %s already exists"
	USER_CREATION_FAILED        = "failed to create user"
	DECODING_REQUEST_FAILED     = "failed to decode the request body"
	USER_DOES_NOT_EXIST         = "user does not exist"
	REGION_ID_MANDATORY         = "region id is missing in request"
	REGION_NAME_MANDATORY       = "region name is missing in request"
	REGION_ALREADY_EXISTS       = "region with region id- %d already exists"
	ADD_REGION_FAILED           = "failed to add new region"
	REGION_TYPE_MANDATORY       = "region type is missing in request"
	REGION_DOES_NOT_EXIST       = "region with region id- %d do not exist"
)
