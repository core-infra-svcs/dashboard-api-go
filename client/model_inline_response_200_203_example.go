/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200203Example Example alert type
type InlineResponse200203Example struct {
	// Version of the alert
	Version *string `json:"version,omitempty"`
	// Shared secret for the alert
	SharedSecret *string `json:"sharedSecret,omitempty"`
	// When the alert notification was sent, in ISO8601 format
	SentAt *time.Time `json:"sentAt,omitempty"`
	// ID for the alert instance
	AlertId *string `json:"alertId,omitempty"`
	// Severity level of the alert
	AlertLevel *string `json:"alertLevel,omitempty"`
	// When the alert occurred, in ISO8601 format
	OccurredAt *time.Time `json:"occurredAt,omitempty"`
	// Data for the specific alert. Contents depend on the type of the alert
	AlertData map[string]interface{} `json:"alertData,omitempty"`
	// ID of the organization associated with the alert
	OrganizationId *string `json:"organizationId,omitempty"`
	// Name of the organization associated with the alert
	OrganizationName *string `json:"organizationName,omitempty"`
	// URL of the organization associated with the alert
	OrganizationUrl *string `json:"organizationUrl,omitempty"`
	// Serial for the device associated with the alert
	DeviceSerial *string `json:"deviceSerial,omitempty"`
	// Mac address for the device associated with the alert
	DeviceMac *string `json:"deviceMac,omitempty"`
	// Name of the device associated with the alert
	DeviceName *string `json:"deviceName,omitempty"`
	// URL of the device associated with the alert
	DeviceUrl *string `json:"deviceUrl,omitempty"`
	// List of tags for the device associated with the alert
	DeviceTags []string `json:"deviceTags,omitempty"`
	// Model of the device associated with the alert
	DeviceModel *string `json:"deviceModel,omitempty"`
	// ID of the network associated with the alert
	NetworkId *string `json:"networkId,omitempty"`
	// Name of the network associated with the alert
	NetworkName *string `json:"networkName,omitempty"`
	// URL of the network associated with the alert
	NetworkUrl *string `json:"networkUrl,omitempty"`
	// Enrollment string of the network associated with the alert
	EnrollmentString *string `json:"enrollmentString,omitempty"`
	// Notes for the network associated with the alert
	Notes *string `json:"notes,omitempty"`
	// List of product types that are part of the network associated with the alert
	ProductTypes []string `json:"productTypes,omitempty"`
	// Encrypted ID of the network associated with the alert
	EncryptedId *string `json:"encryptedId,omitempty"`
}

// NewInlineResponse200203Example instantiates a new InlineResponse200203Example object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200203Example() *InlineResponse200203Example {
	this := InlineResponse200203Example{}
	return &this
}

// NewInlineResponse200203ExampleWithDefaults instantiates a new InlineResponse200203Example object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200203ExampleWithDefaults() *InlineResponse200203Example {
	this := InlineResponse200203Example{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InlineResponse200203Example) SetVersion(v string) {
	o.Version = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetSharedSecret() string {
	if o == nil || isNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetSharedSecretOk() (*string, bool) {
	if o == nil || isNil(o.SharedSecret) {
    return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasSharedSecret() bool {
	if o != nil && !isNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *InlineResponse200203Example) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetSentAt() time.Time {
	if o == nil || isNil(o.SentAt) {
		var ret time.Time
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetSentAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.SentAt) {
    return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasSentAt() bool {
	if o != nil && !isNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given time.Time and assigns it to the SentAt field.
func (o *InlineResponse200203Example) SetSentAt(v time.Time) {
	o.SentAt = &v
}

// GetAlertId returns the AlertId field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetAlertId() string {
	if o == nil || isNil(o.AlertId) {
		var ret string
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetAlertIdOk() (*string, bool) {
	if o == nil || isNil(o.AlertId) {
    return nil, false
	}
	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasAlertId() bool {
	if o != nil && !isNil(o.AlertId) {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given string and assigns it to the AlertId field.
func (o *InlineResponse200203Example) SetAlertId(v string) {
	o.AlertId = &v
}

// GetAlertLevel returns the AlertLevel field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetAlertLevel() string {
	if o == nil || isNil(o.AlertLevel) {
		var ret string
		return ret
	}
	return *o.AlertLevel
}

// GetAlertLevelOk returns a tuple with the AlertLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetAlertLevelOk() (*string, bool) {
	if o == nil || isNil(o.AlertLevel) {
    return nil, false
	}
	return o.AlertLevel, true
}

// HasAlertLevel returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasAlertLevel() bool {
	if o != nil && !isNil(o.AlertLevel) {
		return true
	}

	return false
}

// SetAlertLevel gets a reference to the given string and assigns it to the AlertLevel field.
func (o *InlineResponse200203Example) SetAlertLevel(v string) {
	o.AlertLevel = &v
}

// GetOccurredAt returns the OccurredAt field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetOccurredAt() time.Time {
	if o == nil || isNil(o.OccurredAt) {
		var ret time.Time
		return ret
	}
	return *o.OccurredAt
}

// GetOccurredAtOk returns a tuple with the OccurredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetOccurredAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.OccurredAt) {
    return nil, false
	}
	return o.OccurredAt, true
}

// HasOccurredAt returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasOccurredAt() bool {
	if o != nil && !isNil(o.OccurredAt) {
		return true
	}

	return false
}

// SetOccurredAt gets a reference to the given time.Time and assigns it to the OccurredAt field.
func (o *InlineResponse200203Example) SetOccurredAt(v time.Time) {
	o.OccurredAt = &v
}

// GetAlertData returns the AlertData field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetAlertData() map[string]interface{} {
	if o == nil || isNil(o.AlertData) {
		var ret map[string]interface{}
		return ret
	}
	return o.AlertData
}

// GetAlertDataOk returns a tuple with the AlertData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetAlertDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.AlertData) {
    return map[string]interface{}{}, false
	}
	return o.AlertData, true
}

// HasAlertData returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasAlertData() bool {
	if o != nil && !isNil(o.AlertData) {
		return true
	}

	return false
}

// SetAlertData gets a reference to the given map[string]interface{} and assigns it to the AlertData field.
func (o *InlineResponse200203Example) SetAlertData(v map[string]interface{}) {
	o.AlertData = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *InlineResponse200203Example) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetOrganizationName() string {
	if o == nil || isNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetOrganizationNameOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationName) {
    return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasOrganizationName() bool {
	if o != nil && !isNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *InlineResponse200203Example) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetOrganizationUrl returns the OrganizationUrl field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetOrganizationUrl() string {
	if o == nil || isNil(o.OrganizationUrl) {
		var ret string
		return ret
	}
	return *o.OrganizationUrl
}

// GetOrganizationUrlOk returns a tuple with the OrganizationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetOrganizationUrlOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationUrl) {
    return nil, false
	}
	return o.OrganizationUrl, true
}

// HasOrganizationUrl returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasOrganizationUrl() bool {
	if o != nil && !isNil(o.OrganizationUrl) {
		return true
	}

	return false
}

// SetOrganizationUrl gets a reference to the given string and assigns it to the OrganizationUrl field.
func (o *InlineResponse200203Example) SetOrganizationUrl(v string) {
	o.OrganizationUrl = &v
}

// GetDeviceSerial returns the DeviceSerial field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetDeviceSerial() string {
	if o == nil || isNil(o.DeviceSerial) {
		var ret string
		return ret
	}
	return *o.DeviceSerial
}

// GetDeviceSerialOk returns a tuple with the DeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetDeviceSerialOk() (*string, bool) {
	if o == nil || isNil(o.DeviceSerial) {
    return nil, false
	}
	return o.DeviceSerial, true
}

// HasDeviceSerial returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasDeviceSerial() bool {
	if o != nil && !isNil(o.DeviceSerial) {
		return true
	}

	return false
}

// SetDeviceSerial gets a reference to the given string and assigns it to the DeviceSerial field.
func (o *InlineResponse200203Example) SetDeviceSerial(v string) {
	o.DeviceSerial = &v
}

// GetDeviceMac returns the DeviceMac field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetDeviceMac() string {
	if o == nil || isNil(o.DeviceMac) {
		var ret string
		return ret
	}
	return *o.DeviceMac
}

// GetDeviceMacOk returns a tuple with the DeviceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetDeviceMacOk() (*string, bool) {
	if o == nil || isNil(o.DeviceMac) {
    return nil, false
	}
	return o.DeviceMac, true
}

// HasDeviceMac returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasDeviceMac() bool {
	if o != nil && !isNil(o.DeviceMac) {
		return true
	}

	return false
}

// SetDeviceMac gets a reference to the given string and assigns it to the DeviceMac field.
func (o *InlineResponse200203Example) SetDeviceMac(v string) {
	o.DeviceMac = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetDeviceName() string {
	if o == nil || isNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetDeviceNameOk() (*string, bool) {
	if o == nil || isNil(o.DeviceName) {
    return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasDeviceName() bool {
	if o != nil && !isNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *InlineResponse200203Example) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceUrl returns the DeviceUrl field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetDeviceUrl() string {
	if o == nil || isNil(o.DeviceUrl) {
		var ret string
		return ret
	}
	return *o.DeviceUrl
}

// GetDeviceUrlOk returns a tuple with the DeviceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetDeviceUrlOk() (*string, bool) {
	if o == nil || isNil(o.DeviceUrl) {
    return nil, false
	}
	return o.DeviceUrl, true
}

// HasDeviceUrl returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasDeviceUrl() bool {
	if o != nil && !isNil(o.DeviceUrl) {
		return true
	}

	return false
}

// SetDeviceUrl gets a reference to the given string and assigns it to the DeviceUrl field.
func (o *InlineResponse200203Example) SetDeviceUrl(v string) {
	o.DeviceUrl = &v
}

// GetDeviceTags returns the DeviceTags field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetDeviceTags() []string {
	if o == nil || isNil(o.DeviceTags) {
		var ret []string
		return ret
	}
	return o.DeviceTags
}

// GetDeviceTagsOk returns a tuple with the DeviceTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetDeviceTagsOk() ([]string, bool) {
	if o == nil || isNil(o.DeviceTags) {
    return nil, false
	}
	return o.DeviceTags, true
}

// HasDeviceTags returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasDeviceTags() bool {
	if o != nil && !isNil(o.DeviceTags) {
		return true
	}

	return false
}

// SetDeviceTags gets a reference to the given []string and assigns it to the DeviceTags field.
func (o *InlineResponse200203Example) SetDeviceTags(v []string) {
	o.DeviceTags = v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetDeviceModel() string {
	if o == nil || isNil(o.DeviceModel) {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetDeviceModelOk() (*string, bool) {
	if o == nil || isNil(o.DeviceModel) {
    return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasDeviceModel() bool {
	if o != nil && !isNil(o.DeviceModel) {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *InlineResponse200203Example) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200203Example) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetNetworkName returns the NetworkName field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetNetworkName() string {
	if o == nil || isNil(o.NetworkName) {
		var ret string
		return ret
	}
	return *o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetNetworkNameOk() (*string, bool) {
	if o == nil || isNil(o.NetworkName) {
    return nil, false
	}
	return o.NetworkName, true
}

// HasNetworkName returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasNetworkName() bool {
	if o != nil && !isNil(o.NetworkName) {
		return true
	}

	return false
}

// SetNetworkName gets a reference to the given string and assigns it to the NetworkName field.
func (o *InlineResponse200203Example) SetNetworkName(v string) {
	o.NetworkName = &v
}

// GetNetworkUrl returns the NetworkUrl field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetNetworkUrl() string {
	if o == nil || isNil(o.NetworkUrl) {
		var ret string
		return ret
	}
	return *o.NetworkUrl
}

// GetNetworkUrlOk returns a tuple with the NetworkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetNetworkUrlOk() (*string, bool) {
	if o == nil || isNil(o.NetworkUrl) {
    return nil, false
	}
	return o.NetworkUrl, true
}

// HasNetworkUrl returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasNetworkUrl() bool {
	if o != nil && !isNil(o.NetworkUrl) {
		return true
	}

	return false
}

// SetNetworkUrl gets a reference to the given string and assigns it to the NetworkUrl field.
func (o *InlineResponse200203Example) SetNetworkUrl(v string) {
	o.NetworkUrl = &v
}

// GetEnrollmentString returns the EnrollmentString field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetEnrollmentString() string {
	if o == nil || isNil(o.EnrollmentString) {
		var ret string
		return ret
	}
	return *o.EnrollmentString
}

// GetEnrollmentStringOk returns a tuple with the EnrollmentString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetEnrollmentStringOk() (*string, bool) {
	if o == nil || isNil(o.EnrollmentString) {
    return nil, false
	}
	return o.EnrollmentString, true
}

// HasEnrollmentString returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasEnrollmentString() bool {
	if o != nil && !isNil(o.EnrollmentString) {
		return true
	}

	return false
}

// SetEnrollmentString gets a reference to the given string and assigns it to the EnrollmentString field.
func (o *InlineResponse200203Example) SetEnrollmentString(v string) {
	o.EnrollmentString = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineResponse200203Example) SetNotes(v string) {
	o.Notes = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetProductTypes() []string {
	if o == nil || isNil(o.ProductTypes) {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetProductTypesOk() ([]string, bool) {
	if o == nil || isNil(o.ProductTypes) {
    return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasProductTypes() bool {
	if o != nil && !isNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *InlineResponse200203Example) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetEncryptedId returns the EncryptedId field value if set, zero value otherwise.
func (o *InlineResponse200203Example) GetEncryptedId() string {
	if o == nil || isNil(o.EncryptedId) {
		var ret string
		return ret
	}
	return *o.EncryptedId
}

// GetEncryptedIdOk returns a tuple with the EncryptedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200203Example) GetEncryptedIdOk() (*string, bool) {
	if o == nil || isNil(o.EncryptedId) {
    return nil, false
	}
	return o.EncryptedId, true
}

// HasEncryptedId returns a boolean if a field has been set.
func (o *InlineResponse200203Example) HasEncryptedId() bool {
	if o != nil && !isNil(o.EncryptedId) {
		return true
	}

	return false
}

// SetEncryptedId gets a reference to the given string and assigns it to the EncryptedId field.
func (o *InlineResponse200203Example) SetEncryptedId(v string) {
	o.EncryptedId = &v
}

func (o InlineResponse200203Example) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !isNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	if !isNil(o.AlertId) {
		toSerialize["alertId"] = o.AlertId
	}
	if !isNil(o.AlertLevel) {
		toSerialize["alertLevel"] = o.AlertLevel
	}
	if !isNil(o.OccurredAt) {
		toSerialize["occurredAt"] = o.OccurredAt
	}
	if !isNil(o.AlertData) {
		toSerialize["alertData"] = o.AlertData
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !isNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !isNil(o.OrganizationUrl) {
		toSerialize["organizationUrl"] = o.OrganizationUrl
	}
	if !isNil(o.DeviceSerial) {
		toSerialize["deviceSerial"] = o.DeviceSerial
	}
	if !isNil(o.DeviceMac) {
		toSerialize["deviceMac"] = o.DeviceMac
	}
	if !isNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !isNil(o.DeviceUrl) {
		toSerialize["deviceUrl"] = o.DeviceUrl
	}
	if !isNil(o.DeviceTags) {
		toSerialize["deviceTags"] = o.DeviceTags
	}
	if !isNil(o.DeviceModel) {
		toSerialize["deviceModel"] = o.DeviceModel
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.NetworkName) {
		toSerialize["networkName"] = o.NetworkName
	}
	if !isNil(o.NetworkUrl) {
		toSerialize["networkUrl"] = o.NetworkUrl
	}
	if !isNil(o.EnrollmentString) {
		toSerialize["enrollmentString"] = o.EnrollmentString
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.ProductTypes) {
		toSerialize["productTypes"] = o.ProductTypes
	}
	if !isNil(o.EncryptedId) {
		toSerialize["encryptedId"] = o.EncryptedId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200203Example struct {
	value *InlineResponse200203Example
	isSet bool
}

func (v NullableInlineResponse200203Example) Get() *InlineResponse200203Example {
	return v.value
}

func (v *NullableInlineResponse200203Example) Set(val *InlineResponse200203Example) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200203Example) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200203Example) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200203Example(val *InlineResponse200203Example) *NullableInlineResponse200203Example {
	return &NullableInlineResponse200203Example{value: val, isSet: true}
}

func (v NullableInlineResponse200203Example) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200203Example) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

