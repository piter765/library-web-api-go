package service_errors

const (
	//Token
	UnexpectedError = "Unexpected error"
	ClaimsNotFound  = "Claims not found"
	TokenRequired   = "token required"
	TokenExpired    = "token expired"
	TokenInvalid    = "token invalid"

	// User
	EmailExists    = "Email exists"
	UsernameExists = "Username exists"
	PermissionDenied = "Permission denied"

	// DB
	RecordNotFound = "Record not found"
)