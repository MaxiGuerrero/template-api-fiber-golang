package response

// Resoponse when exist an internal error. Is necessary to set a message.
func InternalError(message string) *Response {
	return &Response{Message: message, Code: 500}
}
