# InlineResponse200321Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Alert Type | 
**CategoryType** | **string** | Alert category | 
**Severity** | **string** | Alert severity | 
**LastAlertedAt** | **time.Time** | Last time an alert of this type was triggered | 
**LastResolvedAt** | Pointer to **time.Time** | Last time an alert of this type was resolved | [optional] 
**Count** | **int32** | Total count of the given alert type | 
**NetworkCount** | **int32** | Number of affected networks for the alert type | 
**DeviceTypes** | **[]string** | Affected device types for alerts in the group | 
**DeviceTags** | **[]string** | Unique device tags for alerts in the group | 
**Networks** | [**[]InlineResponse200321Networks**](InlineResponse200321Networks.md) | Affected networks for alerts in the group | 

## Methods

### NewInlineResponse200321Items

`func NewInlineResponse200321Items(type_ string, categoryType string, severity string, lastAlertedAt time.Time, count int32, networkCount int32, deviceTypes []string, deviceTags []string, networks []InlineResponse200321Networks, ) *InlineResponse200321Items`

NewInlineResponse200321Items instantiates a new InlineResponse200321Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200321ItemsWithDefaults

`func NewInlineResponse200321ItemsWithDefaults() *InlineResponse200321Items`

NewInlineResponse200321ItemsWithDefaults instantiates a new InlineResponse200321Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse200321Items) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200321Items) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200321Items) SetType(v string)`

SetType sets Type field to given value.


### GetCategoryType

`func (o *InlineResponse200321Items) GetCategoryType() string`

GetCategoryType returns the CategoryType field if non-nil, zero value otherwise.

### GetCategoryTypeOk

`func (o *InlineResponse200321Items) GetCategoryTypeOk() (*string, bool)`

GetCategoryTypeOk returns a tuple with the CategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryType

`func (o *InlineResponse200321Items) SetCategoryType(v string)`

SetCategoryType sets CategoryType field to given value.


### GetSeverity

`func (o *InlineResponse200321Items) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *InlineResponse200321Items) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *InlineResponse200321Items) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetLastAlertedAt

`func (o *InlineResponse200321Items) GetLastAlertedAt() time.Time`

GetLastAlertedAt returns the LastAlertedAt field if non-nil, zero value otherwise.

### GetLastAlertedAtOk

`func (o *InlineResponse200321Items) GetLastAlertedAtOk() (*time.Time, bool)`

GetLastAlertedAtOk returns a tuple with the LastAlertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAlertedAt

`func (o *InlineResponse200321Items) SetLastAlertedAt(v time.Time)`

SetLastAlertedAt sets LastAlertedAt field to given value.


### GetLastResolvedAt

`func (o *InlineResponse200321Items) GetLastResolvedAt() time.Time`

GetLastResolvedAt returns the LastResolvedAt field if non-nil, zero value otherwise.

### GetLastResolvedAtOk

`func (o *InlineResponse200321Items) GetLastResolvedAtOk() (*time.Time, bool)`

GetLastResolvedAtOk returns a tuple with the LastResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResolvedAt

`func (o *InlineResponse200321Items) SetLastResolvedAt(v time.Time)`

SetLastResolvedAt sets LastResolvedAt field to given value.

### HasLastResolvedAt

`func (o *InlineResponse200321Items) HasLastResolvedAt() bool`

HasLastResolvedAt returns a boolean if a field has been set.

### GetCount

`func (o *InlineResponse200321Items) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InlineResponse200321Items) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InlineResponse200321Items) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNetworkCount

`func (o *InlineResponse200321Items) GetNetworkCount() int32`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *InlineResponse200321Items) GetNetworkCountOk() (*int32, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *InlineResponse200321Items) SetNetworkCount(v int32)`

SetNetworkCount sets NetworkCount field to given value.


### GetDeviceTypes

`func (o *InlineResponse200321Items) GetDeviceTypes() []string`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *InlineResponse200321Items) GetDeviceTypesOk() (*[]string, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *InlineResponse200321Items) SetDeviceTypes(v []string)`

SetDeviceTypes sets DeviceTypes field to given value.


### GetDeviceTags

`func (o *InlineResponse200321Items) GetDeviceTags() []string`

GetDeviceTags returns the DeviceTags field if non-nil, zero value otherwise.

### GetDeviceTagsOk

`func (o *InlineResponse200321Items) GetDeviceTagsOk() (*[]string, bool)`

GetDeviceTagsOk returns a tuple with the DeviceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTags

`func (o *InlineResponse200321Items) SetDeviceTags(v []string)`

SetDeviceTags sets DeviceTags field to given value.


### GetNetworks

`func (o *InlineResponse200321Items) GetNetworks() []InlineResponse200321Networks`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *InlineResponse200321Items) GetNetworksOk() (*[]InlineResponse200321Networks, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *InlineResponse200321Items) SetNetworks(v []InlineResponse200321Networks)`

SetNetworks sets Networks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


