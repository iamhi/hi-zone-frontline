package errors

const USER_EXISTS_CODE = "user-exists"
const USER_NOT_FOUND_CODE = "user-not-found"
const USER_WRONG_PASSWORD_CODE = "user-wrong-password"
const USER_INVALID_TOKEN_CODE = "user-invalid-token"

type UserHandlerError interface {
	GetCode() string
}

type UserExistsError struct {
	Username string
	Email    string
}

func (e UserExistsError) Error() string {
	return "User with username or email exists"
}

func (e UserExistsError) GetCode() string {
	return USER_EXISTS_CODE
}

type UserNotFoundError struct {
	Username string
}

func (e UserNotFoundError) Error() string {
	return "User with username or email exists"
}

func (e UserNotFoundError) GetCode() string {
	return USER_NOT_FOUND_CODE
}

type UserBadCredentialsError struct {
	Username string
}

func (e UserBadCredentialsError) Error() string {
	return "Incorrect username or password provided"
}

func (e UserBadCredentialsError) GetCode() string {
	return USER_WRONG_PASSWORD_CODE
}

type UserInvalidTokenError struct {
	Token string
}

func (e UserInvalidTokenError) Error() string {
	return "Invalid token provided"
}

func (e UserInvalidTokenError) GetCode() string {
	return USER_INVALID_TOKEN_CODE
}
