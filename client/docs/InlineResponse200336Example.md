# InlineResponse200336Example

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version of the alert | [optional] 
**SharedSecret** | Pointer to **string** | Shared secret for the alert | [optional] 
**SentAt** | Pointer to **time.Time** | When the alert notification was sent, in ISO8601 format | [optional] 
**AlertId** | Pointer to **string** | ID for the alert instance | [optional] 
**AlertLevel** | Pointer to **string** | Severity level of the alert | [optional] 
**OccurredAt** | Pointer to **time.Time** | When the alert occurred, in ISO8601 format | [optional] 
**AlertData** | Pointer to **map[string]interface{}** | Data for the specific alert. Contents depend on the type of the alert | [optional] 
**OrganizationId** | Pointer to **string** | ID of the organization associated with the alert | [optional] 
**OrganizationName** | Pointer to **string** | Name of the organization associated with the alert | [optional] 
**OrganizationUrl** | Pointer to **string** | URL of the organization associated with the alert | [optional] 
**DeviceSerial** | Pointer to **string** | Serial for the device associated with the alert | [optional] 
**DeviceMac** | Pointer to **string** | Mac address for the device associated with the alert | [optional] 
**DeviceName** | Pointer to **string** | Name of the device associated with the alert | [optional] 
**DeviceUrl** | Pointer to **string** | URL of the device associated with the alert | [optional] 
**DeviceTags** | Pointer to **[]string** | List of tags for the device associated with the alert | [optional] 
**DeviceModel** | Pointer to **string** | Model of the device associated with the alert | [optional] 
**NetworkId** | Pointer to **string** | ID of the network associated with the alert | [optional] 
**NetworkName** | Pointer to **string** | Name of the network associated with the alert | [optional] 
**NetworkUrl** | Pointer to **string** | URL of the network associated with the alert | [optional] 
**EnrollmentString** | Pointer to **string** | Enrollment string of the network associated with the alert | [optional] 
**Notes** | Pointer to **string** | Notes for the network associated with the alert | [optional] 
**ProductTypes** | Pointer to **[]string** | List of product types that are part of the network associated with the alert | [optional] 
**EncryptedId** | Pointer to **string** | Encrypted ID of the network associated with the alert | [optional] 

## Methods

### NewInlineResponse200336Example

`func NewInlineResponse200336Example() *InlineResponse200336Example`

NewInlineResponse200336Example instantiates a new InlineResponse200336Example object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200336ExampleWithDefaults

`func NewInlineResponse200336ExampleWithDefaults() *InlineResponse200336Example`

NewInlineResponse200336ExampleWithDefaults instantiates a new InlineResponse200336Example object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *InlineResponse200336Example) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InlineResponse200336Example) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InlineResponse200336Example) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InlineResponse200336Example) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSharedSecret

`func (o *InlineResponse200336Example) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *InlineResponse200336Example) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *InlineResponse200336Example) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *InlineResponse200336Example) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetSentAt

`func (o *InlineResponse200336Example) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *InlineResponse200336Example) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *InlineResponse200336Example) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *InlineResponse200336Example) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### GetAlertId

`func (o *InlineResponse200336Example) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *InlineResponse200336Example) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *InlineResponse200336Example) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *InlineResponse200336Example) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetAlertLevel

`func (o *InlineResponse200336Example) GetAlertLevel() string`

GetAlertLevel returns the AlertLevel field if non-nil, zero value otherwise.

### GetAlertLevelOk

`func (o *InlineResponse200336Example) GetAlertLevelOk() (*string, bool)`

GetAlertLevelOk returns a tuple with the AlertLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLevel

`func (o *InlineResponse200336Example) SetAlertLevel(v string)`

SetAlertLevel sets AlertLevel field to given value.

### HasAlertLevel

`func (o *InlineResponse200336Example) HasAlertLevel() bool`

HasAlertLevel returns a boolean if a field has been set.

### GetOccurredAt

`func (o *InlineResponse200336Example) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *InlineResponse200336Example) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *InlineResponse200336Example) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *InlineResponse200336Example) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetAlertData

`func (o *InlineResponse200336Example) GetAlertData() map[string]interface{}`

GetAlertData returns the AlertData field if non-nil, zero value otherwise.

### GetAlertDataOk

`func (o *InlineResponse200336Example) GetAlertDataOk() (*map[string]interface{}, bool)`

GetAlertDataOk returns a tuple with the AlertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertData

`func (o *InlineResponse200336Example) SetAlertData(v map[string]interface{})`

SetAlertData sets AlertData field to given value.

### HasAlertData

`func (o *InlineResponse200336Example) HasAlertData() bool`

HasAlertData returns a boolean if a field has been set.

### GetOrganizationId

`func (o *InlineResponse200336Example) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InlineResponse200336Example) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InlineResponse200336Example) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *InlineResponse200336Example) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *InlineResponse200336Example) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *InlineResponse200336Example) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *InlineResponse200336Example) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *InlineResponse200336Example) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationUrl

`func (o *InlineResponse200336Example) GetOrganizationUrl() string`

GetOrganizationUrl returns the OrganizationUrl field if non-nil, zero value otherwise.

### GetOrganizationUrlOk

`func (o *InlineResponse200336Example) GetOrganizationUrlOk() (*string, bool)`

GetOrganizationUrlOk returns a tuple with the OrganizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUrl

`func (o *InlineResponse200336Example) SetOrganizationUrl(v string)`

SetOrganizationUrl sets OrganizationUrl field to given value.

### HasOrganizationUrl

`func (o *InlineResponse200336Example) HasOrganizationUrl() bool`

HasOrganizationUrl returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *InlineResponse200336Example) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *InlineResponse200336Example) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *InlineResponse200336Example) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *InlineResponse200336Example) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetDeviceMac

`func (o *InlineResponse200336Example) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *InlineResponse200336Example) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *InlineResponse200336Example) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *InlineResponse200336Example) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDeviceName

`func (o *InlineResponse200336Example) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *InlineResponse200336Example) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *InlineResponse200336Example) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *InlineResponse200336Example) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceUrl

`func (o *InlineResponse200336Example) GetDeviceUrl() string`

GetDeviceUrl returns the DeviceUrl field if non-nil, zero value otherwise.

### GetDeviceUrlOk

`func (o *InlineResponse200336Example) GetDeviceUrlOk() (*string, bool)`

GetDeviceUrlOk returns a tuple with the DeviceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUrl

`func (o *InlineResponse200336Example) SetDeviceUrl(v string)`

SetDeviceUrl sets DeviceUrl field to given value.

### HasDeviceUrl

`func (o *InlineResponse200336Example) HasDeviceUrl() bool`

HasDeviceUrl returns a boolean if a field has been set.

### GetDeviceTags

`func (o *InlineResponse200336Example) GetDeviceTags() []string`

GetDeviceTags returns the DeviceTags field if non-nil, zero value otherwise.

### GetDeviceTagsOk

`func (o *InlineResponse200336Example) GetDeviceTagsOk() (*[]string, bool)`

GetDeviceTagsOk returns a tuple with the DeviceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTags

`func (o *InlineResponse200336Example) SetDeviceTags(v []string)`

SetDeviceTags sets DeviceTags field to given value.

### HasDeviceTags

`func (o *InlineResponse200336Example) HasDeviceTags() bool`

HasDeviceTags returns a boolean if a field has been set.

### GetDeviceModel

`func (o *InlineResponse200336Example) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *InlineResponse200336Example) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *InlineResponse200336Example) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *InlineResponse200336Example) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200336Example) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200336Example) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200336Example) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200336Example) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *InlineResponse200336Example) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse200336Example) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse200336Example) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *InlineResponse200336Example) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNetworkUrl

`func (o *InlineResponse200336Example) GetNetworkUrl() string`

GetNetworkUrl returns the NetworkUrl field if non-nil, zero value otherwise.

### GetNetworkUrlOk

`func (o *InlineResponse200336Example) GetNetworkUrlOk() (*string, bool)`

GetNetworkUrlOk returns a tuple with the NetworkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkUrl

`func (o *InlineResponse200336Example) SetNetworkUrl(v string)`

SetNetworkUrl sets NetworkUrl field to given value.

### HasNetworkUrl

`func (o *InlineResponse200336Example) HasNetworkUrl() bool`

HasNetworkUrl returns a boolean if a field has been set.

### GetEnrollmentString

`func (o *InlineResponse200336Example) GetEnrollmentString() string`

GetEnrollmentString returns the EnrollmentString field if non-nil, zero value otherwise.

### GetEnrollmentStringOk

`func (o *InlineResponse200336Example) GetEnrollmentStringOk() (*string, bool)`

GetEnrollmentStringOk returns a tuple with the EnrollmentString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentString

`func (o *InlineResponse200336Example) SetEnrollmentString(v string)`

SetEnrollmentString sets EnrollmentString field to given value.

### HasEnrollmentString

`func (o *InlineResponse200336Example) HasEnrollmentString() bool`

HasEnrollmentString returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse200336Example) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse200336Example) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse200336Example) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse200336Example) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetProductTypes

`func (o *InlineResponse200336Example) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *InlineResponse200336Example) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *InlineResponse200336Example) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *InlineResponse200336Example) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetEncryptedId

`func (o *InlineResponse200336Example) GetEncryptedId() string`

GetEncryptedId returns the EncryptedId field if non-nil, zero value otherwise.

### GetEncryptedIdOk

`func (o *InlineResponse200336Example) GetEncryptedIdOk() (*string, bool)`

GetEncryptedIdOk returns a tuple with the EncryptedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedId

`func (o *InlineResponse200336Example) SetEncryptedId(v string)`

SetEncryptedId sets EncryptedId field to given value.

### HasEncryptedId

`func (o *InlineResponse200336Example) HasEncryptedId() bool`

HasEncryptedId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


