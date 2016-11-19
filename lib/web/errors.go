package web

// APIError is used to return internal errors to http client.
type APIError struct {
	Description string
	Solution    string
}
