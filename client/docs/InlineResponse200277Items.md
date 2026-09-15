# InlineResponse200277Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for the Catalyst appliance | [optional] 
**Ports** | Pointer to [**[]InlineResponse200277Ports**](InlineResponse200277Ports.md) | Ports associated with the given appliance | [optional] 
**Network** | Pointer to [**InlineResponse200277Network**](InlineResponse200277Network.md) |  | [optional] 

## Methods

### NewInlineResponse200277Items

`func NewInlineResponse200277Items() *InlineResponse200277Items`

NewInlineResponse200277Items instantiates a new InlineResponse200277Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200277ItemsWithDefaults

`func NewInlineResponse200277ItemsWithDefaults() *InlineResponse200277Items`

NewInlineResponse200277ItemsWithDefaults instantiates a new InlineResponse200277Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200277Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200277Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200277Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200277Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetPorts

`func (o *InlineResponse200277Items) GetPorts() []InlineResponse200277Ports`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InlineResponse200277Items) GetPortsOk() (*[]InlineResponse200277Ports, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InlineResponse200277Items) SetPorts(v []InlineResponse200277Ports)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InlineResponse200277Items) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200277Items) GetNetwork() InlineResponse200277Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200277Items) GetNetworkOk() (*InlineResponse200277Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200277Items) SetNetwork(v InlineResponse200277Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200277Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


