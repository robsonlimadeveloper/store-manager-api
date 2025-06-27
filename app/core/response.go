package core

// JsonResponse represents a standard JSON response structure.
// It includes a message, optional data, and an optional error field.
// This structure can be used to provide consistent responses across the application.

type JsonResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}