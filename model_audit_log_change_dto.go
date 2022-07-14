/*
 * LogCom API
 *
 * LogCom Swagger documentation
 *
 * API version: 1.0.56
 * Contact: laborit@blutspende.de
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package logcomapi

import (
	"github.com/google/uuid"
)

type AuditLogChangeDto struct {
	// The category of the change
	Category string `json:"category,omitempty"`
	// The ID
	Id *uuid.UUID `json:"id,omitempty"`
	// The message
	Message string `json:"message,omitempty"`
	// The new value
	NewValue string `json:"newValue,omitempty"`
	// The old value
	OldValue string `json:"oldValue,omitempty"`
	// The short description of the change
	Subject string `json:"subject,omitempty"`
	// The name of the subject
	SubjectName string `json:"subjectName,omitempty"`
	// The property name of the subject
	SubjectPropertyName string `json:"subjectPropertyName,omitempty"`
}
