# InlineResponse200463Wifi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Wi-Fi Enabled | [optional] 
**Type** | Pointer to **string** | Wi-Fi client type. Valid types are: visible, associated | [optional] 
**Flush** | Pointer to [**InlineResponse200463WifiFlush**](InlineResponse200463WifiFlush.md) |  | [optional] 
**AllowLists** | Pointer to [**InlineResponse200463WifiAllowLists**](InlineResponse200463WifiAllowLists.md) |  | [optional] 
**Hysteresis** | Pointer to [**InlineResponse200463WifiHysteresis**](InlineResponse200463WifiHysteresis.md) |  | [optional] 

## Methods

### NewInlineResponse200463Wifi

`func NewInlineResponse200463Wifi() *InlineResponse200463Wifi`

NewInlineResponse200463Wifi instantiates a new InlineResponse200463Wifi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200463WifiWithDefaults

`func NewInlineResponse200463WifiWithDefaults() *InlineResponse200463Wifi`

NewInlineResponse200463WifiWithDefaults instantiates a new InlineResponse200463Wifi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200463Wifi) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200463Wifi) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200463Wifi) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200463Wifi) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200463Wifi) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200463Wifi) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200463Wifi) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200463Wifi) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlush

`func (o *InlineResponse200463Wifi) GetFlush() InlineResponse200463WifiFlush`

GetFlush returns the Flush field if non-nil, zero value otherwise.

### GetFlushOk

`func (o *InlineResponse200463Wifi) GetFlushOk() (*InlineResponse200463WifiFlush, bool)`

GetFlushOk returns a tuple with the Flush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlush

`func (o *InlineResponse200463Wifi) SetFlush(v InlineResponse200463WifiFlush)`

SetFlush sets Flush field to given value.

### HasFlush

`func (o *InlineResponse200463Wifi) HasFlush() bool`

HasFlush returns a boolean if a field has been set.

### GetAllowLists

`func (o *InlineResponse200463Wifi) GetAllowLists() InlineResponse200463WifiAllowLists`

GetAllowLists returns the AllowLists field if non-nil, zero value otherwise.

### GetAllowListsOk

`func (o *InlineResponse200463Wifi) GetAllowListsOk() (*InlineResponse200463WifiAllowLists, bool)`

GetAllowListsOk returns a tuple with the AllowLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLists

`func (o *InlineResponse200463Wifi) SetAllowLists(v InlineResponse200463WifiAllowLists)`

SetAllowLists sets AllowLists field to given value.

### HasAllowLists

`func (o *InlineResponse200463Wifi) HasAllowLists() bool`

HasAllowLists returns a boolean if a field has been set.

### GetHysteresis

`func (o *InlineResponse200463Wifi) GetHysteresis() InlineResponse200463WifiHysteresis`

GetHysteresis returns the Hysteresis field if non-nil, zero value otherwise.

### GetHysteresisOk

`func (o *InlineResponse200463Wifi) GetHysteresisOk() (*InlineResponse200463WifiHysteresis, bool)`

GetHysteresisOk returns a tuple with the Hysteresis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHysteresis

`func (o *InlineResponse200463Wifi) SetHysteresis(v InlineResponse200463WifiHysteresis)`

SetHysteresis sets Hysteresis field to given value.

### HasHysteresis

`func (o *InlineResponse200463Wifi) HasHysteresis() bool`

HasHysteresis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


