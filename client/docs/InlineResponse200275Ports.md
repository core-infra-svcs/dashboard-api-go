# InlineResponse200275Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | Unique identifier for the port | [optional] 
**InterfaceName** | Pointer to **string** | Interface name for the port (e.g. FiveGigabitEthernet0/0/9) | [optional] 
**Indices** | Pointer to [**InlineResponse200275Indices**](InlineResponse200275Indices.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200275Readings**](InlineResponse200275Readings.md) | The digital optical monitoring readings for the individual port, with the most recent snapshot first | [optional] 

## Methods

### NewInlineResponse200275Ports

`func NewInlineResponse200275Ports() *InlineResponse200275Ports`

NewInlineResponse200275Ports instantiates a new InlineResponse200275Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200275PortsWithDefaults

`func NewInlineResponse200275PortsWithDefaults() *InlineResponse200275Ports`

NewInlineResponse200275PortsWithDefaults instantiates a new InlineResponse200275Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse200275Ports) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse200275Ports) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse200275Ports) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse200275Ports) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *InlineResponse200275Ports) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *InlineResponse200275Ports) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *InlineResponse200275Ports) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *InlineResponse200275Ports) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIndices

`func (o *InlineResponse200275Ports) GetIndices() InlineResponse200275Indices`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *InlineResponse200275Ports) GetIndicesOk() (*InlineResponse200275Indices, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *InlineResponse200275Ports) SetIndices(v InlineResponse200275Indices)`

SetIndices sets Indices field to given value.

### HasIndices

`func (o *InlineResponse200275Ports) HasIndices() bool`

HasIndices returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200275Ports) GetReadings() []InlineResponse200275Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200275Ports) GetReadingsOk() (*[]InlineResponse200275Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200275Ports) SetReadings(v []InlineResponse200275Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200275Ports) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


