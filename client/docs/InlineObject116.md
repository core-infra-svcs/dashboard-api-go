# InlineObject116

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the sensor alert profile. | 
**Schedule** | Pointer to [**NetworksNetworkIdSensorAlertsProfilesSchedule1**](NetworksNetworkIdSensorAlertsProfilesSchedule1.md) |  | [optional] 
**Conditions** | [**[]NetworksNetworkIdSensorAlertsProfilesConditions**](NetworksNetworkIdSensorAlertsProfilesConditions.md) | List of conditions that will cause the profile to send an alert. | 
**Recipients** | Pointer to [**NetworksNetworkIdSensorAlertsProfilesRecipients**](NetworksNetworkIdSensorAlertsProfilesRecipients.md) |  | [optional] 
**Serials** | Pointer to **[]string** | List of device serials assigned to this sensor alert profile. | [optional] 
**IncludeSensorUrl** | Pointer to **bool** | Include dashboard link to sensor in messages (default: true). | [optional] [default to true]
**Message** | Pointer to **string** | A custom message that will appear in email and text message alerts. | [optional] 

## Methods

### NewInlineObject116

`func NewInlineObject116(name string, conditions []NetworksNetworkIdSensorAlertsProfilesConditions, ) *InlineObject116`

NewInlineObject116 instantiates a new InlineObject116 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject116WithDefaults

`func NewInlineObject116WithDefaults() *InlineObject116`

NewInlineObject116WithDefaults instantiates a new InlineObject116 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject116) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject116) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject116) SetName(v string)`

SetName sets Name field to given value.


### GetSchedule

`func (o *InlineObject116) GetSchedule() NetworksNetworkIdSensorAlertsProfilesSchedule1`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *InlineObject116) GetScheduleOk() (*NetworksNetworkIdSensorAlertsProfilesSchedule1, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *InlineObject116) SetSchedule(v NetworksNetworkIdSensorAlertsProfilesSchedule1)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *InlineObject116) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetConditions

`func (o *InlineObject116) GetConditions() []NetworksNetworkIdSensorAlertsProfilesConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InlineObject116) GetConditionsOk() (*[]NetworksNetworkIdSensorAlertsProfilesConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InlineObject116) SetConditions(v []NetworksNetworkIdSensorAlertsProfilesConditions)`

SetConditions sets Conditions field to given value.


### GetRecipients

`func (o *InlineObject116) GetRecipients() NetworksNetworkIdSensorAlertsProfilesRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineObject116) GetRecipientsOk() (*NetworksNetworkIdSensorAlertsProfilesRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineObject116) SetRecipients(v NetworksNetworkIdSensorAlertsProfilesRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InlineObject116) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSerials

`func (o *InlineObject116) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject116) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject116) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineObject116) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetIncludeSensorUrl

`func (o *InlineObject116) GetIncludeSensorUrl() bool`

GetIncludeSensorUrl returns the IncludeSensorUrl field if non-nil, zero value otherwise.

### GetIncludeSensorUrlOk

`func (o *InlineObject116) GetIncludeSensorUrlOk() (*bool, bool)`

GetIncludeSensorUrlOk returns a tuple with the IncludeSensorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSensorUrl

`func (o *InlineObject116) SetIncludeSensorUrl(v bool)`

SetIncludeSensorUrl sets IncludeSensorUrl field to given value.

### HasIncludeSensorUrl

`func (o *InlineObject116) HasIncludeSensorUrl() bool`

HasIncludeSensorUrl returns a boolean if a field has been set.

### GetMessage

`func (o *InlineObject116) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineObject116) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineObject116) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineObject116) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


