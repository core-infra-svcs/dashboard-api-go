# InlineResponse200284Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | Unique identifier for the port | [optional] 
**InterfaceName** | Pointer to **string** | Interface name for the port (e.g. FiveGigabitEthernet0/0/9) | [optional] 
**Indices** | Pointer to [**InlineResponse200284Indices**](InlineResponse200284Indices.md) |  | [optional] 
**Readings** | Pointer to [**[]InlineResponse200284Readings**](InlineResponse200284Readings.md) | The digital optical monitoring readings for the individual port, with the most recent snapshot first | [optional] 

## Methods

### NewInlineResponse200284Ports

`func NewInlineResponse200284Ports() *InlineResponse200284Ports`

NewInlineResponse200284Ports instantiates a new InlineResponse200284Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200284PortsWithDefaults

`func NewInlineResponse200284PortsWithDefaults() *InlineResponse200284Ports`

NewInlineResponse200284PortsWithDefaults instantiates a new InlineResponse200284Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse200284Ports) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse200284Ports) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse200284Ports) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse200284Ports) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *InlineResponse200284Ports) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *InlineResponse200284Ports) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *InlineResponse200284Ports) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *InlineResponse200284Ports) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIndices

`func (o *InlineResponse200284Ports) GetIndices() InlineResponse200284Indices`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *InlineResponse200284Ports) GetIndicesOk() (*InlineResponse200284Indices, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *InlineResponse200284Ports) SetIndices(v InlineResponse200284Indices)`

SetIndices sets Indices field to given value.

### HasIndices

`func (o *InlineResponse200284Ports) HasIndices() bool`

HasIndices returns a boolean if a field has been set.

### GetReadings

`func (o *InlineResponse200284Ports) GetReadings() []InlineResponse200284Readings`

GetReadings returns the Readings field if non-nil, zero value otherwise.

### GetReadingsOk

`func (o *InlineResponse200284Ports) GetReadingsOk() (*[]InlineResponse200284Readings, bool)`

GetReadingsOk returns a tuple with the Readings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadings

`func (o *InlineResponse200284Ports) SetReadings(v []InlineResponse200284Readings)`

SetReadings sets Readings field to given value.

### HasReadings

`func (o *InlineResponse200284Ports) HasReadings() bool`

HasReadings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


