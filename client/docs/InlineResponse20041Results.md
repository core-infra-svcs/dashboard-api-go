# InlineResponse20041Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **int32** | The identifier for the port | [optional] 
**Interface** | Pointer to [**InlineResponse20041Interface**](InlineResponse20041Interface.md) |  | [optional] 
**Speed** | Pointer to **string** | The negotiated port speed | [optional] 
**Status** | Pointer to **string** | The current link status of the port | [optional] 
**Duplex** | Pointer to **bool** | Whether the port is in full-duplex mode | [optional] 
**Enabled** | Pointer to **bool** | Whether the port is administratively enabled | [optional] 
**Power** | Pointer to [**InlineResponse20041Power**](InlineResponse20041Power.md) |  | [optional] 

## Methods

### NewInlineResponse20041Results

`func NewInlineResponse20041Results() *InlineResponse20041Results`

NewInlineResponse20041Results instantiates a new InlineResponse20041Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20041ResultsWithDefaults

`func NewInlineResponse20041ResultsWithDefaults() *InlineResponse20041Results`

NewInlineResponse20041ResultsWithDefaults instantiates a new InlineResponse20041Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse20041Results) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse20041Results) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse20041Results) SetPortId(v int32)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse20041Results) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetInterface

`func (o *InlineResponse20041Results) GetInterface() InlineResponse20041Interface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse20041Results) GetInterfaceOk() (*InlineResponse20041Interface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse20041Results) SetInterface(v InlineResponse20041Interface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse20041Results) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetSpeed

`func (o *InlineResponse20041Results) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *InlineResponse20041Results) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *InlineResponse20041Results) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *InlineResponse20041Results) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20041Results) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20041Results) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20041Results) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20041Results) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDuplex

`func (o *InlineResponse20041Results) GetDuplex() bool`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *InlineResponse20041Results) GetDuplexOk() (*bool, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *InlineResponse20041Results) SetDuplex(v bool)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *InlineResponse20041Results) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20041Results) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20041Results) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20041Results) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20041Results) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPower

`func (o *InlineResponse20041Results) GetPower() InlineResponse20041Power`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *InlineResponse20041Results) GetPowerOk() (*InlineResponse20041Power, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *InlineResponse20041Results) SetPower(v InlineResponse20041Power)`

SetPower sets Power field to given value.

### HasPower

`func (o *InlineResponse20041Results) HasPower() bool`

HasPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


