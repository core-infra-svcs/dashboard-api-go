# InlineResponse200476Ports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module. | [optional] 
**Counts** | Pointer to [**InlineResponse200476Counts**](InlineResponse200476Counts.md) |  | [optional] 

## Methods

### NewInlineResponse200476Ports

`func NewInlineResponse200476Ports() *InlineResponse200476Ports`

NewInlineResponse200476Ports instantiates a new InlineResponse200476Ports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200476PortsWithDefaults

`func NewInlineResponse200476PortsWithDefaults() *InlineResponse200476Ports`

NewInlineResponse200476PortsWithDefaults instantiates a new InlineResponse200476Ports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *InlineResponse200476Ports) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InlineResponse200476Ports) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InlineResponse200476Ports) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *InlineResponse200476Ports) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetCounts

`func (o *InlineResponse200476Ports) GetCounts() InlineResponse200476Counts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *InlineResponse200476Ports) GetCountsOk() (*InlineResponse200476Counts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *InlineResponse200476Ports) SetCounts(v InlineResponse200476Counts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *InlineResponse200476Ports) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


