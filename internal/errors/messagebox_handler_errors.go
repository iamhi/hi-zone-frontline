package errors

const MESSAGEBOX_NOT_FOUND_ERROR = "messagebox-not-found"
const MESSAGEBOX_ACCESS_NOT_ALLOWED_ERROR = "messagebox-access-not-allowed"
const MESSAGEBOX_BOX_ALREADY_EXISTS_ERROR = "message-box-already-exists"

type MessageboxHandlerError interface {
	GetCode() string
}

type MessageboxNotFoundError struct{}

func (*MessageboxNotFoundError) GetCode() string {
	return MESSAGEBOX_NOT_FOUND_ERROR
}

type MessageboxAccessNotAllowedError struct{}

func (*MessageboxAccessNotAllowedError) GetCode() string {
	return MESSAGEBOX_ACCESS_NOT_ALLOWED_ERROR
}

type MessageboxAlreadyExistsError struct{}

func (*MessageboxAlreadyExistsError) GetCode() string {
	return MESSAGEBOX_BOX_ALREADY_EXISTS_ERROR
}
