package response

// Define how to send response messages in a request, this struct is use along of the system.
type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
