package constants

const (
	REQUEST_INVALID                      = "request_invalid"
	FAILED_FETCHING_RESULT_FROM_DATABASE = "select_query_failed"
	FAILED_INSERT_DATABASE               = "insert_query_failed"
	FAILED_CREATING_USER                 = "user_creation_failed"
	REQUEST_DECODING_FAILED              = "decoding_request_failed"
	USER_DO_NOT_EXIST                    = "user_not_present"
	LOGIN_FAILED                         = "login_failed"
	REGION_CREATION_FAILED               = "region_creation_failed"
	HALL_CREATION_FAILED                 = "hall_creation_failed"
	REGION_DO_NOT_EXIST                  = "region_do_not_exist"
	GET_REGION_CALL_FAILED               = "get_region_by_id_failed"
	NOT_VALID_INTEGER                    = "invalid_integer"
	THEATRE_CREATION_FAILED              = "theatre_creation_failed"
	THEATRE_ALREADY_EXIST                = "theatre_already_exist"
	THEATRE_DO_NOT_EXIST                 = "theatre_do_not_exist"
	GET_THEATRES_CALL_FAILED             = "get_theatres_call_failed"
	HALLS_DO_NOT_EXIST                   = "halls_do_not_exist"
	GET_HALLS_BY_THEATRE_ID_FAILED       = "get_halls_by_theatre_id_failed"
)
