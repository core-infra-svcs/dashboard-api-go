# InlineObject139

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the sensor alert profile. | [optional] 
**Schedule** | Pointer to [**NetworksNetworkIdSensorAlertsProfilesSchedule1**](NetworksNetworkIdSensorAlertsProfilesSchedule1.md) |  | [optional] 
**Conditions** | Pointer to [**[]NetworksNetworkIdSensorAlertsProfilesConditions**](NetworksNetworkIdSensorAlertsProfilesConditions.md) | List of conditions that will cause the profile to send an alert. | [optional] 
**Recipients** | Pointer to [**NetworksNetworkIdSensorAlertsProfilesRecipients**](NetworksNetworkIdSensorAlertsProfilesRecipients.md) |  | [optional] 
**Serials** | Pointer to **[]string** | List of device serials assigned to this sensor alert profile. | [optional] 
**IncludeSensorUrl** | Pointer to **bool** | Include dashboard link to sensor in messages (default: true). | [optional] 
**Message** | Pointer to **NullableString** | A custom message that will appear in email and text message alerts. | [optional] 

## Methods

### NewInlineObject139

`func NewInlineObject139() *InlineObject139`

NewInlineObject139 instantiates a new InlineObject139 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject139WithDefaults

`func NewInlineObject139WithDefaults() *InlineObject139`

NewInlineObject139WithDefaults instantiates a new InlineObject139 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject139) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject139) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject139) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject139) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *InlineObject139) GetSchedule() NetworksNetworkIdSensorAlertsProfilesSchedule1`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *InlineObject139) GetScheduleOk() (*NetworksNetworkIdSensorAlertsProfilesSchedule1, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *InlineObject139) SetSchedule(v NetworksNetworkIdSensorAlertsProfilesSchedule1)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *InlineObject139) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetConditions

`func (o *InlineObject139) GetConditions() []NetworksNetworkIdSensorAlertsProfilesConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InlineObject139) GetConditionsOk() (*[]NetworksNetworkIdSensorAlertsProfilesConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InlineObject139) SetConditions(v []NetworksNetworkIdSensorAlertsProfilesConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *InlineObject139) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetRecipients

`func (o *InlineObject139) GetRecipients() NetworksNetworkIdSensorAlertsProfilesRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineObject139) GetRecipientsOk() (*NetworksNetworkIdSensorAlertsProfilesRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineObject139) SetRecipients(v NetworksNetworkIdSensorAlertsProfilesRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InlineObject139) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSerials

`func (o *InlineObject139) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject139) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject139) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineObject139) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetIncludeSensorUrl

`func (o *InlineObject139) GetIncludeSensorUrl() bool`

GetIncludeSensorUrl returns the IncludeSensorUrl field if non-nil, zero value otherwise.

### GetIncludeSensorUrlOk

`func (o *InlineObject139) GetIncludeSensorUrlOk() (*bool, bool)`

GetIncludeSensorUrlOk returns a tuple with the IncludeSensorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSensorUrl

`func (o *InlineObject139) SetIncludeSensorUrl(v bool)`

SetIncludeSensorUrl sets IncludeSensorUrl field to given value.

### HasIncludeSensorUrl

`func (o *InlineObject139) HasIncludeSensorUrl() bool`

HasIncludeSensorUrl returns a boolean if a field has been set.

### GetMessage

`func (o *InlineObject139) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineObject139) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineObject139) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineObject139) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *InlineObject139) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *InlineObject139) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


