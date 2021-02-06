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
)
