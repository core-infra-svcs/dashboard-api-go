# InlineResponse20040Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **int32** | The identifier for the port | [optional] 
**Interface** | Pointer to [**InlineResponse20040Interface**](InlineResponse20040Interface.md) |  | [optional] 
**Speed** | Pointer to **string** | The negotiated port speed | [optional] 
**Status** | Pointer to **string** | The current link status of the port | [optional] 
**Duplex** | Pointer to **bool** | Whether the port is in full-duplex mode | [optional] 
**Enabled** | Pointer to **bool** | Whether the port is administratively enabled | [optional] 
**Power** | Pointer to [**InlineResponse20040Power**](InlineResponse20040Power.md) |  | [optional] 

## Methods

### NewInlineResponse20040Results

`func NewInlineResponse20040Results() *InlineResponse20040Results`

NewInlineResponse20040Results instantiates a new InlineResponse20040Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20040ResultsWithDefaults

`func NewInlineResponse20040ResultsWithDefaults() *InlineResponse20040Results`

NewInlineResponse20040ResultsWithDefaults instantiates a new InlineResponse20040Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse20040Results) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse20040Results) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse20040Results) SetPortId(v int32)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse20040Results) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetInterface

`func (o *InlineResponse20040Results) GetInterface() InlineResponse20040Interface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse20040Results) GetInterfaceOk() (*InlineResponse20040Interface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse20040Results) SetInterface(v InlineResponse20040Interface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse20040Results) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetSpeed

`func (o *InlineResponse20040Results) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *InlineResponse20040Results) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *InlineResponse20040Results) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *InlineResponse20040Results) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20040Results) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20040Results) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20040Results) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20040Results) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDuplex

`func (o *InlineResponse20040Results) GetDuplex() bool`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *InlineResponse20040Results) GetDuplexOk() (*bool, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *InlineResponse20040Results) SetDuplex(v bool)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *InlineResponse20040Results) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20040Results) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20040Results) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20040Results) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20040Results) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPower

`func (o *InlineResponse20040Results) GetPower() InlineResponse20040Power`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *InlineResponse20040Results) GetPowerOk() (*InlineResponse20040Power, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *InlineResponse20040Results) SetPower(v InlineResponse20040Power)`

SetPower sets Power field to given value.

### HasPower

`func (o *InlineResponse20040Results) HasPower() bool`

HasPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


