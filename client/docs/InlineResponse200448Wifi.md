# InlineResponse200448Wifi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Wi-Fi Enabled | [optional] 
**Type** | Pointer to **string** | Wi-Fi client type. Valid types are: visible, associated | [optional] 
**Flush** | Pointer to [**InlineResponse200448WifiFlush**](InlineResponse200448WifiFlush.md) |  | [optional] 
**AllowLists** | Pointer to [**InlineResponse200448WifiAllowLists**](InlineResponse200448WifiAllowLists.md) |  | [optional] 
**Hysteresis** | Pointer to [**InlineResponse200448WifiHysteresis**](InlineResponse200448WifiHysteresis.md) |  | [optional] 

## Methods

### NewInlineResponse200448Wifi

`func NewInlineResponse200448Wifi() *InlineResponse200448Wifi`

NewInlineResponse200448Wifi instantiates a new InlineResponse200448Wifi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200448WifiWithDefaults

`func NewInlineResponse200448WifiWithDefaults() *InlineResponse200448Wifi`

NewInlineResponse200448WifiWithDefaults instantiates a new InlineResponse200448Wifi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200448Wifi) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200448Wifi) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200448Wifi) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200448Wifi) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200448Wifi) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200448Wifi) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200448Wifi) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200448Wifi) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlush

`func (o *InlineResponse200448Wifi) GetFlush() InlineResponse200448WifiFlush`

GetFlush returns the Flush field if non-nil, zero value otherwise.

### GetFlushOk

`func (o *InlineResponse200448Wifi) GetFlushOk() (*InlineResponse200448WifiFlush, bool)`

GetFlushOk returns a tuple with the Flush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlush

`func (o *InlineResponse200448Wifi) SetFlush(v InlineResponse200448WifiFlush)`

SetFlush sets Flush field to given value.

### HasFlush

`func (o *InlineResponse200448Wifi) HasFlush() bool`

HasFlush returns a boolean if a field has been set.

### GetAllowLists

`func (o *InlineResponse200448Wifi) GetAllowLists() InlineResponse200448WifiAllowLists`

GetAllowLists returns the AllowLists field if non-nil, zero value otherwise.

### GetAllowListsOk

`func (o *InlineResponse200448Wifi) GetAllowListsOk() (*InlineResponse200448WifiAllowLists, bool)`

GetAllowListsOk returns a tuple with the AllowLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLists

`func (o *InlineResponse200448Wifi) SetAllowLists(v InlineResponse200448WifiAllowLists)`

SetAllowLists sets AllowLists field to given value.

### HasAllowLists

`func (o *InlineResponse200448Wifi) HasAllowLists() bool`

HasAllowLists returns a boolean if a field has been set.

### GetHysteresis

`func (o *InlineResponse200448Wifi) GetHysteresis() InlineResponse200448WifiHysteresis`

GetHysteresis returns the Hysteresis field if non-nil, zero value otherwise.

### GetHysteresisOk

`func (o *InlineResponse200448Wifi) GetHysteresisOk() (*InlineResponse200448WifiHysteresis, bool)`

GetHysteresisOk returns a tuple with the Hysteresis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHysteresis

`func (o *InlineResponse200448Wifi) SetHysteresis(v InlineResponse200448WifiHysteresis)`

SetHysteresis sets Hysteresis field to given value.

### HasHysteresis

`func (o *InlineResponse200448Wifi) HasHysteresis() bool`

HasHysteresis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


