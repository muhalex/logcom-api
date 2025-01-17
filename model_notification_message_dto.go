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

// NotificationMessageDTO struct for NotificationMessageDTO
type NotificationMessageDTO struct {
	// The notification timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user's ID who created
	CreatedById *uuid.UUID `json:"createdById,omitempty"`
	// The user's name who created
	CreatedByName *string `json:"createdByName,omitempty"`
	// The ID
	Id *uuid.UUID `json:"id,omitempty"`
	// The notification message
	Message *string `json:"message,omitempty"`
	// The service which sent the notification
	Service *string `json:"service,omitempty"`
	// The status of the message
	Status *NotificationStatus `json:"status,omitempty"`
}

// NewNotificationMessageDTO instantiates a new NotificationMessageDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationMessageDTO() *NotificationMessageDTO {
	this := NotificationMessageDTO{}
	return &this
}

// NewNotificationMessageDTOWithDefaults instantiates a new NotificationMessageDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationMessageDTOWithDefaults() *NotificationMessageDTO {
	this := NotificationMessageDTO{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NotificationMessageDTO) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageDTO) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationMessageDTO) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NotificationMessageDTO) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *NotificationMessageDTO) GetCreatedById() uuid.UUID {
	if o == nil || isNil(o.CreatedById) {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageDTO) GetCreatedByIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.CreatedById) {
    return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *NotificationMessageDTO) HasCreatedById() bool {
	if o != nil && !isNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given uuid.UUID and assigns it to the CreatedById field.
func (o *NotificationMessageDTO) SetCreatedById(v uuid.UUID) {
	o.CreatedById = &v
}

// GetCreatedByName returns the CreatedByName field value if set, zero value otherwise.
func (o *NotificationMessageDTO) GetCreatedByName() string {
	if o == nil || isNil(o.CreatedByName) {
		var ret string
		return ret
	}
	return *o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageDTO) GetCreatedByNameOk() (*string, bool) {
	if o == nil || isNil(o.CreatedByName) {
    return nil, false
	}
	return o.CreatedByName, true
}

// HasCreatedByName returns a boolean if a field has been set.
func (o *NotificationMessageDTO) HasCreatedByName() bool {
	if o != nil && !isNil(o.CreatedByName) {
		return true
	}

	return false
}

// SetCreatedByName gets a reference to the given string and assigns it to the CreatedByName field.
func (o *NotificationMessageDTO) SetCreatedByName(v string) {
	o.CreatedByName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationMessageDTO) GetId() uuid.UUID {
	if o == nil || isNil(o.Id) {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageDTO) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationMessageDTO) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *NotificationMessageDTO) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NotificationMessageDTO) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageDTO) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NotificationMessageDTO) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NotificationMessageDTO) SetMessage(v string) {
	o.Message = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *NotificationMessageDTO) GetService() string {
	if o == nil || isNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageDTO) GetServiceOk() (*string, bool) {
	if o == nil || isNil(o.Service) {
    return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *NotificationMessageDTO) HasService() bool {
	if o != nil && !isNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *NotificationMessageDTO) SetService(v string) {
	o.Service = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NotificationMessageDTO) GetStatus() NotificationStatus {
	if o == nil || isNil(o.Status) {
		var ret NotificationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageDTO) GetStatusOk() (*NotificationStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NotificationMessageDTO) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NotificationStatus and assigns it to the Status field.
func (o *NotificationMessageDTO) SetStatus(v NotificationStatus) {
	o.Status = &v
}

func (o NotificationMessageDTO) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationMessageDTO struct {
	value *NotificationMessageDTO
	isSet bool
}

func (v NullableNotificationMessageDTO) Get() *NotificationMessageDTO {
	return v.value
}

func (v *NullableNotificationMessageDTO) Set(val *NotificationMessageDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationMessageDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationMessageDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationMessageDTO(val *NotificationMessageDTO) *NullableNotificationMessageDTO {
	return &NullableNotificationMessageDTO{value: val, isSet: true}
}

func (v NullableNotificationMessageDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationMessageDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


