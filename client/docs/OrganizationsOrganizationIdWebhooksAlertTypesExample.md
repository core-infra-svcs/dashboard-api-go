# OrganizationsOrganizationIdWebhooksAlertTypesExample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version of the alert | [optional] 
**SharedSecret** | Pointer to **string** | Shared secret for the alert | [optional] 
**SentAt** | Pointer to **time.Time** | When the alert notification was sent, in ISO8601 format | [optional] 
**AlertId** | Pointer to **string** | ID for the alert instance | [optional] 
**AlertType** | Pointer to **string** | The type of Meraki alert | [optional] 
**AlertTypeId** | Pointer to **string** | The type ID of Meraki alert | [optional] 
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
**NetworkTags** | Pointer to **[]string** | List of tags for the network associated with the alert | [optional] 
**EnrollmentString** | Pointer to **string** | Enrollment string of the network associated with the alert | [optional] 
**Notes** | Pointer to **string** | Notes for the network associated with the alert | [optional] 
**ProductTypes** | Pointer to **[]string** | List of product types that are part of the network associated with the alert | [optional] 
**EncryptedId** | Pointer to **string** | Encrypted ID of the network associated with the alert | [optional] 

## Methods

### NewOrganizationsOrganizationIdWebhooksAlertTypesExample

`func NewOrganizationsOrganizationIdWebhooksAlertTypesExample() *OrganizationsOrganizationIdWebhooksAlertTypesExample`

NewOrganizationsOrganizationIdWebhooksAlertTypesExample instantiates a new OrganizationsOrganizationIdWebhooksAlertTypesExample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdWebhooksAlertTypesExampleWithDefaults

`func NewOrganizationsOrganizationIdWebhooksAlertTypesExampleWithDefaults() *OrganizationsOrganizationIdWebhooksAlertTypesExample`

NewOrganizationsOrganizationIdWebhooksAlertTypesExampleWithDefaults instantiates a new OrganizationsOrganizationIdWebhooksAlertTypesExample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSharedSecret

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetSentAt

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### GetAlertId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetAlertType

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetAlertTypeId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetAlertTypeId() string`

GetAlertTypeId returns the AlertTypeId field if non-nil, zero value otherwise.

### GetAlertTypeIdOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetAlertTypeIdOk() (*string, bool)`

GetAlertTypeIdOk returns a tuple with the AlertTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypeId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetAlertTypeId(v string)`

SetAlertTypeId sets AlertTypeId field to given value.

### HasAlertTypeId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasAlertTypeId() bool`

HasAlertTypeId returns a boolean if a field has been set.

### GetAlertLevel

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetAlertLevel() string`

GetAlertLevel returns the AlertLevel field if non-nil, zero value otherwise.

### GetAlertLevelOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetAlertLevelOk() (*string, bool)`

GetAlertLevelOk returns a tuple with the AlertLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLevel

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetAlertLevel(v string)`

SetAlertLevel sets AlertLevel field to given value.

### HasAlertLevel

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasAlertLevel() bool`

HasAlertLevel returns a boolean if a field has been set.

### GetOccurredAt

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetAlertData

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetAlertData() map[string]interface{}`

GetAlertData returns the AlertData field if non-nil, zero value otherwise.

### GetAlertDataOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetAlertDataOk() (*map[string]interface{}, bool)`

GetAlertDataOk returns a tuple with the AlertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertData

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetAlertData(v map[string]interface{})`

SetAlertData sets AlertData field to given value.

### HasAlertData

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasAlertData() bool`

HasAlertData returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationUrl

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetOrganizationUrl() string`

GetOrganizationUrl returns the OrganizationUrl field if non-nil, zero value otherwise.

### GetOrganizationUrlOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetOrganizationUrlOk() (*string, bool)`

GetOrganizationUrlOk returns a tuple with the OrganizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUrl

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetOrganizationUrl(v string)`

SetOrganizationUrl sets OrganizationUrl field to given value.

### HasOrganizationUrl

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasOrganizationUrl() bool`

HasOrganizationUrl returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetDeviceMac

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDeviceName

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceUrl

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetDeviceUrl() string`

GetDeviceUrl returns the DeviceUrl field if non-nil, zero value otherwise.

### GetDeviceUrlOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetDeviceUrlOk() (*string, bool)`

GetDeviceUrlOk returns a tuple with the DeviceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUrl

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetDeviceUrl(v string)`

SetDeviceUrl sets DeviceUrl field to given value.

### HasDeviceUrl

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasDeviceUrl() bool`

HasDeviceUrl returns a boolean if a field has been set.

### GetDeviceTags

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetDeviceTags() []string`

GetDeviceTags returns the DeviceTags field if non-nil, zero value otherwise.

### GetDeviceTagsOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetDeviceTagsOk() (*[]string, bool)`

GetDeviceTagsOk returns a tuple with the DeviceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTags

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetDeviceTags(v []string)`

SetDeviceTags sets DeviceTags field to given value.

### HasDeviceTags

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasDeviceTags() bool`

HasDeviceTags returns a boolean if a field has been set.

### GetDeviceModel

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetNetworkId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNetworkUrl

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetNetworkUrl() string`

GetNetworkUrl returns the NetworkUrl field if non-nil, zero value otherwise.

### GetNetworkUrlOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetNetworkUrlOk() (*string, bool)`

GetNetworkUrlOk returns a tuple with the NetworkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkUrl

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetNetworkUrl(v string)`

SetNetworkUrl sets NetworkUrl field to given value.

### HasNetworkUrl

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasNetworkUrl() bool`

HasNetworkUrl returns a boolean if a field has been set.

### GetNetworkTags

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetEnrollmentString

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetEnrollmentString() string`

GetEnrollmentString returns the EnrollmentString field if non-nil, zero value otherwise.

### GetEnrollmentStringOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetEnrollmentStringOk() (*string, bool)`

GetEnrollmentStringOk returns a tuple with the EnrollmentString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentString

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetEnrollmentString(v string)`

SetEnrollmentString sets EnrollmentString field to given value.

### HasEnrollmentString

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasEnrollmentString() bool`

HasEnrollmentString returns a boolean if a field has been set.

### GetNotes

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetProductTypes

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetEncryptedId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetEncryptedId() string`

GetEncryptedId returns the EncryptedId field if non-nil, zero value otherwise.

### GetEncryptedIdOk

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) GetEncryptedIdOk() (*string, bool)`

GetEncryptedIdOk returns a tuple with the EncryptedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) SetEncryptedId(v string)`

SetEncryptedId sets EncryptedId field to given value.

### HasEncryptedId

`func (o *OrganizationsOrganizationIdWebhooksAlertTypesExample) HasEncryptedId() bool`

HasEncryptedId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


