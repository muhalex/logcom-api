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
)

// CreateAuditLogChangeDTO struct for CreateAuditLogChangeDTO
type CreateAuditLogChangeDTO struct {
	// The category of the change
	Category *string `json:"category,omitempty"`
	// The message
	Message *string `json:"message,omitempty"`
	// The new value
	NewValue *string `json:"newValue,omitempty"`
	// The old value
	OldValue *string `json:"oldValue,omitempty"`
	// The short description of the change
	Subject *string `json:"subject,omitempty"`
	// The name of the subject
	SubjectName *string `json:"subjectName,omitempty"`
	// The property name of the subject
	SubjectPropertyName *string `json:"subjectPropertyName,omitempty"`
}

// NewCreateAuditLogChangeDTO instantiates a new CreateAuditLogChangeDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuditLogChangeDTO() *CreateAuditLogChangeDTO {
	this := CreateAuditLogChangeDTO{}
	return &this
}

// NewCreateAuditLogChangeDTOWithDefaults instantiates a new CreateAuditLogChangeDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuditLogChangeDTOWithDefaults() *CreateAuditLogChangeDTO {
	this := CreateAuditLogChangeDTO{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CreateAuditLogChangeDTO) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogChangeDTO) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CreateAuditLogChangeDTO) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CreateAuditLogChangeDTO) SetCategory(v string) {
	o.Category = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateAuditLogChangeDTO) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogChangeDTO) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateAuditLogChangeDTO) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateAuditLogChangeDTO) SetMessage(v string) {
	o.Message = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *CreateAuditLogChangeDTO) GetNewValue() string {
	if o == nil || isNil(o.NewValue) {
		var ret string
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogChangeDTO) GetNewValueOk() (*string, bool) {
	if o == nil || isNil(o.NewValue) {
    return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *CreateAuditLogChangeDTO) HasNewValue() bool {
	if o != nil && !isNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given string and assigns it to the NewValue field.
func (o *CreateAuditLogChangeDTO) SetNewValue(v string) {
	o.NewValue = &v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *CreateAuditLogChangeDTO) GetOldValue() string {
	if o == nil || isNil(o.OldValue) {
		var ret string
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogChangeDTO) GetOldValueOk() (*string, bool) {
	if o == nil || isNil(o.OldValue) {
    return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *CreateAuditLogChangeDTO) HasOldValue() bool {
	if o != nil && !isNil(o.OldValue) {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given string and assigns it to the OldValue field.
func (o *CreateAuditLogChangeDTO) SetOldValue(v string) {
	o.OldValue = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CreateAuditLogChangeDTO) GetSubject() string {
	if o == nil || isNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogChangeDTO) GetSubjectOk() (*string, bool) {
	if o == nil || isNil(o.Subject) {
    return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CreateAuditLogChangeDTO) HasSubject() bool {
	if o != nil && !isNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CreateAuditLogChangeDTO) SetSubject(v string) {
	o.Subject = &v
}

// GetSubjectName returns the SubjectName field value if set, zero value otherwise.
func (o *CreateAuditLogChangeDTO) GetSubjectName() string {
	if o == nil || isNil(o.SubjectName) {
		var ret string
		return ret
	}
	return *o.SubjectName
}

// GetSubjectNameOk returns a tuple with the SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogChangeDTO) GetSubjectNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectName) {
    return nil, false
	}
	return o.SubjectName, true
}

// HasSubjectName returns a boolean if a field has been set.
func (o *CreateAuditLogChangeDTO) HasSubjectName() bool {
	if o != nil && !isNil(o.SubjectName) {
		return true
	}

	return false
}

// SetSubjectName gets a reference to the given string and assigns it to the SubjectName field.
func (o *CreateAuditLogChangeDTO) SetSubjectName(v string) {
	o.SubjectName = &v
}

// GetSubjectPropertyName returns the SubjectPropertyName field value if set, zero value otherwise.
func (o *CreateAuditLogChangeDTO) GetSubjectPropertyName() string {
	if o == nil || isNil(o.SubjectPropertyName) {
		var ret string
		return ret
	}
	return *o.SubjectPropertyName
}

// GetSubjectPropertyNameOk returns a tuple with the SubjectPropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuditLogChangeDTO) GetSubjectPropertyNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectPropertyName) {
    return nil, false
	}
	return o.SubjectPropertyName, true
}

// HasSubjectPropertyName returns a boolean if a field has been set.
func (o *CreateAuditLogChangeDTO) HasSubjectPropertyName() bool {
	if o != nil && !isNil(o.SubjectPropertyName) {
		return true
	}

	return false
}

// SetSubjectPropertyName gets a reference to the given string and assigns it to the SubjectPropertyName field.
func (o *CreateAuditLogChangeDTO) SetSubjectPropertyName(v string) {
	o.SubjectPropertyName = &v
}

func (o CreateAuditLogChangeDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.NewValue) {
		toSerialize["newValue"] = o.NewValue
	}
	if !isNil(o.OldValue) {
		toSerialize["oldValue"] = o.OldValue
	}
	if !isNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !isNil(o.SubjectName) {
		toSerialize["subjectName"] = o.SubjectName
	}
	if !isNil(o.SubjectPropertyName) {
		toSerialize["subjectPropertyName"] = o.SubjectPropertyName
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAuditLogChangeDTO struct {
	value *CreateAuditLogChangeDTO
	isSet bool
}

func (v NullableCreateAuditLogChangeDTO) Get() *CreateAuditLogChangeDTO {
	return v.value
}

func (v *NullableCreateAuditLogChangeDTO) Set(val *CreateAuditLogChangeDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuditLogChangeDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuditLogChangeDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuditLogChangeDTO(val *CreateAuditLogChangeDTO) *NullableCreateAuditLogChangeDTO {
	return &NullableCreateAuditLogChangeDTO{value: val, isSet: true}
}

func (v NullableCreateAuditLogChangeDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuditLogChangeDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


