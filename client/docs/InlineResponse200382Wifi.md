# InlineResponse200382Wifi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Wi-Fi Enabled | [optional] 
**Type** | Pointer to **string** | Wi-Fi client type. Valid types are: visible, associated | [optional] 
**Flush** | Pointer to [**InlineResponse200382WifiFlush**](InlineResponse200382WifiFlush.md) |  | [optional] 
**AllowLists** | Pointer to [**InlineResponse200382WifiAllowLists**](InlineResponse200382WifiAllowLists.md) |  | [optional] 
**Hysteresis** | Pointer to [**InlineResponse200382WifiHysteresis**](InlineResponse200382WifiHysteresis.md) |  | [optional] 

## Methods

### NewInlineResponse200382Wifi

`func NewInlineResponse200382Wifi() *InlineResponse200382Wifi`

NewInlineResponse200382Wifi instantiates a new InlineResponse200382Wifi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200382WifiWithDefaults

`func NewInlineResponse200382WifiWithDefaults() *InlineResponse200382Wifi`

NewInlineResponse200382WifiWithDefaults instantiates a new InlineResponse200382Wifi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200382Wifi) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200382Wifi) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200382Wifi) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200382Wifi) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200382Wifi) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200382Wifi) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200382Wifi) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200382Wifi) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlush

`func (o *InlineResponse200382Wifi) GetFlush() InlineResponse200382WifiFlush`

GetFlush returns the Flush field if non-nil, zero value otherwise.

### GetFlushOk

`func (o *InlineResponse200382Wifi) GetFlushOk() (*InlineResponse200382WifiFlush, bool)`

GetFlushOk returns a tuple with the Flush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlush

`func (o *InlineResponse200382Wifi) SetFlush(v InlineResponse200382WifiFlush)`

SetFlush sets Flush field to given value.

### HasFlush

`func (o *InlineResponse200382Wifi) HasFlush() bool`

HasFlush returns a boolean if a field has been set.

### GetAllowLists

`func (o *InlineResponse200382Wifi) GetAllowLists() InlineResponse200382WifiAllowLists`

GetAllowLists returns the AllowLists field if non-nil, zero value otherwise.

### GetAllowListsOk

`func (o *InlineResponse200382Wifi) GetAllowListsOk() (*InlineResponse200382WifiAllowLists, bool)`

GetAllowListsOk returns a tuple with the AllowLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLists

`func (o *InlineResponse200382Wifi) SetAllowLists(v InlineResponse200382WifiAllowLists)`

SetAllowLists sets AllowLists field to given value.

### HasAllowLists

`func (o *InlineResponse200382Wifi) HasAllowLists() bool`

HasAllowLists returns a boolean if a field has been set.

### GetHysteresis

`func (o *InlineResponse200382Wifi) GetHysteresis() InlineResponse200382WifiHysteresis`

GetHysteresis returns the Hysteresis field if non-nil, zero value otherwise.

### GetHysteresisOk

`func (o *InlineResponse200382Wifi) GetHysteresisOk() (*InlineResponse200382WifiHysteresis, bool)`

GetHysteresisOk returns a tuple with the Hysteresis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHysteresis

`func (o *InlineResponse200382Wifi) SetHysteresis(v InlineResponse200382WifiHysteresis)`

SetHysteresis sets Hysteresis field to given value.

### HasHysteresis

`func (o *InlineResponse200382Wifi) HasHysteresis() bool`

HasHysteresis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


