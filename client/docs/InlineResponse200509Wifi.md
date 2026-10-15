# InlineResponse200509Wifi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Wi-Fi Enabled | [optional] 
**Type** | Pointer to **string** | Wi-Fi client type. Valid types are: visible, associated | [optional] 
**Flush** | Pointer to [**InlineResponse200509WifiFlush**](InlineResponse200509WifiFlush.md) |  | [optional] 
**AllowLists** | Pointer to [**InlineResponse200509WifiAllowLists**](InlineResponse200509WifiAllowLists.md) |  | [optional] 
**Hysteresis** | Pointer to [**InlineResponse200509WifiHysteresis**](InlineResponse200509WifiHysteresis.md) |  | [optional] 

## Methods

### NewInlineResponse200509Wifi

`func NewInlineResponse200509Wifi() *InlineResponse200509Wifi`

NewInlineResponse200509Wifi instantiates a new InlineResponse200509Wifi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200509WifiWithDefaults

`func NewInlineResponse200509WifiWithDefaults() *InlineResponse200509Wifi`

NewInlineResponse200509WifiWithDefaults instantiates a new InlineResponse200509Wifi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200509Wifi) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200509Wifi) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200509Wifi) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200509Wifi) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200509Wifi) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200509Wifi) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200509Wifi) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200509Wifi) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlush

`func (o *InlineResponse200509Wifi) GetFlush() InlineResponse200509WifiFlush`

GetFlush returns the Flush field if non-nil, zero value otherwise.

### GetFlushOk

`func (o *InlineResponse200509Wifi) GetFlushOk() (*InlineResponse200509WifiFlush, bool)`

GetFlushOk returns a tuple with the Flush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlush

`func (o *InlineResponse200509Wifi) SetFlush(v InlineResponse200509WifiFlush)`

SetFlush sets Flush field to given value.

### HasFlush

`func (o *InlineResponse200509Wifi) HasFlush() bool`

HasFlush returns a boolean if a field has been set.

### GetAllowLists

`func (o *InlineResponse200509Wifi) GetAllowLists() InlineResponse200509WifiAllowLists`

GetAllowLists returns the AllowLists field if non-nil, zero value otherwise.

### GetAllowListsOk

`func (o *InlineResponse200509Wifi) GetAllowListsOk() (*InlineResponse200509WifiAllowLists, bool)`

GetAllowListsOk returns a tuple with the AllowLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLists

`func (o *InlineResponse200509Wifi) SetAllowLists(v InlineResponse200509WifiAllowLists)`

SetAllowLists sets AllowLists field to given value.

### HasAllowLists

`func (o *InlineResponse200509Wifi) HasAllowLists() bool`

HasAllowLists returns a boolean if a field has been set.

### GetHysteresis

`func (o *InlineResponse200509Wifi) GetHysteresis() InlineResponse200509WifiHysteresis`

GetHysteresis returns the Hysteresis field if non-nil, zero value otherwise.

### GetHysteresisOk

`func (o *InlineResponse200509Wifi) GetHysteresisOk() (*InlineResponse200509WifiHysteresis, bool)`

GetHysteresisOk returns a tuple with the Hysteresis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHysteresis

`func (o *InlineResponse200509Wifi) SetHysteresis(v InlineResponse200509WifiHysteresis)`

SetHysteresis sets Hysteresis field to given value.

### HasHysteresis

`func (o *InlineResponse200509Wifi) HasHysteresis() bool`

HasHysteresis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


