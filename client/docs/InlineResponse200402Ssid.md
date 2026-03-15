# InlineResponse200402Ssid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The SSID name | [optional] 
**Number** | Pointer to **int32** | The SSID number | [optional] 
**Enabled** | Pointer to **bool** | Whether the SSID is enabled | [optional] 
**OpenRoaming** | Pointer to [**InlineResponse200402OpenRoaming**](InlineResponse200402OpenRoaming.md) |  | [optional] 

## Methods

### NewInlineResponse200402Ssid

`func NewInlineResponse200402Ssid() *InlineResponse200402Ssid`

NewInlineResponse200402Ssid instantiates a new InlineResponse200402Ssid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200402SsidWithDefaults

`func NewInlineResponse200402SsidWithDefaults() *InlineResponse200402Ssid`

NewInlineResponse200402SsidWithDefaults instantiates a new InlineResponse200402Ssid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200402Ssid) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200402Ssid) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200402Ssid) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200402Ssid) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *InlineResponse200402Ssid) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse200402Ssid) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse200402Ssid) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse200402Ssid) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200402Ssid) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200402Ssid) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200402Ssid) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200402Ssid) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOpenRoaming

`func (o *InlineResponse200402Ssid) GetOpenRoaming() InlineResponse200402OpenRoaming`

GetOpenRoaming returns the OpenRoaming field if non-nil, zero value otherwise.

### GetOpenRoamingOk

`func (o *InlineResponse200402Ssid) GetOpenRoamingOk() (*InlineResponse200402OpenRoaming, bool)`

GetOpenRoamingOk returns a tuple with the OpenRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenRoaming

`func (o *InlineResponse200402Ssid) SetOpenRoaming(v InlineResponse200402OpenRoaming)`

SetOpenRoaming sets OpenRoaming field to given value.

### HasOpenRoaming

`func (o *InlineResponse200402Ssid) HasOpenRoaming() bool`

HasOpenRoaming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


