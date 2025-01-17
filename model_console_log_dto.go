/*
LogCom API

LogCom Swagger documentation

API version: 1.2.21
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

// ConsoleLogDTO struct for ConsoleLogDTO
type ConsoleLogDTO struct {
	// The log timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user's ID who created
	CreatedById *uuid.UUID `json:"createdById,omitempty"`
	// The user's name who created
	CreatedByName *string `json:"createdByName,omitempty"`
	// The ID
	Id *uuid.UUID `json:"id,omitempty"`
	// The log level (Trace=-1, Debug=0, Info=1, Warning=2, Error=3, Fatal=4, Panic=5)
	Level *LogLevel `json:"level,omitempty"`
	// The log message
	Message *string `json:"message,omitempty"`
	// The request ID making dependent logs trackable
	RequestId *uuid.UUID `json:"requestId,omitempty"`
	// The service which sent the log
	Service *string `json:"service,omitempty"`
}

// NewConsoleLogDTO instantiates a new ConsoleLogDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsoleLogDTO() *ConsoleLogDTO {
	this := ConsoleLogDTO{}
	return &this
}

// NewConsoleLogDTOWithDefaults instantiates a new ConsoleLogDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsoleLogDTOWithDefaults() *ConsoleLogDTO {
	this := ConsoleLogDTO{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ConsoleLogDTO) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogDTO) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ConsoleLogDTO) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ConsoleLogDTO) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *ConsoleLogDTO) GetCreatedById() uuid.UUID {
	if o == nil || isNil(o.CreatedById) {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogDTO) GetCreatedByIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.CreatedById) {
    return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *ConsoleLogDTO) HasCreatedById() bool {
	if o != nil && !isNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given uuid.UUID and assigns it to the CreatedById field.
func (o *ConsoleLogDTO) SetCreatedById(v uuid.UUID) {
	o.CreatedById = &v
}

// GetCreatedByName returns the CreatedByName field value if set, zero value otherwise.
func (o *ConsoleLogDTO) GetCreatedByName() string {
	if o == nil || isNil(o.CreatedByName) {
		var ret string
		return ret
	}
	return *o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogDTO) GetCreatedByNameOk() (*string, bool) {
	if o == nil || isNil(o.CreatedByName) {
    return nil, false
	}
	return o.CreatedByName, true
}

// HasCreatedByName returns a boolean if a field has been set.
func (o *ConsoleLogDTO) HasCreatedByName() bool {
	if o != nil && !isNil(o.CreatedByName) {
		return true
	}

	return false
}

// SetCreatedByName gets a reference to the given string and assigns it to the CreatedByName field.
func (o *ConsoleLogDTO) SetCreatedByName(v string) {
	o.CreatedByName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConsoleLogDTO) GetId() uuid.UUID {
	if o == nil || isNil(o.Id) {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogDTO) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConsoleLogDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *ConsoleLogDTO) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *ConsoleLogDTO) GetLevel() LogLevel {
	if o == nil || isNil(o.Level) {
		var ret LogLevel
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogDTO) GetLevelOk() (*LogLevel, bool) {
	if o == nil || isNil(o.Level) {
    return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ConsoleLogDTO) HasLevel() bool {
	if o != nil && !isNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given LogLevel and assigns it to the Level field.
func (o *ConsoleLogDTO) SetLevel(v LogLevel) {
	o.Level = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ConsoleLogDTO) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogDTO) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ConsoleLogDTO) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ConsoleLogDTO) SetMessage(v string) {
	o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ConsoleLogDTO) GetRequestId() uuid.UUID {
	if o == nil || isNil(o.RequestId) {
		var ret uuid.UUID
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogDTO) GetRequestIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.RequestId) {
    return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ConsoleLogDTO) HasRequestId() bool {
	if o != nil && !isNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given uuid.UUID and assigns it to the RequestId field.
func (o *ConsoleLogDTO) SetRequestId(v uuid.UUID) {
	o.RequestId = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ConsoleLogDTO) GetService() string {
	if o == nil || isNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogDTO) GetServiceOk() (*string, bool) {
	if o == nil || isNil(o.Service) {
    return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ConsoleLogDTO) HasService() bool {
	if o != nil && !isNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ConsoleLogDTO) SetService(v string) {
	o.Service = &v
}

func (o ConsoleLogDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if !isNil(o.CreatedByName) {
		toSerialize["createdByName"] = o.CreatedByName
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !isNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	return json.Marshal(toSerialize)
}

type NullableConsoleLogDTO struct {
	value *ConsoleLogDTO
	isSet bool
}

func (v NullableConsoleLogDTO) Get() *ConsoleLogDTO {
	return v.value
}

func (v *NullableConsoleLogDTO) Set(val *ConsoleLogDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableConsoleLogDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableConsoleLogDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsoleLogDTO(val *ConsoleLogDTO) *NullableConsoleLogDTO {
	return &NullableConsoleLogDTO{value: val, isSet: true}
}

func (v NullableConsoleLogDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsoleLogDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


