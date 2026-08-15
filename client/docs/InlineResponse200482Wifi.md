# InlineResponse200482Wifi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Wi-Fi Enabled | [optional] 
**Type** | Pointer to **string** | Wi-Fi client type. Valid types are: visible, associated | [optional] 
**Flush** | Pointer to [**InlineResponse200482WifiFlush**](InlineResponse200482WifiFlush.md) |  | [optional] 
**AllowLists** | Pointer to [**InlineResponse200482WifiAllowLists**](InlineResponse200482WifiAllowLists.md) |  | [optional] 
**Hysteresis** | Pointer to [**InlineResponse200482WifiHysteresis**](InlineResponse200482WifiHysteresis.md) |  | [optional] 

## Methods

### NewInlineResponse200482Wifi

`func NewInlineResponse200482Wifi() *InlineResponse200482Wifi`

NewInlineResponse200482Wifi instantiates a new InlineResponse200482Wifi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200482WifiWithDefaults

`func NewInlineResponse200482WifiWithDefaults() *InlineResponse200482Wifi`

NewInlineResponse200482WifiWithDefaults instantiates a new InlineResponse200482Wifi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200482Wifi) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200482Wifi) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200482Wifi) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200482Wifi) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200482Wifi) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200482Wifi) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200482Wifi) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200482Wifi) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlush

`func (o *InlineResponse200482Wifi) GetFlush() InlineResponse200482WifiFlush`

GetFlush returns the Flush field if non-nil, zero value otherwise.

### GetFlushOk

`func (o *InlineResponse200482Wifi) GetFlushOk() (*InlineResponse200482WifiFlush, bool)`

GetFlushOk returns a tuple with the Flush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlush

`func (o *InlineResponse200482Wifi) SetFlush(v InlineResponse200482WifiFlush)`

SetFlush sets Flush field to given value.

### HasFlush

`func (o *InlineResponse200482Wifi) HasFlush() bool`

HasFlush returns a boolean if a field has been set.

### GetAllowLists

`func (o *InlineResponse200482Wifi) GetAllowLists() InlineResponse200482WifiAllowLists`

GetAllowLists returns the AllowLists field if non-nil, zero value otherwise.

### GetAllowListsOk

`func (o *InlineResponse200482Wifi) GetAllowListsOk() (*InlineResponse200482WifiAllowLists, bool)`

GetAllowListsOk returns a tuple with the AllowLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLists

`func (o *InlineResponse200482Wifi) SetAllowLists(v InlineResponse200482WifiAllowLists)`

SetAllowLists sets AllowLists field to given value.

### HasAllowLists

`func (o *InlineResponse200482Wifi) HasAllowLists() bool`

HasAllowLists returns a boolean if a field has been set.

### GetHysteresis

`func (o *InlineResponse200482Wifi) GetHysteresis() InlineResponse200482WifiHysteresis`

GetHysteresis returns the Hysteresis field if non-nil, zero value otherwise.

### GetHysteresisOk

`func (o *InlineResponse200482Wifi) GetHysteresisOk() (*InlineResponse200482WifiHysteresis, bool)`

GetHysteresisOk returns a tuple with the Hysteresis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHysteresis

`func (o *InlineResponse200482Wifi) SetHysteresis(v InlineResponse200482WifiHysteresis)`

SetHysteresis sets Hysteresis field to given value.

### HasHysteresis

`func (o *InlineResponse200482Wifi) HasHysteresis() bool`

HasHysteresis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


