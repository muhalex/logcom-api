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
	"time"
)

type CreateConsoleLogRequestDto struct {
	// The log timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user's ID who created
	CreatedById *uuid.UUID `json:"createdById,omitempty"`
	// The user's name who created
	CreatedByName string `json:"createdByName,omitempty"`
	// The log level (Trace=-1, Debug=0, Info=1, Warning=2, Error=3, Fatal=4, Panic=5)
	Level int32 `json:"level"`
	// The log message
	Message string `json:"message"`
	// The service which sent the log
	Service string `json:"service"`
}
