package errors

//This package is for making custom errors. In this case, it's used to create a custom error at
//any place of our code, and be able to add anywhere attributes like (in this case, for our needs):
//	-Status code of the error
//	-String with reason of the error
//	-Some description of the error, maybe an error code... (not implemented here)
//This way, when the error reaches the "surface" (handler), you already have this information,
//and can simply return the error as it is.

type ApiError struct {
	originalError error
	message       string `default:"-"` //struct tag, specifies a custom zero-value for this struct field
	statusCode    int
}

func NewFromError(err error) *ApiError {
	return &ApiError{
		originalError: err,
	}
}

func NewFromMessage(message string) *ApiError {
	return &ApiError{
		message: message,
	}
}

// GetStatusCode: a getter. In other packages, you could make use of the struct, but not have direct access to the fields
// (they are initial lowercase, so unexported)
func (e *ApiError) GetStatusCode() int {
	return e.statusCode
}

// GetOriginalMessage: obtain the source error message for debugging and logging purposes.
// It's not the message given by us (struct field "message")
func (e *ApiError) GetOriginalMessage() string {
	if e.originalError != nil {
		return e.originalError.Error()
	} else {
		return ""
	}
}

// WithMessage: set the message of the error
func (e *ApiError) WithMessage(message string) *ApiError {
	e.message = message
	return e
}

// WithStatusCode: set the http code of the error
func (e *ApiError) WithStatusCode(code int) *ApiError {
	e.statusCode = code
	return e
}

// ToJSON: return a json serializable object from the ApiError
func (e *ApiError) ToJSON() map[string]interface{} {
	//Note: maps of type map[string]interface{} can be sent as JSON straight up.
	errorMap := map[string]interface{}{
		"message": e.message,
	}

	return errorMap
}

// Error: method needed to satisfy the base golang "error" interface.
// It generally prints out a description of the error.
func (e *ApiError) Error() string {
	if e.message == "-" {
		//message was not set
		return e.originalError.Error()
	} else {
		return e.message
	}
}
