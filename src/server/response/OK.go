package response

// Default response when a operation is success.
func OK() *Response {
	return &Response{Message: "Successful operation", Code: 200}
}
