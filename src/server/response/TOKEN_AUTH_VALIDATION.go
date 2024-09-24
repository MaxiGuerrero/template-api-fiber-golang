package response

func TokenAuthValidation() *Response {
	return &Response{Message: "Token is correct", Code: 200}
}
