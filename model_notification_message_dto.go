/*
 * LogCom API
 *
 * LogCom Swagger documentation
 *
 * API version: 1.0.65
 * Contact: laborit@blutspende.de
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package logcomapi

import (
	"github.com/google/uuid"
	"time"
)

type NotificationMessageDto struct {
	// The notification timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user's ID who created
	CreatedById *uuid.UUID `json:"createdById,omitempty"`
	// The user's name who created
	CreatedByName string `json:"createdByName,omitempty"`
	// The ID
	Id *uuid.UUID `json:"id,omitempty"`
	// The notification message
	Message string `json:"message,omitempty"`
	// The service which sent the notification
	Service string `json:"service,omitempty"`
	// The status of the message
	Status string `json:"status,omitempty"`
}
