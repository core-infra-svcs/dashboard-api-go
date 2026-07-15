# InlineResponse200434Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module. | [optional] 
**Counts** | Pointer to [**InlineResponse200434Counts**](InlineResponse200434Counts.md) |  | [optional] 

## Methods

### NewInlineResponse200434Ports

`func NewInlineResponse200434Ports() *InlineResponse200434Ports`

NewInlineResponse200434Ports instantiates a new InlineResponse200434Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200434PortsWithDefaults

`func NewInlineResponse200434PortsWithDefaults() *InlineResponse200434Ports`

NewInlineResponse200434PortsWithDefaults instantiates a new InlineResponse200434Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse200434Ports) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse200434Ports) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse200434Ports) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse200434Ports) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200434Ports) GetCounts() InlineResponse200434Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200434Ports) GetCountsOk() (*InlineResponse200434Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200434Ports) SetCounts(v InlineResponse200434Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200434Ports) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


