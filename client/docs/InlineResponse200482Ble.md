# InlineResponse200482Ble

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | BLE Enabled | [optional] 
**Type** | Pointer to **string** | BLE type of clients to publish telemetry. Valid types are: all, ibeacon, eddystone, unknown | [optional] 
**Flush** | Pointer to [**InlineResponse200482BleFlush**](InlineResponse200482BleFlush.md) |  | [optional] 
**AllowLists** | Pointer to [**InlineResponse200482BleAllowLists**](InlineResponse200482BleAllowLists.md) |  | [optional] 
**Hysteresis** | Pointer to [**InlineResponse200482BleHysteresis**](InlineResponse200482BleHysteresis.md) |  | [optional] 

## Methods

### NewInlineResponse200482Ble

`func NewInlineResponse200482Ble() *InlineResponse200482Ble`

NewInlineResponse200482Ble instantiates a new InlineResponse200482Ble object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200482BleWithDefaults

`func NewInlineResponse200482BleWithDefaults() *InlineResponse200482Ble`

NewInlineResponse200482BleWithDefaults instantiates a new InlineResponse200482Ble object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200482Ble) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200482Ble) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200482Ble) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200482Ble) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200482Ble) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200482Ble) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200482Ble) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200482Ble) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFlush

`func (o *InlineResponse200482Ble) GetFlush() InlineResponse200482BleFlush`

GetFlush returns the Flush field if non-nil, zero value otherwise.

### GetFlushOk

`func (o *InlineResponse200482Ble) GetFlushOk() (*InlineResponse200482BleFlush, bool)`

GetFlushOk returns a tuple with the Flush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlush

`func (o *InlineResponse200482Ble) SetFlush(v InlineResponse200482BleFlush)`

SetFlush sets Flush field to given value.

### HasFlush

`func (o *InlineResponse200482Ble) HasFlush() bool`

HasFlush returns a boolean if a field has been set.

### GetAllowLists

`func (o *InlineResponse200482Ble) GetAllowLists() InlineResponse200482BleAllowLists`

GetAllowLists returns the AllowLists field if non-nil, zero value otherwise.

### GetAllowListsOk

`func (o *InlineResponse200482Ble) GetAllowListsOk() (*InlineResponse200482BleAllowLists, bool)`

GetAllowListsOk returns a tuple with the AllowLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLists

`func (o *InlineResponse200482Ble) SetAllowLists(v InlineResponse200482BleAllowLists)`

SetAllowLists sets AllowLists field to given value.

### HasAllowLists

`func (o *InlineResponse200482Ble) HasAllowLists() bool`

HasAllowLists returns a boolean if a field has been set.

### GetHysteresis

`func (o *InlineResponse200482Ble) GetHysteresis() InlineResponse200482BleHysteresis`

GetHysteresis returns the Hysteresis field if non-nil, zero value otherwise.

### GetHysteresisOk

`func (o *InlineResponse200482Ble) GetHysteresisOk() (*InlineResponse200482BleHysteresis, bool)`

GetHysteresisOk returns a tuple with the Hysteresis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHysteresis

`func (o *InlineResponse200482Ble) SetHysteresis(v InlineResponse200482BleHysteresis)`

SetHysteresis sets Hysteresis field to given value.

### HasHysteresis

`func (o *InlineResponse200482Ble) HasHysteresis() bool`

HasHysteresis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


