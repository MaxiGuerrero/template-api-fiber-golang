package response

// Default response when an user doesn't have permission to use an endpoint.
func Unauthorized() *Response {
	return &Response{Message: "Unauthorized access", Code: 401}
}
