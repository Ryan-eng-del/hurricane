// Package code defines the error codes used in the application.
package code

// HTTP error codes for the IAM system
const (
	// ErrDatabase - 500: Database error.
	ErrDatabase int = 100101

	// ErrUnauthorized - 401: Unauthorized access.
	ErrUnauthorized int = 100102

	// ErrNotFound - 404: Resource not found.
	ErrNotFound int = 100103

	// ErrBadRequest - 400: Bad request.
	ErrBadRequest int = 100104

	// ErrForbidden - 403: Access forbidden.
	ErrForbidden int = 100105

	// ErrInternalServerError - 500: Internal server error.
	ErrInternalServerError int = 100106
)
