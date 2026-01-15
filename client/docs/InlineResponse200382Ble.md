# InlineResponse200382Ble

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | BLE Enabled | [optional] 
**Type** | Pointer to **string** | BLE type of clients to publish telemetry. Valid types are: all, ibeacon, eddystone, unknown | [optional] 
**Flush** | Pointer to [**InlineResponse200382BleFlush**](InlineResponse200382BleFlush.md) |  | [optional] 
**AllowLists** | Pointer to [**InlineResponse200382BleAllowLists**](InlineResponse200382BleAllowLists.md) |  | [optional] 
**Hysteresis** | Pointer to [**InlineResponse200382BleHysteresis**](InlineResponse200382BleHysteresis.md) |  | [optional] 

## Methods

### NewInlineResponse200382Ble

`func NewInlineResponse200382Ble() *InlineResponse200382Ble`

NewInlineResponse200382Ble instantiates a new InlineResponse200382Ble object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200382BleWithDefaults

`func NewInlineResponse200382BleWithDefaults() *InlineResponse200382Ble`

NewInlineResponse200382BleWithDefaults instantiates a new InlineResponse200382Ble object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200382Ble) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200382Ble) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200382Ble) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200382Ble) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200382Ble) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200382Ble) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200382Ble) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200382Ble) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlush

`func (o *InlineResponse200382Ble) GetFlush() InlineResponse200382BleFlush`

GetFlush returns the Flush field if non-nil, zero value otherwise.

### GetFlushOk

`func (o *InlineResponse200382Ble) GetFlushOk() (*InlineResponse200382BleFlush, bool)`

GetFlushOk returns a tuple with the Flush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlush

`func (o *InlineResponse200382Ble) SetFlush(v InlineResponse200382BleFlush)`

SetFlush sets Flush field to given value.

### HasFlush

`func (o *InlineResponse200382Ble) HasFlush() bool`

HasFlush returns a boolean if a field has been set.

### GetAllowLists

`func (o *InlineResponse200382Ble) GetAllowLists() InlineResponse200382BleAllowLists`

GetAllowLists returns the AllowLists field if non-nil, zero value otherwise.

### GetAllowListsOk

`func (o *InlineResponse200382Ble) GetAllowListsOk() (*InlineResponse200382BleAllowLists, bool)`

GetAllowListsOk returns a tuple with the AllowLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLists

`func (o *InlineResponse200382Ble) SetAllowLists(v InlineResponse200382BleAllowLists)`

SetAllowLists sets AllowLists field to given value.

### HasAllowLists

`func (o *InlineResponse200382Ble) HasAllowLists() bool`

HasAllowLists returns a boolean if a field has been set.

### GetHysteresis

`func (o *InlineResponse200382Ble) GetHysteresis() InlineResponse200382BleHysteresis`

GetHysteresis returns the Hysteresis field if non-nil, zero value otherwise.

### GetHysteresisOk

`func (o *InlineResponse200382Ble) GetHysteresisOk() (*InlineResponse200382BleHysteresis, bool)`

GetHysteresisOk returns a tuple with the Hysteresis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHysteresis

`func (o *InlineResponse200382Ble) SetHysteresis(v InlineResponse200382BleHysteresis)`

SetHysteresis sets Hysteresis field to given value.

### HasHysteresis

`func (o *InlineResponse200382Ble) HasHysteresis() bool`

HasHysteresis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


