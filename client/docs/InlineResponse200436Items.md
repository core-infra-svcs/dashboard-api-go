# InlineResponse200436Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the switch. | [optional] 
**Serial** | Pointer to **string** | The serial number of the switch. | [optional] 
**Mac** | Pointer to **string** | The MAC address of the switch. | [optional] 
**Network** | Pointer to [**InlineResponse200434Network**](InlineResponse200434Network.md) |  | [optional] 
**Model** | Pointer to **string** | The model of the switch. | [optional] 
**Ports** | Pointer to [**[]InlineResponse200436Ports**](InlineResponse200436Ports.md) | The statuses of the ports on the switch. | [optional] 

## Methods

### NewInlineResponse200436Items

`func NewInlineResponse200436Items() *InlineResponse200436Items`

NewInlineResponse200436Items instantiates a new InlineResponse200436Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200436ItemsWithDefaults

`func NewInlineResponse200436ItemsWithDefaults() *InlineResponse200436Items`

NewInlineResponse200436ItemsWithDefaults instantiates a new InlineResponse200436Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200436Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200436Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200436Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200436Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200436Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200436Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200436Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200436Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200436Items) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200436Items) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200436Items) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200436Items) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200436Items) GetNetwork() InlineResponse200434Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200436Items) GetNetworkOk() (*InlineResponse200434Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200436Items) SetNetwork(v InlineResponse200434Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200436Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200436Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200436Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200436Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200436Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200436Items) GetPorts() []InlineResponse200436Ports`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200436Items) GetPortsOk() (*[]InlineResponse200436Ports, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200436Items) SetPorts(v []InlineResponse200436Ports)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200436Items) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


