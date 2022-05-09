/*
 * LogCom API
 *
 * LogCom Swagger documentation
 *
 * API version: 1.0.22
 * Contact: laborit@blutspende.de
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package logcomapi

type GroupedAuditLogChangesDto struct {
	// The category of the change
	Category string `json:"category,omitempty"`
	// The message
	Message string `json:"message,omitempty"`
	// The new value
	NewValue string `json:"newValue,omitempty"`
	// The old value
	OldValue string `json:"oldValue,omitempty"`
	// The short description of the change
	Subject string `json:"subject,omitempty"`
}