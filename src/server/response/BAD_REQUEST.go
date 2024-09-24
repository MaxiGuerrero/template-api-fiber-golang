package response

// Response when a operation is has a bad request. Is necessary to set a message.
func BadRequest(message string) *Response {
	return &Response{Message: message, Code: 400}
}
