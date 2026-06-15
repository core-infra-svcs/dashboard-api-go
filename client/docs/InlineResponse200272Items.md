# InlineResponse200272Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for the Catalyst appliance | [optional] 
**Ports** | Pointer to [**[]InlineResponse200272Ports**](InlineResponse200272Ports.md) | Ports associated with the given appliance | [optional] 
**Network** | Pointer to [**InlineResponse200272Network**](InlineResponse200272Network.md) |  | [optional] 

## Methods

### NewInlineResponse200272Items

`func NewInlineResponse200272Items() *InlineResponse200272Items`

NewInlineResponse200272Items instantiates a new InlineResponse200272Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200272ItemsWithDefaults

`func NewInlineResponse200272ItemsWithDefaults() *InlineResponse200272Items`

NewInlineResponse200272ItemsWithDefaults instantiates a new InlineResponse200272Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200272Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200272Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200272Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200272Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200272Items) GetPorts() []InlineResponse200272Ports`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200272Items) GetPortsOk() (*[]InlineResponse200272Ports, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200272Items) SetPorts(v []InlineResponse200272Ports)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200272Items) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200272Items) GetNetwork() InlineResponse200272Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200272Items) GetNetworkOk() (*InlineResponse200272Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200272Items) SetNetwork(v InlineResponse200272Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200272Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


