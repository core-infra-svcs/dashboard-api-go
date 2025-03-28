# InlineResponse200117

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the sensor alert profile. | [optional] 
**Name** | Pointer to **string** | Name of the sensor alert profile. | [optional] 
**Schedule** | Pointer to [**NetworksNetworkIdSensorAlertsProfilesSchedule**](NetworksNetworkIdSensorAlertsProfilesSchedule.md) |  | [optional] 
**Conditions** | [**[]NetworksNetworkIdSensorAlertsProfilesConditions**](NetworksNetworkIdSensorAlertsProfilesConditions.md) | List of conditions that will cause the profile to send an alert. | 
**Recipients** | Pointer to [**NetworksNetworkIdSensorAlertsProfilesRecipients**](NetworksNetworkIdSensorAlertsProfilesRecipients.md) |  | [optional] 
**Serials** | Pointer to **[]string** | List of device serials assigned to this sensor alert profile. | [optional] 

## Methods

### NewInlineResponse200117

`func NewInlineResponse200117(conditions []NetworksNetworkIdSensorAlertsProfilesConditions, ) *InlineResponse200117`

NewInlineResponse200117 instantiates a new InlineResponse200117 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200117WithDefaults

`func NewInlineResponse200117WithDefaults() *InlineResponse200117`

NewInlineResponse200117WithDefaults instantiates a new InlineResponse200117 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *InlineResponse200117) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *InlineResponse200117) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *InlineResponse200117) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *InlineResponse200117) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200117) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200117) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200117) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200117) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *InlineResponse200117) GetSchedule() NetworksNetworkIdSensorAlertsProfilesSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *InlineResponse200117) GetScheduleOk() (*NetworksNetworkIdSensorAlertsProfilesSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *InlineResponse200117) SetSchedule(v NetworksNetworkIdSensorAlertsProfilesSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *InlineResponse200117) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetConditions

`func (o *InlineResponse200117) GetConditions() []NetworksNetworkIdSensorAlertsProfilesConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InlineResponse200117) GetConditionsOk() (*[]NetworksNetworkIdSensorAlertsProfilesConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InlineResponse200117) SetConditions(v []NetworksNetworkIdSensorAlertsProfilesConditions)`

SetConditions sets Conditions field to given value.


### GetRecipients

`func (o *InlineResponse200117) GetRecipients() NetworksNetworkIdSensorAlertsProfilesRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineResponse200117) GetRecipientsOk() (*NetworksNetworkIdSensorAlertsProfilesRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineResponse200117) SetRecipients(v NetworksNetworkIdSensorAlertsProfilesRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InlineResponse200117) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSerials

`func (o *InlineResponse200117) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse200117) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse200117) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse200117) HasSerials() bool`

HasSerials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


