# InlineResponse200275Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for the Catalyst appliance | [optional] 
**Ports** | Pointer to [**[]InlineResponse200275Ports**](InlineResponse200275Ports.md) | Ports associated with the given appliance | [optional] 
**Network** | Pointer to [**InlineResponse200275Network**](InlineResponse200275Network.md) |  | [optional] 

## Methods

### NewInlineResponse200275Items

`func NewInlineResponse200275Items() *InlineResponse200275Items`

NewInlineResponse200275Items instantiates a new InlineResponse200275Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200275ItemsWithDefaults

`func NewInlineResponse200275ItemsWithDefaults() *InlineResponse200275Items`

NewInlineResponse200275ItemsWithDefaults instantiates a new InlineResponse200275Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200275Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200275Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200275Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200275Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200275Items) GetPorts() []InlineResponse200275Ports`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200275Items) GetPortsOk() (*[]InlineResponse200275Ports, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200275Items) SetPorts(v []InlineResponse200275Ports)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200275Items) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200275Items) GetNetwork() InlineResponse200275Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200275Items) GetNetworkOk() (*InlineResponse200275Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200275Items) SetNetwork(v InlineResponse200275Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200275Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


