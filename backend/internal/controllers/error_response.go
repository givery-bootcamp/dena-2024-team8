package controllers

// ErrorResponse represents an error response
type ErrorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}