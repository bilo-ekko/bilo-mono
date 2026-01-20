// Package errors contains domain-specific error types.
package errors

import (
	"fmt"
)

// DomainError represents a domain-specific error
type DomainError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Cause   error  `json:"-"`
}

func (e *DomainError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s: %s: %v", e.Domain, e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s: %s", e.Domain, e.Code, e.Message)
}

func (e *DomainError) Unwrap() error {
	return e.Cause
}

// Common error codes
const (
	ErrCodeNotFound       = "NOT_FOUND"
	ErrCodeValidation     = "VALIDATION_ERROR"
	ErrCodeUnauthorized   = "UNAUTHORIZED"
	ErrCodeForbidden      = "FORBIDDEN"
	ErrCodeInternalError  = "INTERNAL_ERROR"
	ErrCodeExternalAPI    = "EXTERNAL_API_ERROR"
)

// NewNotFoundError creates a not found error
func NewNotFoundError(domain, message string) *DomainError {
	return &DomainError{
		Code:    ErrCodeNotFound,
		Message: message,
		Domain:  domain,
	}
}

// NewValidationError creates a validation error
func NewValidationError(domain, message string) *DomainError {
	return &DomainError{
		Code:    ErrCodeValidation,
		Message: message,
		Domain:  domain,
	}
}

// NewUnauthorizedError creates an unauthorized error
func NewUnauthorizedError(domain, message string) *DomainError {
	return &DomainError{
		Code:    ErrCodeUnauthorized,
		Message: message,
		Domain:  domain,
	}
}

// NewForbiddenError creates a forbidden error
func NewForbiddenError(domain, message string) *DomainError {
	return &DomainError{
		Code:    ErrCodeForbidden,
		Message: message,
		Domain:  domain,
	}
}

// NewInternalError creates an internal error
func NewInternalError(domain, message string, cause error) *DomainError {
	return &DomainError{
		Code:    ErrCodeInternalError,
		Message: message,
		Domain:  domain,
		Cause:   cause,
	}
}

// Wrap wraps an error with domain context
func Wrap(domain string, err error, message string) *DomainError {
	return &DomainError{
		Code:    ErrCodeInternalError,
		Message: message,
		Domain:  domain,
		Cause:   err,
	}
}

// IsDomainError checks if an error is a DomainError and assigns it to target
func IsDomainError(err error, target **DomainError) bool {
	if err == nil {
		return false
	}
	if de, ok := err.(*DomainError); ok {
		*target = de
		return true
	}
	return false
}
